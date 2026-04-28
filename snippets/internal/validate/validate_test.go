package validate

import (
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
