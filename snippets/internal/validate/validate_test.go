package validate

import (
	"strings"
	"testing"

	"github.com/launchdarkly/sdk-meta/snippets/internal/model"
)

// Regression for review #2: snippet author-controlled values may not escape
// the staging directory.
func TestCheckEntrypointRejectsTraversal(t *testing.T) {
	bad := []string{
		"../etc/passwd",
		"../../home/x/.ssh/authorized_keys",
		"a/b/c.py",
		"./main.py",
		".",
		"..",
	}
	for _, e := range bad {
		if err := checkEntrypoint(e); err == nil {
			t.Errorf("checkEntrypoint(%q): expected error", e)
		}
	}
	good := []string{"main.py", "app.py", "snippet_test.py"}
	for _, e := range good {
		if err := checkEntrypoint(e); err != nil {
			t.Errorf("checkEntrypoint(%q): unexpected error %v", e, err)
		}
	}
	// Empty is allowed (no validation entrypoint declared).
	if err := checkEntrypoint(""); err != nil {
		t.Errorf("empty entrypoint should be allowed: %v", err)
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

// Regression for review #7: sdk-key inputs must come from the env, not from
// a snippet's runtime-default.
func TestRuntimeInputs(t *testing.T) {
	s := &model.Snippet{
		Frontmatter: model.Frontmatter{
			Inputs: map[string]model.Input{
				"apiKey":     {Type: "sdk-key"},
				"featureKey": {Type: "flag-key"},
				"version":    {Type: "string", RuntimeDefault: "1.2.3"},
			},
		},
	}
	got, err := runtimeInputs(s, "real-sdk-key", "real-flag-key")
	if err != nil {
		t.Fatal(err)
	}
	if got["apiKey"] != "real-sdk-key" {
		t.Errorf("apiKey: got %q want real-sdk-key", got["apiKey"])
	}
	if got["featureKey"] != "real-flag-key" {
		t.Errorf("featureKey: got %q want real-flag-key", got["featureKey"])
	}
	if got["version"] != "1.2.3" {
		t.Errorf("version: got %q want 1.2.3", got["version"])
	}
}

func TestRuntimeInputsRejectsRuntimeDefaultOnKeys(t *testing.T) {
	for _, kind := range []string{"sdk-key", "flag-key"} {
		s := &model.Snippet{
			Frontmatter: model.Frontmatter{
				Inputs: map[string]model.Input{
					"k": {Type: kind, RuntimeDefault: "should-not-be-here"},
				},
			},
		}
		_, err := runtimeInputs(s, "x", "y")
		if err == nil || !strings.Contains(err.Error(), "runtime-default") {
			t.Errorf("%s: want runtime-default rejection, got %v", kind, err)
		}
	}
}
