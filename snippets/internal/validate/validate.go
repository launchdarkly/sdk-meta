package validate

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/launchdarkly/sdk-meta/snippets/internal/model"
	"github.com/launchdarkly/sdk-meta/snippets/internal/render"
)

// Config controls a validator run.
type Config struct {
	SDKsDir       string // path to sdks/
	ValidatorsDir string // path to validators/
	SDK           string // sdk id to validate (empty = all)
}

// Run finds validatable snippets under cfg.SDKsDir and runs each through
// the per-language Docker validator. First-pass implementation: python only.
//
// Snippets are run against a real LaunchDarkly environment. Required env vars,
// matching the convention used by the hello-* sample apps:
//
//	LAUNCHDARKLY_SDK_KEY    server-side SDK key for the test environment
//	LAUNCHDARKLY_FLAG_KEY   the flag key the snippet should evaluate
//
// These are read from the caller's environment and forwarded into the
// per-snippet Docker run. They are never written to a file in the repo.
func Run(cfg Config) error {
	sdkKey := os.Getenv("LAUNCHDARKLY_SDK_KEY")
	flagKey := os.Getenv("LAUNCHDARKLY_FLAG_KEY")
	if sdkKey == "" || flagKey == "" {
		return fmt.Errorf("LAUNCHDARKLY_SDK_KEY and LAUNCHDARKLY_FLAG_KEY must be set in the caller environment")
	}

	snippets, err := model.LoadAll(cfg.SDKsDir)
	if err != nil {
		return err
	}

	any := false
	for _, id := range model.SortedIDs(snippets) {
		s := snippets[id]
		if cfg.SDK != "" && s.Frontmatter.SDK != cfg.SDK {
			continue
		}
		if s.Frontmatter.Validation.Entrypoint == "" {
			continue
		}
		any = true
		if err := runOne(cfg, s, sdkKey, flagKey); err != nil {
			return fmt.Errorf("validate %s: %w", id, err)
		}
	}
	if !any {
		return fmt.Errorf("no validatable snippets found (sdk=%q)", cfg.SDK)
	}
	return nil
}

func runOne(cfg Config, s *model.Snippet, sdkKey, flagKey string) error {
	switch s.CodeLang {
	case "python":
		return runPython(cfg, s, sdkKey, flagKey)
	default:
		return fmt.Errorf("no validator for lang=%q", s.CodeLang)
	}
}

func runPython(cfg Config, s *model.Snippet, sdkKey, flagKey string) error {
	if err := checkEntrypoint(s.Frontmatter.Validation.Entrypoint); err != nil {
		return err
	}
	if err := checkRequirements(s.Frontmatter.Validation.Requirements); err != nil {
		return err
	}

	inputs, err := runtimeInputs(s, sdkKey, flagKey)
	if err != nil {
		return err
	}
	nodes, err := render.Parse(s.CodeBody)
	if err != nil {
		return err
	}
	code, err := render.RenderRuntime(nodes, inputs)
	if err != nil {
		return err
	}

	stageDir, err := os.MkdirTemp("", "snippets-validate-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(stageDir)

	entrypoint := s.Frontmatter.Validation.Entrypoint
	if err := os.WriteFile(filepath.Join(stageDir, entrypoint), []byte(code), 0o644); err != nil {
		return err
	}
	if s.Frontmatter.Validation.Requirements != "" {
		if err := os.WriteFile(filepath.Join(stageDir, "requirements.txt"),
			[]byte(s.Frontmatter.Validation.Requirements+"\n"), 0o644); err != nil {
			return err
		}
	}

	validatorDir := filepath.Join(cfg.ValidatorsDir, "languages", "python")
	if _, err := os.Stat(filepath.Join(validatorDir, "Dockerfile")); err != nil {
		return fmt.Errorf("validator Dockerfile not found at %s: %w", validatorDir, err)
	}

	// Tag the image by hash of its build context. Two concurrent validate runs
	// on the same Docker host won't race on a shared mutable tag, and rebuilds
	// are skipped automatically when the validator hasn't changed.
	tag, err := validatorImageTag(validatorDir)
	if err != nil {
		return err
	}
	build := exec.Command("docker", "build", "--quiet", "-t", tag, validatorDir)
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		return fmt.Errorf("docker build failed: %w", err)
	}

	fmt.Printf("--- validate %s (lang=%s, entrypoint=%s) ---\n", s.Frontmatter.ID, s.CodeLang, entrypoint)

	run := exec.Command("docker", "run", "--rm",
		"-v", stageDir+":/snippet:ro",
		"-e", "SNIPPET_ENTRYPOINT="+entrypoint,
		"-e", "LAUNCHDARKLY_SDK_KEY="+sdkKey,
		"-e", "LAUNCHDARKLY_FLAG_KEY="+flagKey,
		tag,
	)
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	if err := run.Run(); err != nil {
		return fmt.Errorf("snippet runtime validation failed: %w", err)
	}
	return nil
}

