package validate

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"maps"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/launchdarkly/sdk-meta/snippets/internal/model"
	"github.com/launchdarkly/sdk-meta/snippets/internal/render"
)

// Config controls a validator run.
type Config struct {
	SDKsFS        fs.FS  // sdks/ as an fs.FS (embedded or os.DirFS)
	ValidatorsDir string // path to validators/ (must be on disk — Docker COPY needs it)
	SDK           string // sdk id to validate (empty = all)
	Snippet       string // snippet id to validate (empty = all in the SDK)
	SnippetSkip   string // comma-separated snippet ids to skip (empty = none)
	Group         string // snippet group to validate (empty = all; e.g. "sdk-info" or "sdk-docs")
	Jobs          int    // max concurrent batch-harness invocations (0 = NumCPU); batch-mode validators only

	// ImageCache turns on cross-run Docker layer caching for `mode: docker`
	// validators. Empty (the default) keeps a plain `docker build`, which is
	// what local runs and fork CI use — no buildx, no external cache. When
	// set, the build runs through `docker buildx build` with a per-runtime
	// layer cache:
	//   "gha"                     — GitHub Actions cache (type=gha)
	//   "<registry>/<repo>/<img>" — a registry cache ref (type=registry)
	// CI sets this only on non-fork builds; see snippets-validate.yml.
	ImageCache string
}

