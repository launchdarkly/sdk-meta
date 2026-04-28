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

// envInputs holds the environment-derived input values that get substituted
// into snippets at validation time. Each field maps to one EXAM-HELLO env
// var and to a snippet-input type.
type envInputs struct {
	sdkKey        string // LAUNCHDARKLY_SDK_KEY        ↔ type: sdk-key
	flagKey       string // LAUNCHDARKLY_FLAG_KEY       ↔ type: flag-key
	mobileKey     string // LAUNCHDARKLY_MOBILE_KEY     ↔ type: mobile-key
	clientSideID  string // LAUNCHDARKLY_CLIENT_SIDE_ID ↔ type: client-side-id
}

// Run finds validatable snippets under cfg.SDKsDir and routes each through
// its language harness. A snippet is considered validatable when its
// frontmatter declares any one of:
//   - validation.runtime
//   - validation.entrypoint  (back-compat with the python first slice)
//
// EXAM-HELLO env vars (LAUNCHDARKLY_SDK_KEY, LAUNCHDARKLY_FLAG_KEY,
// LAUNCHDARKLY_MOBILE_KEY, LAUNCHDARKLY_CLIENT_SIDE_ID) are read from the
// caller's environment. Snippets that need a particular key declare an input
// of the matching type; the dispatcher refuses to run if a needed key is
// not set.
func Run(cfg Config) error {
	env := envInputs{
		sdkKey:       os.Getenv("LAUNCHDARKLY_SDK_KEY"),
		flagKey:      os.Getenv("LAUNCHDARKLY_FLAG_KEY"),
		mobileKey:    os.Getenv("LAUNCHDARKLY_MOBILE_KEY"),
		clientSideID: os.Getenv("LAUNCHDARKLY_CLIENT_SIDE_ID"),
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
		if !isValidatable(s) {
			continue
		}
		any = true
		if err := runOne(cfg, s, snippets, env); err != nil {
			return fmt.Errorf("validate %s: %w", id, err)
		}
	}
	if !any {
		return fmt.Errorf("no validatable snippets found (sdk=%q)", cfg.SDK)
	}
	return nil
}

func isValidatable(s *model.Snippet) bool {
	return s.Frontmatter.Validation.Runtime != "" || s.Frontmatter.Validation.Entrypoint != ""
}

func runOne(cfg Config, s *model.Snippet, all map[string]*model.Snippet, env envInputs) error {
	runtime := s.Frontmatter.Validation.Runtime
	if runtime == "" {
		runtime = s.CodeLang
	}
	if runtime == "" {
		return fmt.Errorf("snippet %q: cannot determine validator runtime (set validation.runtime or lang)", s.Frontmatter.ID)
	}

	runner, runnerDir, err := loadRunner(cfg.ValidatorsDir, runtime)
	if err != nil {
		return err
	}

	if err := requireEnvForInputs(s, all, env); err != nil {
		return err
	}

	stageDir, err := stageSnippet(s, all, env)
	if err != nil {
		return err
	}
	defer os.RemoveAll(stageDir)

	entrypoint := entrypointPath(s)
	fmt.Printf("--- validate %s (runtime=%s, entrypoint=%s) ---\n", s.Frontmatter.ID, runtime, entrypoint)

	switch runner.Mode {
	case "docker":
		return runDocker(cfg, runner, runnerDir, stageDir, entrypoint, env)
	case "native":
		return runNative(runnerDir, stageDir, entrypoint, env)
	default:
		return fmt.Errorf("validator runtime %q: unknown mode %q", runtime, runner.Mode)
	}
}

// entrypointPath returns the relative path the harness should invoke. If the
// snippet declares validation.entrypoint use that; otherwise fall back to
// the file: field (which is also where the body is staged).
func entrypointPath(s *model.Snippet) string {
	if s.Frontmatter.Validation.Entrypoint != "" {
		return s.Frontmatter.Validation.Entrypoint
	}
	return s.Frontmatter.File
}