// checkEntrypoint rejects any value that isn't a plain filename. Snippet
// frontmatter is author-controlled, so without this guard a malicious
// `entrypoint: ../../../etc/foo` would let `os.WriteFile(stageDir+entrypoint)`
// land outside the staging directory.
func checkEntrypoint(entrypoint string) error {
	if entrypoint == "" {
		return nil
	}
	if entrypoint != filepath.Base(entrypoint) {
		return fmt.Errorf("validation.entrypoint %q must be a plain filename (no path separators or ..)", entrypoint)
	}
	if entrypoint == "." || entrypoint == ".." {
		return fmt.Errorf("validation.entrypoint %q is not a valid filename", entrypoint)
	}
	return nil
}

// checkRequirements rejects values that would let a snippet author smuggle
// pip flags through the requirements.txt that the validator writes. The
// allow-list is "one or more requirement specifiers, separated by single
// newlines, none starting with `-`". This blocks `--extra-index-url` /
// `--index-url` / `-r other.txt` style escapes.
func checkRequirements(req string) error {
	if req == "" {
		return nil
	}
	for i, line := range strings.Split(req, "\n") {
		trim := strings.TrimSpace(line)
		if trim == "" {
			continue
		}
		if strings.HasPrefix(trim, "-") {
			return fmt.Errorf("validation.requirements line %d %q starts with '-' (pip flags are not allowed)", i+1, trim)
		}
	}
	return nil
}

// validatorImageTag produces a Docker tag that's a content hash of the
// validator directory: a deterministic tag that changes only when a file in
// the validator changes. Concurrent validate runs against the same validator
// thus reuse the cached image; runs against different validators get different
// tags so they cannot interleave.
func validatorImageTag(dir string) (string, error) {
	h := sha256.New()
	err := filepath.WalkDir(dir, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		f, err := os.Open(p)
		if err != nil {
			return err
		}
		defer f.Close()
		rel, _ := filepath.Rel(dir, p)
		fmt.Fprintf(h, "%s\x00", rel)
		_, err = io.Copy(h, f)
		return err
	})
	if err != nil {
		return "", err
	}
	return "sdk-snippets/python-validator:" + hex.EncodeToString(h.Sum(nil))[:16], nil
}

// runtimeInputs derives concrete values for every declared input.
//
// Inputs typed as `flag-key` use LAUNCHDARKLY_FLAG_KEY; inputs typed as
// `sdk-key` use LAUNCHDARKLY_SDK_KEY. Both come from the caller's env so the
// snippet's rendered output never embeds a real key. Other inputs fall back
// to the snippet's own runtime-default. Declaring runtime-default for either
// keyed type is an error: the value must always come from the environment.
func runtimeInputs(s *model.Snippet, sdkKey, flagKey string) (map[string]string, error) {
	out := map[string]string{}
	for name, in := range s.Frontmatter.Inputs {
		switch in.Type {
		case "flag-key":
			if in.RuntimeDefault != "" {
				return nil, fmt.Errorf("input %q (type=flag-key) must not declare runtime-default — value comes from LAUNCHDARKLY_FLAG_KEY", name)
			}
			out[name] = flagKey
		case "sdk-key":
			if in.RuntimeDefault != "" {
				return nil, fmt.Errorf("input %q (type=sdk-key) must not declare runtime-default — value comes from LAUNCHDARKLY_SDK_KEY", name)
			}
			out[name] = sdkKey
		default:
			out[name] = in.RuntimeDefault
		}
	}
	return out, nil
}