// envInputs holds the environment-derived input values that get substituted
// into snippets at validation time. Each field maps to one EXAM-HELLO env
// var and to a snippet-input type.
type envInputs struct {
	sdkKey       string // LAUNCHDARKLY_SDK_KEY        ↔ type: sdk-key
	flagKey      string // LAUNCHDARKLY_FLAG_KEY       ↔ type: flag-key
	mobileKey    string // LAUNCHDARKLY_MOBILE_KEY     ↔ type: mobile-key
	clientSideID string // LAUNCHDARKLY_CLIENT_SIDE_ID ↔ type: client-side-id
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
// skipID reports whether id appears in the comma-separated skip list.
// Entries are trimmed so CI matrix values can wrap across lines.
func skipID(skipList, id string) bool {
	if skipList == "" {
		return false
	}
	for _, entry := range strings.Split(skipList, ",") {
		if strings.TrimSpace(entry) == id {
			return true
		}
	}
	return false
}

func Run(cfg Config) error {
	env := envInputs{
		sdkKey:       os.Getenv("LAUNCHDARKLY_SDK_KEY"),
		flagKey:      os.Getenv("LAUNCHDARKLY_FLAG_KEY"),
		mobileKey:    os.Getenv("LAUNCHDARKLY_MOBILE_KEY"),
		clientSideID: os.Getenv("LAUNCHDARKLY_CLIENT_SIDE_ID"),
	}

	snippets, err := model.LoadAll(cfg.SDKsFS)
	if err != nil {
		return err
	}

	// Resolve every (snippet, check) into a unit (runtime + runner +
	// assembled env) up front. Batch-capable runtimes are then grouped and
	// run a handful of times in warm workspaces; everything else keeps the
	// historical one-invocation-per-snippet path.
	var units []*resolvedUnit
	for _, id := range model.SortedIDs(snippets) {
		s := snippets[id]
		if cfg.SDK != "" && s.Frontmatter.SDK != cfg.SDK {
			continue
		}
		if cfg.Snippet != "" && s.Frontmatter.ID != cfg.Snippet {
			continue
		}
		if skipID(cfg.SnippetSkip, s.Frontmatter.ID) {
			continue
		}
		if cfg.Group != "" {
			// Snippet IDs are <sdk>/<group>/<name>; filter on the
			// middle segment so a CI row can validate just one group
			// (e.g. sdk-info installs vs sdk-docs reference fragments).
			parts := strings.Split(s.Frontmatter.ID, "/")
			if len(parts) < 2 || parts[1] != cfg.Group {
				continue
			}
		}
		if !isValidatable(s) {
			continue
		}
		checks := s.Frontmatter.Validation.EffectiveChecks()
		for i, check := range checks {
			u, err := resolveCheck(cfg, s, snippets, env, check, i, len(checks))
			if err != nil {
				return fmt.Errorf("validate %s: %w", id, err)
			}
			units = append(units, u)
		}
	}
	if len(units) == 0 {
		return fmt.Errorf("no validatable snippets found (sdk=%q, snippet=%q)", cfg.SDK, cfg.Snippet)
	}

	var batchUnits []*resolvedUnit
	for _, u := range units {
		if u.runner.Batch {
			batchUnits = append(batchUnits, u)
			continue
		}
		// Non-batch validator: stage and dispatch one snippet at a time,
		// exactly as before.
		if err := u.dispatch(); err != nil {
			return fmt.Errorf("validate %s: %w", u.snippet.Frontmatter.ID, err)
		}
	}
	if len(batchUnits) > 0 {
		if err := runBatches(cfg, batchUnits); err != nil {
			return err
		}
	}
	return nil
}

// isValidatable reports whether the validator should attempt to run a
// given snippet. Scaffolds (snippets that exist only to wrap other
// snippets' bodies) are explicitly excluded: their `{{ body }}` slot is
// unbound when run standalone, so they would always fail.
//
// Everything else with at least one effective check is in scope. The
// effective-check expansion handles the legacy single-validator shape by
// synthesizing one `kind: runtime` check from the top-level fields, so
// existing snippets keep validating without schema changes.
func isValidatable(s *model.Snippet) bool {
	if s.Frontmatter.Kind == "scaffold" {
		return false
	}
	return len(s.Frontmatter.Validation.EffectiveChecks()) > 0
}

// effectiveScaffold resolves a scaffold reference into the actual scaffold
// snippet. Returns (nil, nil) when scaffoldID is empty (plain snippet, no
// wrapping scaffold). Errors when the ID is non-empty but doesn't resolve
// to a snippet of kind=scaffold.
func effectiveScaffold(s *model.Snippet, scaffoldID string, all map[string]*model.Snippet) (*model.Snippet, error) {
	if scaffoldID == "" {
		return nil, nil
	}
	sc, ok := all[scaffoldID]
	if !ok {
		return nil, fmt.Errorf("snippet %s: scaffold %q not found", s.Frontmatter.ID, scaffoldID)
	}
	if sc.Frontmatter.Kind != "scaffold" {
		return nil, fmt.Errorf("snippet %s: validation.scaffold target %q has kind=%q (must be `scaffold`)",
			s.Frontmatter.ID, sc.Frontmatter.ID, sc.Frontmatter.Kind)
	}
	return sc, nil
}

// resolvedUnit is a single (snippet, check) with its runtime, runner, and
// run-time env resolved, ready to stage and validate. Staging is deferred
// (see stage) so batch mode can stage many units into one batch directory
// before invoking the harness once over the whole set.
type resolvedUnit struct {
	cfg        Config
	snippet    *model.Snippet
	all        map[string]*model.Snippet
	env        envInputs
	check      model.Check
	scaffold   *model.Snippet // nil when the snippet renders itself
	runtime    string
	runner     *Runner
	runnerDir  string
	entrypoint string            // path under the stage dir to the entry file
	extraEnv   map[string]string // build/dispatch env forwarded to the harness
	label      string            // "<id>" plus a check suffix, for log headers
}

// resolveCheck resolves one Check against the snippet into a resolvedUnit:
// it determines the runtime, loads the runner, enforces the env-input
// contract for runtime checks, and assembles the harness env. It performs
// no staging and runs no harness. Each Check carries its own (possibly
// inherited) scaffold, runtime, env, and placeholders, so a snippet with
// multiple checks can route each through a different harness branch.
func resolveCheck(cfg Config, s *model.Snippet, all map[string]*model.Snippet, env envInputs, check model.Check, checkIdx, checkTotal int) (*resolvedUnit, error) {
	scaffold, err := effectiveScaffold(s, check.Scaffold, all)
	if err != nil {
		return nil, err
	}

	// `eff` is the snippet whose frontmatter contributes file/inputs/
	// companions/requirements — scaffold when one is bound, the entry
	// snippet otherwise.
	eff := s
	if scaffold != nil {
		eff = scaffold
	}

	runtime := check.Runtime
	if runtime == "" {
		runtime = eff.Frontmatter.Validation.Runtime
	}
	if runtime == "" {
		// Fall back to the effective snippet's `lang:` field, matching
		// the documented contract on Validation.Runtime. (CodeLang — the
		// markdown fence's language tag — is a presentation hint and may
		// diverge from the author's declared `lang:`; using it here would
		// silently pick the wrong validator if they don't match.)
		runtime = eff.Frontmatter.Lang
	}
	if runtime == "" {
		return nil, fmt.Errorf("snippet %q (check %s): cannot determine validator runtime", s.Frontmatter.ID, check.Kind)
	}

	runner, runnerDir, err := loadRunner(cfg.ValidatorsDir, runtime)
	if err != nil {
		return nil, err
	}

	// Parse and typecheck modes don't need real LD credentials — the
	// scaffold's wrappee body is never executed and the harness's parser
	// invocation runs offline. Skip the env-input check for those kinds
	// so a CI row can validate parse/typecheck without provisioning a key.
	if check.Kind == "" || check.Kind == "runtime" {
		// Both the wrappee and the scaffold (when distinct) contribute to
		// the final body, so each one's typed inputs need their env values
		// satisfied. Companions are checked inside requireEnvForInputs.
		if err := requireEnvForInputs(s, all, env); err != nil {
			return nil, err
		}
		if eff != s {
			if err := requireEnvForInputs(eff, all, env); err != nil {
				return nil, err
			}
		}
	}

	entrypoint := check.Entrypoint
	if entrypoint == "" {
		entrypoint = entrypointPath(eff)
	}

	kindLabel := check.Kind
	if kindLabel == "" {
		kindLabel = "runtime"
	}
	checkSuffix := ""
	if checkTotal > 1 {
		checkSuffix = fmt.Sprintf(" check=%s [%d/%d]", kindLabel, checkIdx+1, checkTotal)
	} else if check.Kind != "" && check.Kind != "runtime" {
		checkSuffix = fmt.Sprintf(" check=%s", kindLabel)
	}

	// Forward per-snippet env, scaffold-supplied env, and the Check's own
	// env. Last writer wins (check overrides scaffold overrides wrappee
	// stays the rule that a scaffold can publish defaults a wrappee
	// (and individual check) can override).
	extraEnv := map[string]string{}
	if scaffold != nil {
		for k, v := range scaffold.Frontmatter.Validation.Env {
			extraEnv[k] = v
		}
	}
	for k, v := range s.Frontmatter.Validation.Env {
		extraEnv[k] = v
	}
	for k, v := range check.Env {
		extraEnv[k] = v
	}
	// Always announce the check kind to the harness so it can dispatch
	// even when the snippet didn't set any other env (default "runtime"
	// to preserve historical harness behavior).
	if _, ok := extraEnv["SNIPPET_CHECK"]; !ok {
		extraEnv["SNIPPET_CHECK"] = kindLabel
	}

	return &resolvedUnit{
		cfg:        cfg,
		snippet:    s,
		all:        all,
		env:        env,
		check:      check,
		scaffold:   scaffold,
		runtime:    runtime,
		runner:     runner,
		runnerDir:  runnerDir,
		entrypoint: entrypoint,
		extraEnv:   extraEnv,
		label:      s.Frontmatter.ID + checkSuffix,
	}, nil
}

// stage writes the unit's composed body (and companions/requirements) into
// a fresh temp directory and returns its path. The caller owns cleanup.
func (u *resolvedUnit) stage() (string, error) {
	return stageSnippetForCheck(u.snippet, u.all, u.env, u.check, u.scaffold)
}

// dispatch stages and validates a single non-batch unit, mirroring the
// historical one-invocation-per-snippet behavior.
func (u *resolvedUnit) dispatch() error {
	stageDir, err := u.stage()
	if err != nil {
		return err
	}
	defer os.RemoveAll(stageDir)

	if u.scaffold == nil {
		fmt.Printf("--- validate %s (runtime=%s, entrypoint=%s) ---\n", u.label, u.runtime, u.entrypoint)
	} else {
		fmt.Printf("--- validate %s (scaffold=%s, runtime=%s, entrypoint=%s) ---\n",
			u.label, u.scaffold.Frontmatter.ID, u.runtime, u.entrypoint)
	}

	switch u.runner.Mode {
	case "docker":
		if err := buildImage(u.cfg, u.runner, u.runnerDir, os.Stdout); err != nil {
			return err
		}
		tag, err := validatorImageTag(u.cfg.ValidatorsDir, u.runnerDir, u.runner.ImagePrefix)
		if err != nil {
			return err
		}
		return runContainer(tag, stageDir, u.entrypoint, u.env, u.extraEnv, os.Stdout)
	case "native":
		return runNative(u.runnerDir, stageDir, u.entrypoint, u.env, u.extraEnv, os.Stdout)
	default:
		return fmt.Errorf("validator runtime %q: unknown mode %q", u.runtime, u.runner.Mode)
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

// stageSnippetForCheck wraps stageSnippet with the Check's overrides
// applied. When a Check declares its own scaffold/placeholders/
// scaffold-inputs, those override (or merge into) the snippet's
// validation-block defaults for THIS check's staging only.
//
// The wrapper synthesizes a temporary snippet with the merged Validation
// block and delegates to stageSnippet, which keeps the staging logic
// unchanged for the legacy single-check path.
func stageSnippetForCheck(entry *model.Snippet, all map[string]*model.Snippet, env envInputs, check model.Check, scaffold *model.Snippet) (string, error) {
	synth := *entry
	v := entry.Frontmatter.Validation
	// Always set the resolved scaffold from the check (which already
	// merged the parent default in EffectiveChecks).
	v.Scaffold = check.Scaffold
	// Placeholders: merge parent + check (check wins on conflicts).
	if len(check.Placeholders) > 0 || len(entry.Frontmatter.Validation.Placeholders) > 0 {
		merged := map[string]string{}
		for k, val := range entry.Frontmatter.Validation.Placeholders {
			merged[k] = val
		}
		for k, val := range check.Placeholders {
			merged[k] = val
		}
		v.Placeholders = merged
	}
	if len(check.ScaffoldInputs) > 0 {
		merged := map[string]string{}
		for k, val := range entry.Frontmatter.Validation.ScaffoldInputs {
			merged[k] = val
		}
		for k, val := range check.ScaffoldInputs {
			merged[k] = val
		}
		v.ScaffoldInputs = merged
	}
	// Requirements and Companions are already resolved on the check by
	// EffectiveChecks (the check's own value, or the parent default).
	// Carry them onto the synthesized snippet so a per-check override
	// reaches staging; stageSnippet prefers the entry's values over the
	// scaffold's when set. Without this the override would be silently
	// dropped in favor of the scaffold's defaults.
	v.Requirements = check.Requirements
	v.Companions = check.Companions
	// Clear the multi-check list on the synthesized snippet so
	// stageSnippet's legacy code path doesn't recurse.
	v.Checks = nil
	synth.Frontmatter.Validation = v
	return stageSnippet(&synth, all, env)
}

// stageSnippet writes the entry snippet's body (or its scaffold-composed
// body) plus any companion bodies into a temp directory shaped exactly
// like the project the harness expects.
//
// Plain case: each snippet (entry + companions) is rendered with runtime
// inputs and written at its `file:` path under stageDir.
//
// Scaffold case: entry's body is rendered first; that string becomes the
// `body` input for the scaffold's render, and the scaffold's rendered
// output is staged at the scaffold's `file:` path. Companions and
// requirements come from the scaffold, unless the entry carries its own
// (a per-check override resolved onto the synthesized snippet) — entry
// values win when present.
func stageSnippet(entry *model.Snippet, all map[string]*model.Snippet, env envInputs) (string, error) {
	stageDir, err := os.MkdirTemp("", "snippets-validate-")
	if err != nil {
		return "", err
	}

	scaffold, err := effectiveScaffold(entry, entry.Frontmatter.Validation.Scaffold, all)
	if err != nil {
		os.RemoveAll(stageDir)
		return "", err
	}
	eff := entry
	if scaffold != nil {
		eff = scaffold
	}

	if eff == entry {
		// Plain (no scaffold): render entry with its own inputs.
		if err := stageRender(stageDir, entry, nil, env); err != nil {
			os.RemoveAll(stageDir)
			return "", err
		}
	} else {
		// Scaffolded: render the wrappee body first.
		wrappeeInputs, err := runtimeInputs(entry, env)
		if err != nil {
			os.RemoveAll(stageDir)
			return "", fmt.Errorf("snippet %s: %w", entry.Frontmatter.ID, err)
		}
		wrappeeNodes, err := render.Parse(entry.CodeBody)
		if err != nil {
			os.RemoveAll(stageDir)
			return "", fmt.Errorf("snippet %s: %w", entry.Frontmatter.ID, err)
		}
		wrappeeBody, err := render.RenderRuntime(wrappeeNodes, wrappeeInputs)
		if err != nil {
			os.RemoveAll(stageDir)
			return "", fmt.Errorf("snippet %s: %w", entry.Frontmatter.ID, err)
		}
		// Wrappee-declared placeholders apply to the wrappee body before
		// it's spliced into the scaffold. Scaffold-side substitution would
		// also catch the strings, but doing it here matches authors'
		// expectations: "this snippet's `'YOUR_SDK_KEY'` placeholder
		// becomes the env value."
		wrappeeBody, err = applyPlaceholders(wrappeeBody, entry.Frontmatter.Validation.Placeholders, env)
		if err != nil {
			os.RemoveAll(stageDir)
			return "", fmt.Errorf("snippet %s: %w", entry.Frontmatter.ID, err)
		}

		// Build the scaffold's input map: env-derived typed inputs from
		// the scaffold itself, overlaid with the wrappee's
		// scaffold-inputs, and finally the special `body` slot.
		scaffoldInputs, err := runtimeInputs(eff, env)
		if err != nil {
			os.RemoveAll(stageDir)
			return "", fmt.Errorf("scaffold %s: %w", eff.Frontmatter.ID, err)
		}
		maps.Copy(scaffoldInputs, entry.Frontmatter.Validation.ScaffoldInputs)
		scaffoldInputs["body"] = wrappeeBody

		if err := stageRender(stageDir, eff, scaffoldInputs, env); err != nil {
			os.RemoveAll(stageDir)
			return "", err
		}
	}

	// Plain-snippet placeholders: rewrite the staged file in place. (For
	// scaffolded snippets the substitution already happened on the
	// wrappee body, before scaffold composition.)
	if eff == entry && len(entry.Frontmatter.Validation.Placeholders) > 0 {
		if err := applyPlaceholdersToFile(filepath.Join(stageDir, entry.Frontmatter.File),
			entry.Frontmatter.Validation.Placeholders, env); err != nil {
			os.RemoveAll(stageDir)
			return "", fmt.Errorf("snippet %s: %w", entry.Frontmatter.ID, err)
		}
	}

	// Companions and requirements come from the effective snippet (the
	// scaffold when one is in use), unless the entry carries its own. A
	// per-check override (resolved onto the synthesized snippet in
	// stageSnippetForCheck) takes precedence over the scaffold's defaults;
	// in the legacy/no-scaffold case eff == entry, so this is a no-op.
	companions := entry.Frontmatter.Validation.Companions
	if len(companions) == 0 {
		companions = eff.Frontmatter.Validation.Companions
	}
	for _, cid := range companions {
		comp, ok := all[cid]
		if !ok {
			os.RemoveAll(stageDir)
			return "", fmt.Errorf("snippet %s: companion %q not found", eff.Frontmatter.ID, cid)
		}
		if err := stageRender(stageDir, comp, nil, env); err != nil {
			os.RemoveAll(stageDir)
			return "", err
		}
	}

	// Python convention: validation.requirements becomes requirements.txt.
	// Other runtimes carry their dependency manifest as a companion snippet
	// (pom.xml, Cargo.toml, etc.).
	req := entry.Frontmatter.Validation.Requirements
	if req == "" {
		req = eff.Frontmatter.Validation.Requirements
	}
	if req != "" {
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

// stageRender renders a snippet's body and writes it at its `file:` path
// under stageDir. If overrideInputs is non-nil it's used verbatim as the
// render-input map (used for scaffolds where the wrappee's body is
// supplied via the special `body` key); otherwise the snippet's own
// runtime inputs are derived from env.
func stageRender(stageDir string, s *model.Snippet, overrideInputs map[string]string, env envInputs) error {
	rel := s.Frontmatter.File
	if rel == "" {
		return fmt.Errorf("snippet %s: frontmatter.file is required for staging", s.Frontmatter.ID)
	}
	if err := checkStagePath(rel); err != nil {
		return fmt.Errorf("snippet %s: %w", s.Frontmatter.ID, err)
	}
	inputs := overrideInputs
	if inputs == nil {
		var err error
		inputs, err = runtimeInputs(s, env)
		if err != nil {
			return fmt.Errorf("snippet %s: %w", s.Frontmatter.ID, err)
		}
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

// buildImage builds the validator's Dockerfile, tagged with a content hash
// of the shared lib + runner dir so concurrent and repeat builds share the
// cached image. Idempotent: a second call with the same inputs is a cache
// hit. Build context is the entire `validators/` tree so each Dockerfile can
// pull from `shared/` as well as its own `languages/<runtime>/` subtree.
func buildImage(cfg Config, runner *Runner, runnerDir string, out io.Writer) error {
	dockerfile := filepath.Join(runnerDir, "Dockerfile")
	if _, err := os.Stat(dockerfile); err != nil {
		return fmt.Errorf("validator Dockerfile not found at %s: %w", runnerDir, err)
	}
	tag, err := validatorImageTag(cfg.ValidatorsDir, runnerDir, runner.ImagePrefix)
	if err != nil {
		return err
	}
	// `--progress=plain` keeps stdout tame (a one-line-per-step log
	// rather than the interactive multi-line redraws) while leaving
	// failure output visible — important for diagnosing apt/network
	// failures inside the build that --quiet would otherwise swallow.
	var build *exec.Cmd
	if cfg.ImageCache == "" {
		build = exec.Command("docker", "build", "--progress=plain",
			"-f", dockerfile,
			"-t", tag,
			cfg.ValidatorsDir,
		)
	} else {
		// buildx with a per-runtime layer cache. `--load` puts the built
		// image into the local daemon so the subsequent `docker run` finds
		// it. mode=min exports the final image's layers — these are
		// single-stage Dockerfiles, so the layer that does the expensive
		// work (compile the SDK, restore deps, warm the project) is part of
		// the final image and gets cached. A buildx builder that supports
		// cache export must already be configured (CI does this via
		// docker/setup-buildx-action).
		args := []string{"buildx", "build", "--progress=plain",
			"-f", dockerfile,
			"-t", tag,
			"--load",
		}
		args = append(args, cacheArgs(cfg.ImageCache, cacheScope(runnerDir))...)
		args = append(args, cfg.ValidatorsDir)
		build = exec.Command("docker", args...)
	}
	build.Stdout = out
	build.Stderr = out
	if err := build.Run(); err != nil {
		return fmt.Errorf("docker build failed: %w", err)
	}
	return nil
}

// cacheScope derives a stable per-runtime cache key from the runner
// directory (e.g. validators/languages/cpp-client -> "cpp-client"). Each
// validator gets its own cache so an edit to one doesn't invalidate the
// others. The result is lowercased and stripped to [a-z0-9._-] so it's a
// valid registry tag as well as a gha scope.
func cacheScope(runnerDir string) string {
	base := strings.ToLower(filepath.Base(runnerDir))
	var b strings.Builder
	for _, r := range base {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9', r == '.', r == '_', r == '-':
			b.WriteRune(r)
		default:
			b.WriteByte('-')
		}
	}
	return b.String()
}

// cacheArgs returns the --cache-from/--cache-to flags for buildx given the
// configured cache backend and a per-runtime scope. "gha" selects the
// GitHub Actions cache; anything else is treated as a registry ref prefix
// and the scope becomes the tag.
func cacheArgs(backend, scope string) []string {
	if backend == "gha" {
		return []string{
			"--cache-from", "type=gha,scope=" + scope,
			"--cache-to", "type=gha,scope=" + scope + ",mode=min",
		}
	}
	ref := backend + ":" + scope
	return []string{
		"--cache-from", "type=registry,ref=" + ref,
		"--cache-to", "type=registry,ref=" + ref + ",mode=min",
	}
}

// runContainer runs harness/run.sh inside the validator image with the
// staged snippet (single-snippet stage dir, or a batch dir) bind-mounted at
// /snippet. The image must already be built (see buildImage).
func runContainer(tag, stageDir, entrypoint string, env envInputs, extraEnv map[string]string, out io.Writer) error {
	args := []string{"run", "--rm",
		"-v", stageDir + ":/snippet:ro",
		"-e", "SNIPPET_ENTRYPOINT=" + entrypoint,
	}
	for _, kv := range envForRun(env) {
		args = append(args, "-e", kv)
	}
	for k, v := range extraEnv {
		args = append(args, "-e", k+"="+v)
	}
	args = append(args, tag)
	run := exec.Command("docker", args...)
	run.Stdout = out
	run.Stderr = out
	if err := run.Run(); err != nil {
		return fmt.Errorf("snippet runtime validation failed: %w", err)
	}
	return nil
}

// runNative execs the harness's run.sh on the host with the staged snippet
// path passed as $SNIPPET_DIR. Used for runtimes whose toolchains can't run
// in a Linux container (iOS / xcodebuild) or are too heavy to dockerize for
// CI (Android emulator, Flutter).
func runNative(runnerDir, stageDir, entrypoint string, env envInputs, extraEnv map[string]string, out io.Writer) error {
	script := filepath.Join(runnerDir, "harness", "run.sh")
	if _, err := os.Stat(script); err != nil {
		return fmt.Errorf("native validator run.sh not found at %s: %w", script, err)
	}
	cmd := exec.Command("/bin/sh", script)
	cmd.Stdout = out
	cmd.Stderr = out
	cmd.Env = append(os.Environ(),
		"SNIPPET_DIR="+stageDir,
		"SNIPPET_ENTRYPOINT="+entrypoint,
	)
	cmd.Env = append(cmd.Env, envForRun(env)...)
	for k, v := range extraEnv {
		cmd.Env = append(cmd.Env, k+"="+v)
	}
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
// env var must be set. Same check is applied to validation.placeholders.
// This produces a clear error before a downstream pip-install or
// docker-build has wasted time.
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
		for needle, envName := range s.Frontmatter.Validation.Placeholders {
			switch envName {
			case "LAUNCHDARKLY_SDK_KEY":
				if env.sdkKey == "" {
					return fmt.Errorf("snippet %s placeholder %q requires LAUNCHDARKLY_SDK_KEY to be set", s.Frontmatter.ID, needle)
				}
			case "LAUNCHDARKLY_FLAG_KEY":
				if env.flagKey == "" {
					return fmt.Errorf("snippet %s placeholder %q requires LAUNCHDARKLY_FLAG_KEY to be set", s.Frontmatter.ID, needle)
				}
			case "LAUNCHDARKLY_MOBILE_KEY":
				if env.mobileKey == "" {
					return fmt.Errorf("snippet %s placeholder %q requires LAUNCHDARKLY_MOBILE_KEY to be set", s.Frontmatter.ID, needle)
				}
			case "LAUNCHDARKLY_CLIENT_SIDE_ID":
				if env.clientSideID == "" {
					return fmt.Errorf("snippet %s placeholder %q requires LAUNCHDARKLY_CLIENT_SIDE_ID to be set", s.Frontmatter.ID, needle)
				}
			default:
				return fmt.Errorf("snippet %s placeholder %q maps to unknown env var %q", s.Frontmatter.ID, needle, envName)
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

// applyPlaceholders rewrites every literal occurrence of each key in the
// placeholders map with the corresponding env-var value. Returns an error
// if a placeholder maps to an env var that's empty or to a name outside
// the allow-list.
func applyPlaceholders(body string, placeholders map[string]string, env envInputs) (string, error) {
	if len(placeholders) == 0 {
		return body, nil
	}
	for needle, envName := range placeholders {
		var val string
		switch envName {
		case "LAUNCHDARKLY_SDK_KEY":
			val = env.sdkKey
		case "LAUNCHDARKLY_FLAG_KEY":
			val = env.flagKey
		case "LAUNCHDARKLY_MOBILE_KEY":
			val = env.mobileKey
		case "LAUNCHDARKLY_CLIENT_SIDE_ID":
			val = env.clientSideID
		default:
			return "", fmt.Errorf("placeholder %q maps to unknown env var %q (allowed: LAUNCHDARKLY_SDK_KEY, LAUNCHDARKLY_FLAG_KEY, LAUNCHDARKLY_MOBILE_KEY, LAUNCHDARKLY_CLIENT_SIDE_ID)", needle, envName)
		}
		if val == "" {
			return "", fmt.Errorf("placeholder %q requires %s to be set", needle, envName)
		}
		if !strings.Contains(body, needle) {
			return "", fmt.Errorf("placeholder %q not found in snippet body", needle)
		}
		body = strings.ReplaceAll(body, needle, val)
	}
	return body, nil
}

// applyPlaceholdersToFile reads the file, applies placeholders, and writes
// the result back. Used for plain (non-scaffolded) snippets.
func applyPlaceholdersToFile(path string, placeholders map[string]string, env envInputs) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	out, err := applyPlaceholders(string(raw), placeholders, env)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(out), 0o644)
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
