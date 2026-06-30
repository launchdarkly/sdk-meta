package validate

import (
	"strings"
	"testing"
)

func TestCacheScope(t *testing.T) {
	cases := map[string]string{
		"validators/languages/cpp-client":        "cpp-client",
		"validators/languages/haskell-server-v3": "haskell-server-v3",
		"validators/languages/dotnet-server":     "dotnet-server",
		// Anything outside [a-z0-9._-] is replaced so the scope is a valid
		// registry tag and gha scope.
		"validators/languages/Weird Name!": "weird-name-",
	}
	for in, want := range cases {
		if got := cacheScope(in); got != want {
			t.Errorf("cacheScope(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestCacheArgs(t *testing.T) {
	// gha backend.
	gha := cacheArgs("gha", "rust")
	joined := strings.Join(gha, " ")
	if !strings.Contains(joined, "type=gha,scope=rust") {
		t.Errorf("gha cache-from missing scope: %v", gha)
	}
	if !strings.Contains(joined, "type=gha,scope=rust,mode=min") {
		t.Errorf("gha cache-to missing scope/mode: %v", gha)
	}

	// registry backend: scope becomes the tag on the ref prefix.
	reg := cacheArgs("ghcr.io/launchdarkly/sdk-meta/snippet-cache", "cpp-server")
	joined = strings.Join(reg, " ")
	if !strings.Contains(joined, "type=registry,ref=ghcr.io/launchdarkly/sdk-meta/snippet-cache:cpp-server") {
		t.Errorf("registry cache ref wrong: %v", reg)
	}
	if !strings.Contains(joined, ",mode=min") {
		t.Errorf("registry cache-to missing mode: %v", reg)
	}

	// Both backends emit one --cache-from and one --cache-to.
	for _, args := range [][]string{gha, reg} {
		from, to := 0, 0
		for _, a := range args {
			if a == "--cache-from" {
				from++
			}
			if a == "--cache-to" {
				to++
			}
		}
		if from != 1 || to != 1 {
			t.Errorf("want 1 cache-from + 1 cache-to, got %d/%d in %v", from, to, args)
		}
	}
}
