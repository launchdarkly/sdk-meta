package validate

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/launchdarkly/sdk-meta/snippets/internal/model"
)

// Regression for review #2: snippet `file:` paths may not escape the staging
// directory. The check moved from a dedicated entrypoint guard to a stage-
// path guard during the multi-file refactor.
func TestCheckStagePathRejectsTraversal(t *testing.T) {
	bad := []string{
		"../etc/passwd",
		"../../home/x/.ssh/authorized_keys",
		"/etc/passwd",
		"..",
	}
	for _, e := range bad {
		if err := checkStagePath(e); err == nil {
			t.Errorf("checkStagePath(%q): expected error", e)
		}
	}
	good := []string{
		"main.py",
		"src/main.rs",
		"src/main/java/com/launchdarkly/HelloLD.java",
		"./main.py", // filepath.Clean normalizes to main.py
	}
	for _, e := range good {
		if err := checkStagePath(e); err != nil {
			t.Errorf("checkStagePath(%q): unexpected error %v", e, err)
		}
	}
}

// Regression for review #8: pip flag injection via requirements.txt.
func TestCheckRequirementsRejectsPipFlags(t *testing.T) {
	bad := []string{
		"--extra-index-url=https://evil.example/pypi",
		"--index-url https://evil.example/pypi",
		"-r other-requirements.txt",
		"requests\n--extra-index-url=https://evil.example/pypi",
	}
	for _, r := range bad {
		err := checkRequirements(r)
		if err == nil || !strings.Contains(err.Error(), "pip flags") {
			t.Errorf("checkRequirements(%q): want pip-flag error, got %v", r, err)
		}
	}
	good := []string{
		"launchdarkly-server-sdk",
		"launchdarkly-server-sdk==9.2.0",
		"requests\nlaunchdarkly-server-sdk",
	}
	for _, r := range good {
		if err := checkRequirements(r); err != nil {
			t.Errorf("checkRequirements(%q): unexpected error %v", r, err)
		}
	}
}

// All four EXAM-HELLO key types pull from env, not from runtime-default.
func TestRuntimeInputs(t *testing.T) {
	env := envInputs{
		sdkKey:       "real-sdk-key",
		flagKey:      "real-flag-key",
		mobileKey:    "real-mobile-key",
		clientSideID: "real-client-side-id",
	}
	s := &model.Snippet{
		Frontmatter: model.Frontmatter{
			Inputs: map[string]model.Input{
				"apiKey":         {Type: "sdk-key"},
				"featureKey":     {Type: "flag-key"},
				"mobileKey":      {Type: "mobile-key"},
				"clientSideId":   {Type: "client-side-id"},
				"version":        {Type: "string", RuntimeDefault: "1.2.3"},
			},
		},
	}
	got, err := runtimeInputs(s, env)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"apiKey":       "real-sdk-key",
		"featureKey":   "real-flag-key",
		"mobileKey":    "real-mobile-key",
		"clientSideId": "real-client-side-id",
		"version":      "1.2.3",
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("input %q: got %q want %q", k, got[k], v)
		}
	}
}

func TestRuntimeInputsRejectsRuntimeDefaultOnKeys(t *testing.T) {
	for _, kind := range []string{"sdk-key", "flag-key", "mobile-key", "client-side-id"} {
		s := &model.Snippet{
			Frontmatter: model.Frontmatter{
				Inputs: map[string]model.Input{
					"k": {Type: kind, RuntimeDefault: "should-not-be-here"},
				},
			},
		}
		_, err := runtimeInputs(s, envInputs{})
		if err == nil || !strings.Contains(err.Error(), "runtime-default") {
			t.Errorf("%s: want runtime-default rejection, got %v", kind, err)
		}
	}
}

