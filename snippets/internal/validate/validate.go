package validate

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
	inputs := runtimeInputs(s, flagKey)
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

	imageTag := "sdk-snippets/python-validator:dev"
	build := exec.Command("docker", "build", "--quiet", "-t", imageTag, validatorDir)
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
		imageTag,
	)
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	if err := run.Run(); err != nil {
		return fmt.Errorf("snippet runtime validation failed: %w", err)
	}
	return nil
}

// runtimeInputs derives concrete values for every declared input.
// Inputs typed as `flag-key` use the LAUNCHDARKLY_FLAG_KEY env value. Inputs
// typed as `sdk-key` use the LAUNCHDARKLY_SDK_KEY env value (the snippet's
// rendered source needs the literal key when it's interpolated, e.g. in the
// `Run` shell command — but the Python `main.py` reads it from the env and
// never has the key in its source). Other inputs fall back to the snippet's
// own runtime-default.
func runtimeInputs(s *model.Snippet, flagKey string) map[string]string {
	out := map[string]string{}
	for name, in := range s.Frontmatter.Inputs {
		switch in.Type {
		case "flag-key":
			out[name] = flagKey
		default:
			out[name] = in.RuntimeDefault
		}
	}
	return out
}