// stageSnippet writes the snippet body and any companion bodies into a
// temp directory shaped exactly like the project the harness expects.
//
// Each snippet (entrypoint + companions) is rendered with runtime inputs and
// written at its `file:` path under stageDir. A snippet without a `file:`
// field is an error — we need to know where to put its body.
func stageSnippet(entry *model.Snippet, all map[string]*model.Snippet, env envInputs) (string, error) {
	stageDir, err := os.MkdirTemp("", "snippets-validate-")
	if err != nil {
		return "", err
	}

	if err := stageOne(stageDir, entry, env); err != nil {
		os.RemoveAll(stageDir)
		return "", err
	}
	for _, cid := range entry.Frontmatter.Validation.Companions {
		comp, ok := all[cid]
		if !ok {
			os.RemoveAll(stageDir)
			return "", fmt.Errorf("snippet %s: companion %q not found", entry.Frontmatter.ID, cid)
		}
		if err := stageOne(stageDir, comp, env); err != nil {
			os.RemoveAll(stageDir)
			return "", err
		}
	}

	// Python convention: validation.requirements becomes requirements.txt.
	// Other runtimes carry their dependency manifest as a companion snippet
	// (pom.xml, Cargo.toml, etc.).
	if req := entry.Frontmatter.Validation.Requirements; req != "" {
		if err := checkRequirements(req); err != nil {
			os.RemoveAll(stageDir)
			return "", err
		}
		if err := os.WriteFile(filepath.Join(stageDir, "requirements.txt"),
			[]byte(req+"\n"), 0o644); err != nil {
			os.RemoveAll(stageDir)
			return "", err
		}
	}
	return stageDir, nil
}

func stageOne(stageDir string, s *model.Snippet, env envInputs) error {
	rel := s.Frontmatter.File
	if rel == "" {
		return fmt.Errorf("snippet %s: frontmatter.file is required for staging", s.Frontmatter.ID)
	}
	if err := checkStagePath(rel); err != nil {
		return fmt.Errorf("snippet %s: %w", s.Frontmatter.ID, err)
	}
	inputs, err := runtimeInputs(s, env)
	if err != nil {
		return fmt.Errorf("snippet %s: %w", s.Frontmatter.ID, err)
	}
	nodes, err := render.Parse(s.CodeBody)
	if err != nil {
		return fmt.Errorf("snippet %s: %w", s.Frontmatter.ID, err)
	}
	body, err := render.RenderRuntime(nodes, inputs)
	if err != nil {
		return fmt.Errorf("snippet %s: %w", s.Frontmatter.ID, err)
	}
	dst := filepath.Join(stageDir, rel)
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	return os.WriteFile(dst, []byte(body), 0o644)
}

// checkStagePath rejects file paths that escape the staging directory.
func checkStagePath(rel string) error {
	clean := filepath.Clean(rel)
	if filepath.IsAbs(clean) {
		return fmt.Errorf("frontmatter.file %q must be relative", rel)
	}
	if clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
		return fmt.Errorf("frontmatter.file %q escapes staging directory", rel)
	}
	return nil
}

// runDocker builds the validator's Dockerfile and runs harness/run.sh inside
// the resulting container with the staged snippet bind-mounted at /snippet.
//
// Build context is the entire `validators/` tree so each Dockerfile can pull
// from `shared/` (the shared harness library) as well as its own
// `languages/<runtime>/` subtree.
func runDocker(cfg Config, runner *Runner, runnerDir, stageDir, entrypoint string, env envInputs) error {
	dockerfile := filepath.Join(runnerDir, "Dockerfile")
	if _, err := os.Stat(dockerfile); err != nil {
		return fmt.Errorf("validator Dockerfile not found at %s: %w", runnerDir, err)
	}
	tag, err := validatorImageTag(cfg.ValidatorsDir, runnerDir, runner.ImagePrefix)
	if err != nil {
		return err
	}
	build := exec.Command("docker", "build", "--quiet",
		"-f", dockerfile,
		"-t", tag,
		cfg.ValidatorsDir,
	)
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		return fmt.Errorf("docker build failed: %w", err)
	}
	args := []string{"run", "--rm",
		"-v", stageDir + ":/snippet:ro",
		"-e", "SNIPPET_ENTRYPOINT=" + entrypoint,
	}
	for _, kv := range envForRun(env) {
		args = append(args, "-e", kv)
	}
	args = append(args, tag)
	run := exec.Command("docker", args...)
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	if err := run.Run(); err != nil {
		return fmt.Errorf("snippet runtime validation failed: %w", err)
	}
	return nil
}

// runNative execs the harness's run.sh on the host with the staged snippet
// path passed as $SNIPPET_DIR. Used for runtimes whose toolchains can't run
// in a Linux container (iOS / xcodebuild) or are too heavy to dockerize for
// CI (Android emulator, Flutter).
func runNative(runnerDir, stageDir, entrypoint string, env envInputs) error {
	script := filepath.Join(runnerDir, "harness", "run.sh")
	if _, err := os.Stat(script); err != nil {
		return fmt.Errorf("native validator run.sh not found at %s: %w", script, err)
	}
	cmd := exec.Command("/bin/sh", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(),
		"SNIPPET_DIR="+stageDir,
		"SNIPPET_ENTRYPOINT="+entrypoint,
	)
	cmd.Env = append(cmd.Env, envForRun(env)...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("native validator failed: %w", err)
	}
	return nil
}