// stageSnippet composes a wrappee into a scaffold's `{{ body }}` slot.
// The staged file is the scaffold's `file:` path; the bytes are the
// scaffold body with the wrappee body inlined.
func TestStageSnippet_ScaffoldComposition(t *testing.T) {
	wrappee := &model.Snippet{
		Path: "test/wrappee.snippet.md",
		Frontmatter: model.Frontmatter{
			ID:   "test-sdk/docs/eval",
			SDK:  "test-sdk",
			Kind: "reference",
			Lang: "python",
			Validation: model.Validation{
				Scaffold: "test-sdk/scaffolds/with-test-data",
			},
		},
		CodeBody: `flag_value = client.variation("your.feature.key", context, False)`,
	}
	scaffold := &model.Snippet{
		Path: "test/scaffold.snippet.md",
		Frontmatter: model.Frontmatter{
			ID:   "test-sdk/scaffolds/with-test-data",
			SDK:  "test-sdk",
			Kind: "scaffold",
			Lang: "python",
			File: "main.py",
			Inputs: map[string]model.Input{
				"body": {Type: "string", Description: "Wrappee body"},
			},
			Validation: model.Validation{
				Runtime:      "python",
				Entrypoint:   "main.py",
				Requirements: "launchdarkly-server-sdk",
			},
		},
		CodeBody: "import ldclient\n# setup ...\n{{ body }}\nprint(\"feature flag evaluates to\", flag_value)",
	}
	all := map[string]*model.Snippet{
		wrappee.Frontmatter.ID:  wrappee,
		scaffold.Frontmatter.ID: scaffold,
	}

	stageDir, err := stageSnippet(wrappee, all, envInputs{})
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(stageDir)

	// Composed body lives at the scaffold's `file:` path, not the wrappee's.
	out, err := os.ReadFile(filepath.Join(stageDir, "main.py"))
	if err != nil {
		t.Fatal(err)
	}
	got := string(out)
	if !strings.Contains(got, `flag_value = client.variation("your.feature.key", context, False)`) {
		t.Errorf("staged body missing wrappee fragment:\n%s", got)
	}
	if !strings.Contains(got, "import ldclient") {
		t.Errorf("staged body missing scaffold prologue:\n%s", got)
	}
	if !strings.Contains(got, "feature flag evaluates to") {
		t.Errorf("staged body missing scaffold epilogue:\n%s", got)
	}

	// requirements.txt comes from the scaffold.
	req, err := os.ReadFile(filepath.Join(stageDir, "requirements.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(req), "launchdarkly-server-sdk") {
		t.Errorf("requirements.txt missing scaffold dep: %q", string(req))
	}
}

// scaffold-inputs from the wrappee override env-derived defaults on the
// scaffold's render.
func TestStageSnippet_ScaffoldInputsOverride(t *testing.T) {
	wrappee := &model.Snippet{
		Frontmatter: model.Frontmatter{
			ID:   "test-sdk/docs/eval",
			SDK:  "test-sdk",
			Kind: "reference",
			Lang: "python",
			Validation: model.Validation{
				Scaffold: "test-sdk/scaffolds/with-flag",
				ScaffoldInputs: map[string]string{
					"flagName": "your.feature.key",
				},
			},
		},
		CodeBody: `# wrappee body`,
	}
	scaffold := &model.Snippet{
		Frontmatter: model.Frontmatter{
			ID:   "test-sdk/scaffolds/with-flag",
			SDK:  "test-sdk",
			Kind: "scaffold",
			Lang: "python",
			File: "main.py",
			Inputs: map[string]model.Input{
				"body":     {Type: "string"},
				"flagName": {Type: "string", RuntimeDefault: "default.key"},
			},
			Validation: model.Validation{Runtime: "python", Entrypoint: "main.py"},
		},
		CodeBody: "FLAG = \"{{ flagName }}\"\n{{ body }}",
	}
	all := map[string]*model.Snippet{
		wrappee.Frontmatter.ID:  wrappee,
		scaffold.Frontmatter.ID: scaffold,
	}
	stageDir, err := stageSnippet(wrappee, all, envInputs{})
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(stageDir)
	out, _ := os.ReadFile(filepath.Join(stageDir, "main.py"))
	got := string(out)
	if !strings.Contains(got, `FLAG = "your.feature.key"`) {
		t.Errorf("scaffold-inputs override did not apply:\n%s", got)
	}
}

// isValidatable skips scaffolds even when they declare runtime/entrypoint.
// A standalone scaffold run would have an unbound `{{ body }}` slot.
func TestIsValidatable_SkipsScaffolds(t *testing.T) {
	s := &model.Snippet{
		Frontmatter: model.Frontmatter{
			Kind: "scaffold",
			Validation: model.Validation{
				Runtime:    "python",
				Entrypoint: "main.py",
			},
		},
	}
	if isValidatable(s) {
		t.Error("scaffolds must not be validatable on their own")
	}

	// Compare with a plain validatable snippet.
	plain := &model.Snippet{
		Frontmatter: model.Frontmatter{
			Kind: "hello-world",
			Validation: model.Validation{
				Runtime:    "python",
				Entrypoint: "main.py",
			},
		},
	}
	if !isValidatable(plain) {
		t.Error("plain validatable snippet should be validatable")
	}

	// And a scaffold-bound snippet (no runtime of its own) IS validatable.
	scaffolded := &model.Snippet{
		Frontmatter: model.Frontmatter{
			Kind: "reference",
			Validation: model.Validation{
				Scaffold: "test-sdk/scaffolds/with-test-data",
			},
		},
	}
	if !isValidatable(scaffolded) {
		t.Error("scaffold-bound snippet should be validatable via the scaffold")
	}
}

// effectiveScaffold errors loudly when the scaffold ID can't be resolved
// or when the target isn't actually a scaffold.
func TestEffectiveScaffold_RejectsBadScaffold(t *testing.T) {
	s := &model.Snippet{
		Frontmatter: model.Frontmatter{
			ID:         "test/wrappee",
			Validation: model.Validation{Scaffold: "nonexistent"},
		},
	}
	if _, err := effectiveScaffold(s, "nonexistent", map[string]*model.Snippet{}); err == nil ||
		!strings.Contains(err.Error(), "scaffold") {
		t.Errorf("missing scaffold: want error mentioning scaffold, got %v", err)
	}

	notScaffold := &model.Snippet{
		Frontmatter: model.Frontmatter{ID: "test/init", Kind: "hello-world"},
	}
	all := map[string]*model.Snippet{notScaffold.Frontmatter.ID: notScaffold}
	if _, err := effectiveScaffold(s, notScaffold.Frontmatter.ID, all); err == nil ||
		!strings.Contains(err.Error(), "kind=") {
		t.Errorf("non-scaffold target: want kind error, got %v", err)
	}
}
