package validate

import (
	"fmt"
	"io/fs"
	"sort"
	"strings"

	"github.com/launchdarkly/sdk-meta/snippets/internal/model"
)

// ImageTag returns the deterministic Docker image tag the validator will
// build for a runtime (`<image-prefix>:<content-hash-16-hex>`). Exposed so
// CI can pre-build the image out-of-band (e.g. via `docker buildx build`
// with a remote layer cache) with the same tag the Go validator's
// subsequent `docker build` will look up.
//
// Returns an empty string with no error for runtimes whose runner.yaml
// declares `mode: native` — no image is built for those.
func ImageTag(validatorsDir, runtime string) (string, error) {
	runner, runnerDir, err := loadRunner(validatorsDir, runtime)
	if err != nil {
		return "", err
	}
	if runner.Mode != "docker" {
		return "", nil
	}
	return validatorImageTag(validatorsDir, runnerDir, runner.ImagePrefix)
}

// ListRunners returns the deduplicated set of validator runtimes used by
// the validatable snippets matching the given sdk + group filters. Exposed
// so CI's per-row pre-build step knows exactly which Dockerfiles to build
// (without enumerating every validator that happens to exist on disk).
//
// The set is computed by walking every bound snippet's effective checks
// the same way Run() does: a snippet bound via `validation.scaffold:` uses
// its scaffold's runtime; a Check with an explicit `runtime:` overrides
// the parent Validation's; a missing runtime falls back to the effective
// snippet's `lang:` field.
func ListRunners(sdksFS fs.FS, validatorsDir, sdk, group string) ([]string, error) {
	snippets, err := model.LoadAll(sdksFS)
	if err != nil {
		return nil, err
	}
	seen := map[string]struct{}{}
	for _, id := range model.SortedIDs(snippets) {
		s := snippets[id]
		if sdk != "" && s.Frontmatter.SDK != sdk {
			continue
		}
		if group != "" {
			parts := strings.Split(s.Frontmatter.ID, "/")
			if len(parts) < 2 || parts[1] != group {
				continue
			}
		}
		if !isValidatable(s) {
			continue
		}
		for _, check := range s.Frontmatter.Validation.EffectiveChecks() {
			runtime, err := effectiveRuntime(s, snippets, check)
			if err != nil {
				return nil, fmt.Errorf("snippet %s: %w", s.Frontmatter.ID, err)
			}
			seen[runtime] = struct{}{}
		}
	}
	out := make([]string, 0, len(seen))
	for r := range seen {
		out = append(out, r)
	}
	sort.Strings(out)
	return out, nil
}

// effectiveRuntime resolves the same runtime-selection chain that
// runOneCheck uses: explicit check.Runtime > effective snippet's
// Validation.Runtime > effective snippet's lang. The "effective snippet"
// is the scaffold when one is bound, the entry snippet otherwise.
func effectiveRuntime(s *model.Snippet, all map[string]*model.Snippet, check model.Check) (string, error) {
	scaffold, err := effectiveScaffold(s, check.Scaffold, all)
	if err != nil {
		return "", err
	}
	eff := s
	if scaffold != nil {
		eff = scaffold
	}
	runtime := check.Runtime
	if runtime == "" {
		runtime = eff.Frontmatter.Validation.Runtime
	}
	if runtime == "" {
		runtime = eff.Frontmatter.Lang
	}
	if runtime == "" {
		return "", fmt.Errorf("cannot determine validator runtime for check %q", check.Kind)
	}
	return runtime, nil
}