// envForRun returns the EXAM-HELLO env-var KEY=VALUE pairs that should be
// forwarded into the harness. Empty values are still forwarded (the harness
// can decide whether to require them); the per-snippet
// requireEnvForInputs check has already failed fast on missing values that
// the snippet actually needs.
func envForRun(env envInputs) []string {
	return []string{
		"LAUNCHDARKLY_SDK_KEY=" + env.sdkKey,
		"LAUNCHDARKLY_FLAG_KEY=" + env.flagKey,
		"LAUNCHDARKLY_MOBILE_KEY=" + env.mobileKey,
		"LAUNCHDARKLY_CLIENT_SIDE_ID=" + env.clientSideID,
	}
}

// requireEnvForInputs walks the entrypoint snippet AND its companions; for
// every input typed as one of the EXAM-HELLO key types, the corresponding
// env var must be set. This produces a clear error before a downstream pip-
// install or docker-build has wasted time.
func requireEnvForInputs(entry *model.Snippet, all map[string]*model.Snippet, env envInputs) error {
	check := func(s *model.Snippet) error {
		for name, in := range s.Frontmatter.Inputs {
			switch in.Type {
			case "flag-key":
				if env.flagKey == "" {
					return fmt.Errorf("snippet %s input %q (type=flag-key) requires LAUNCHDARKLY_FLAG_KEY to be set", s.Frontmatter.ID, name)
				}
			case "sdk-key":
				if env.sdkKey == "" {
					return fmt.Errorf("snippet %s input %q (type=sdk-key) requires LAUNCHDARKLY_SDK_KEY to be set", s.Frontmatter.ID, name)
				}
			case "mobile-key":
				if env.mobileKey == "" {
					return fmt.Errorf("snippet %s input %q (type=mobile-key) requires LAUNCHDARKLY_MOBILE_KEY to be set", s.Frontmatter.ID, name)
				}
			case "client-side-id":
				if env.clientSideID == "" {
					return fmt.Errorf("snippet %s input %q (type=client-side-id) requires LAUNCHDARKLY_CLIENT_SIDE_ID to be set", s.Frontmatter.ID, name)
				}
			}
		}
		return nil
	}
	if err := check(entry); err != nil {
		return err
	}
	for _, cid := range entry.Frontmatter.Validation.Companions {
		if comp, ok := all[cid]; ok {
			if err := check(comp); err != nil {
				return err
			}
		}
	}
	return nil
}

// runtimeInputs derives concrete values for every declared input.
//
// Inputs typed flag-key / sdk-key / mobile-key / client-side-id pull from
// the EXAM-HELLO env vars carried in `env`. Other inputs fall back to the
// snippet's own runtime-default. Declaring runtime-default for any of the
// EXAM-HELLO key types is an error: those values must always come from the
// caller's environment so real keys never end up committed.
func runtimeInputs(s *model.Snippet, env envInputs) (map[string]string, error) {
	out := map[string]string{}
	for name, in := range s.Frontmatter.Inputs {
		switch in.Type {
		case "flag-key":
			if in.RuntimeDefault != "" {
				return nil, fmt.Errorf("input %q (type=flag-key) must not declare runtime-default", name)
			}
			out[name] = env.flagKey
		case "sdk-key":
			if in.RuntimeDefault != "" {
				return nil, fmt.Errorf("input %q (type=sdk-key) must not declare runtime-default", name)
			}
			out[name] = env.sdkKey
		case "mobile-key":
			if in.RuntimeDefault != "" {
				return nil, fmt.Errorf("input %q (type=mobile-key) must not declare runtime-default", name)
			}
			out[name] = env.mobileKey
		case "client-side-id":
			if in.RuntimeDefault != "" {
				return nil, fmt.Errorf("input %q (type=client-side-id) must not declare runtime-default", name)
			}
			out[name] = env.clientSideID
		default:
			out[name] = in.RuntimeDefault
		}
	}
	return out, nil
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

// validatorImageTag produces a Docker tag that's a content hash of both the
// shared harness library AND the per-language validator directory. A change
// in either place forces a rebuild; concurrent validate runs against the
// same validator share the cached image.
func validatorImageTag(validatorsDir, runnerDir, prefix string) (string, error) {
	h := sha256.New()
	for _, sub := range []string{"shared", ""} {
		// "" means "the runner dir itself"; otherwise sub is rooted at validatorsDir.
		var root string
		if sub == "" {
			root = runnerDir
		} else {
			root = filepath.Join(validatorsDir, sub)
		}
		err := filepath.WalkDir(root, func(p string, d os.DirEntry, err error) error {
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
			rel, _ := filepath.Rel(validatorsDir, p)
			fmt.Fprintf(h, "%s\x00", rel)
			_, err = io.Copy(h, f)
			return err
		})
		if err != nil {
			return "", err
		}
	}
	return prefix + ":" + hex.EncodeToString(h.Sum(nil))[:16], nil
}
