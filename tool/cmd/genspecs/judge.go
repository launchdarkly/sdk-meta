package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var httpClient = &http.Client{Timeout: 5 * time.Minute}

func runJudge(args []string) error {
	fs := flag.NewFlagSet("judge", flag.ExitOnError)
	specsJSON := fs.String("specs-json", "products/specs.json", "Path to products/specs.json (input).")
	harnessJSON := fs.String("harness-json", "products/harness_signals.json", "Path to products/harness_signals.json (input).")
	featuresJSON := fs.String("features-json", "products/features.json", "Path to products/features.json (input).")
	featureInfoJSON := fs.String("feature-info-json", "products/feature_info.json", "Path to products/feature_info.json (input).")
	typesJSON := fs.String("types-json", "products/types.json", "Path to products/types.json (input).")
	namesJSON := fs.String("names-json", "products/names.json", "Path to products/names.json (input).")
	languagesJSON := fs.String("languages-json", "products/languages.json", "Path to products/languages.json (input).")
	reposJSON := fs.String("repos-json", "products/repos.json", "Path to products/repos.json (input).")
	specsRepoPath := fs.String("specs-repo", filepath.Join(defaultReposRoot(), "sdk-specs"), "Path to local sdk-specs checkout (used to read spec READMEs).")
	root := fs.String("sdk-repos-root", defaultReposRoot(), "Directory under which SDK repos live.")
	out := fs.String("out", "products/spec_support.json", "Output path for spec_support.json.")
	provider := fs.String("provider", defaultProvider(), "Judge provider: 'bedrock' (uses AWS_BEARER_TOKEN_BEDROCK), 'anthropic' (uses ANTHROPIC_API_KEY), or 'noop'.")
	model := fs.String("model", "", "Model id. Defaults depend on provider: a Bedrock model id when provider='bedrock', the Anthropic native id when provider='anthropic'.")
	concurrency := fs.Int("concurrency", 4, "Maximum concurrent LLM requests.")
	cacheDir := fs.String("cache-dir", filepath.Join("tool", "specs", ".judge-cache"), "Directory for the per-cell prompt-pack cache.")
	maxCells := fs.Int("max-cells", 0, "If > 0, only judge this many LLM-bound cells (the rest stay 'unknown'). Useful for incremental runs.")
	if err := fs.Parse(args); err != nil {
		return err
	}

	specsProduct, err := loadJSON[SpecsProduct](*specsJSON)
	if err != nil {
		return fmt.Errorf("loading specs.json: %w", err)
	}
	harnessProduct, err := loadJSON[HarnessProduct](*harnessJSON)
	if err != nil {
		return fmt.Errorf("loading harness_signals.json: %w", err)
	}
	features, err := loadJSON[map[string]map[string]FeatureSupportLite](*featuresJSON)
	if err != nil {
		return fmt.Errorf("loading features.json: %w", err)
	}
	featureInfo, err := loadJSON[map[string]FeatureInfoLite](*featureInfoJSON)
	if err != nil {
		return fmt.Errorf("loading feature_info.json: %w", err)
	}
	types, err := loadJSON[map[string]string](*typesJSON)
	if err != nil {
		return fmt.Errorf("loading types.json: %w", err)
	}
	names, err := loadJSON[map[string]string](*namesJSON)
	if err != nil {
		return fmt.Errorf("loading names.json: %w", err)
	}
	languages, err := loadJSON[map[string][]string](*languagesJSON)
	if err != nil {
		return fmt.Errorf("loading languages.json: %w", err)
	}
	repos, err := readReposJSONFull(*reposJSON)
	if err != nil {
		return fmt.Errorf("loading repos.json: %w", err)
	}

	judge, err := buildJudge(*provider, *model)
	if err != nil {
		return fmt.Errorf("building judge: %w", err)
	}

	if err := os.MkdirAll(*cacheDir, 0o755); err != nil {
		return fmt.Errorf("creating cache dir: %w", err)
	}

	// Iterate sdks and specs in stable order.
	sdkIDs := sortedKeys(types)
	specIDs := sortedKeys(specsProduct.Specs)

	result := SpecSupportProduct{
		GeneratedAt:   time.Now().UTC().Format(time.RFC3339),
		SpecsCommit:   specsProduct.SourceCommit,
		HarnessCommit: harnessProduct.Harness.Commit,
		Model:         judge.Model(),
		PromptVersion: promptVersion,
		SDKs:          map[string]map[string]Cell{},
	}

	// Build the work list: every (sdk, spec) cell that survives the applies-to filter.
	type job struct {
		sdkID, specID string
	}
	var work []job
	for _, sdkID := range sdkIDs {
		sdkType := types[sdkID]
		row := map[string]Cell{}
		for _, specID := range specIDs {
			spec := specsProduct.Specs[specID]
			cell, applicable := appliesToCell(spec, sdkID, sdkType)
			if !applicable {
				row[specID] = cell
				continue
			}
			// Reserve an unknown slot; judge step may overwrite.
			row[specID] = Cell{State: StateUnknown, Source: SourceJudgeFailed, Confidence: ConfidenceLow}
			work = append(work, job{sdkID, specID})
		}
		result.SDKs[sdkID] = row
	}

	if *maxCells > 0 && len(work) > *maxCells {
		work = work[:*maxCells]
		fmt.Fprintf(os.Stderr, "Capping work to first %d cells per --max-cells.\n", *maxCells)
	}

	fmt.Fprintf(os.Stderr, "Judging %d (SDK, spec) cells with %s\n", len(work), judge.Describe())

	cache := newCellCache(*cacheDir)
	sem := make(chan struct{}, *concurrency)
	var mu sync.Mutex
	var wg sync.WaitGroup
	progress := 0

	for _, j := range work {
		wg.Add(1)
		j := j
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			pack := buildPromptPack(packInputs{
				spec:            specsProduct.Specs[j.specID],
				sdkID:           j.sdkID,
				sdkName:         names[j.sdkID],
				sdkType:         types[j.sdkID],
				languages:       languages[j.sdkID],
				features:        features[j.sdkID],
				featureInfo:     featureInfo,
				harness:         harnessProduct.Harness,
				sdkHarness:      harnessProduct.SDKs[j.sdkID],
				repoGitHub:      repos[j.sdkID],
				specsRepoPath:   *specsRepoPath,
				sdkRepoRoot:     *root,
				specsCommit:     specsProduct.SourceCommit,
				harnessCommit:   harnessProduct.Harness.Commit,
				promptVersion:   promptVersion,
				modelIdentifier: judge.Model(),
			})

			cell, fromCache, err := evaluate(context.Background(), judge, pack, cache)

			mu.Lock()
			defer mu.Unlock()
			progress++
			if err != nil {
				fmt.Fprintf(os.Stderr, "[%d/%d] %-32s %-12s ERROR: %v\n", progress, len(work), j.sdkID, j.specID, err)
				cell = Cell{
					State:      StateUnknown,
					Source:     SourceJudgeFailed,
					Confidence: ConfidenceLow,
					Rationale:  truncate("judge failed: "+err.Error(), 240),
				}
			} else {
				marker := "*"
				if fromCache {
					marker = "C"
				}
				fmt.Fprintf(os.Stderr, "[%d/%d] %s %-32s %-12s %s (%s)\n", progress, len(work), marker, j.sdkID, j.specID, cell.State, cell.Confidence)
			}
			result.SDKs[j.sdkID][j.specID] = cell
		}()
	}
	wg.Wait()

	return writeJSON(*out, result)
}

func evaluate(ctx context.Context, judge Judge, pack PromptPack, cache *cellCache) (Cell, bool, error) {
	now := time.Now().UTC().Format(time.RFC3339)
	hash := pack.Hash()
	if cached, ok := cache.Get(hash); ok {
		cell := cached
		cell.JudgedAt = &now
		// JudgedAgainst stays as the cached value, so reviewers can see exactly
		// which inputs the cell was judged against.
		return cell, true, nil
	}

	resp, err := judge.Judge(ctx, pack)
	if err != nil {
		return Cell{}, false, err
	}

	state := normalizeState(resp.State)
	if state == "" {
		return Cell{}, false, fmt.Errorf("judge returned unknown state %q", resp.State)
	}

	cell := Cell{
		State:      state,
		Confidence: normalizeConfidence(resp.Confidence),
		Source:     SourceLLMJudge,
		Rationale:  truncate(strings.TrimSpace(resp.Rationale), 480),
		Evidence:   resp.Evidence,
		JudgedAt:   &now,
		JudgedAgainst: &JudgedAgainst{
			SpecsCommit:   pack.SpecsCommit,
			HarnessCommit: pack.HarnessCommit,
			SDKCommit:     pack.SDKCommit,
			Model:         pack.ModelIdentifier,
			PromptVersion: pack.PromptVersion,
		},
	}
	if resp.NotesForHuman != "" {
		n := resp.NotesForHuman
		cell.NotesForHuman = &n
	}
	cache.Put(hash, cell)
	return cell, false, nil
}

func appliesToCell(spec Spec, sdkID, sdkType string) (Cell, bool) {
	// Empty applies-to means foundational; never short-circuit those.
	if len(spec.AppliesTo) == 0 {
		return Cell{}, true
	}
	canonical := canonicalSDKKind(sdkType)
	if canonical == "" {
		return Cell{}, true // can't decide cheaply; let the LLM see it
	}
	for _, a := range spec.AppliesTo {
		if a == canonical {
			return Cell{}, true
		}
	}
	return Cell{
		State:      StateNotApplicable,
		Source:     SourceAppliesTo,
		Confidence: ConfidenceHigh,
		Rationale:  fmt.Sprintf("Spec applies to %s; %s is %s.", strings.Join(spec.AppliesTo, "/"), sdkID, sdkType),
	}, false
}

// canonicalSDKKind maps the sdk-meta types.json values to the spec
// applies-to vocabulary. We treat 'edge' as server-side (matches how the
// specs and existing metadata categorize them) and AI SDKs as server-side
// by default — both are coarse approximations the LLM step can override
// via richer evidence when the cell isn't filtered.
func canonicalSDKKind(t string) string {
	switch t {
	case "client-side":
		return "client-sdk"
	case "server-side":
		return "server-sdk"
	case "edge":
		return "server-sdk"
	case "ai":
		return "server-sdk"
	case "relay":
		return "relay-proxy"
	}
	return ""
}

func normalizeState(s string) string {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "supported":
		return StateSupported
	case "partial":
		return StatePartial
	case "not-supported", "not_supported", "unsupported":
		return StateNotSupported
	case "not-applicable", "not_applicable", "n/a", "na":
		return StateNotApplicable
	case "unknown":
		return StateUnknown
	}
	return ""
}

func normalizeConfidence(s string) string {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "high":
		return ConfidenceHigh
	case "medium", "med", "moderate":
		return ConfidenceMedium
	case "low":
		return ConfidenceLow
	case "n/a", "na", "":
		return ConfidenceNA
	}
	return ConfidenceLow
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-1] + "…"
}

func sortedKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func loadJSON[T any](path string) (T, error) {
	var zero T
	bytes, err := os.ReadFile(path)
	if err != nil {
		return zero, err
	}
	var v T
	if err := json.Unmarshal(bytes, &v); err != nil {
		return zero, err
	}
	return v, nil
}

// FeatureSupportLite is a minimal copy of features.json's per-feature shape,
// used so we don't depend on the genhtml types.
type FeatureSupportLite struct {
	Introduced *string `json:"introduced"`
	Deprecated *string `json:"deprecated"`
	Removed    *string `json:"removed"`
}

type FeatureInfoLite struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ---------- prompt pack ----------

type packInputs struct {
	spec            Spec
	sdkID           string
	sdkName         string
	sdkType         string
	languages       []string
	features        map[string]FeatureSupportLite
	featureInfo     map[string]FeatureInfoLite
	harness         HarnessInfo
	sdkHarness      SDKHarnessSignal
	repoGitHub      string
	specsRepoPath   string
	sdkRepoRoot     string
	specsCommit     string
	harnessCommit   string
	promptVersion   string
	modelIdentifier string
}

// PromptPack is the full set of inputs the LLM sees for a single (sdk, spec)
// cell. It's also what we hash for cache lookups.
type PromptPack struct {
	SpecID            string                        `json:"spec_id"`
	SpecTitle         string                        `json:"spec_title"`
	SpecStatus        string                        `json:"spec_status"`
	SpecAppliesTo     []string                      `json:"spec_applies_to"`
	SpecReadmeText    string                        `json:"spec_readme_text"`
	SpecReadmePath    string                        `json:"spec_readme_path"`
	SpecRequirementCt int                           `json:"spec_requirement_count"`
	SpecSubSpecs      []string                      `json:"spec_sub_specs,omitempty"`

	SDKID             string                        `json:"sdk_id"`
	SDKName           string                        `json:"sdk_name"`
	SDKType           string                        `json:"sdk_type"`
	SDKLanguages      []string                      `json:"sdk_languages,omitempty"`
	SDKRepoGitHub     string                        `json:"sdk_repo_github,omitempty"`

	SDKFeatures       map[string]FeatureSupportLite `json:"sdk_features,omitempty"`
	FeatureInfo       map[string]FeatureInfoLite    `json:"feature_info,omitempty"`

	HarnessCapabilities []HarnessCapability         `json:"harness_capabilities"`
	HarnessTestGroups   []TestGroup                 `json:"harness_test_groups"`
	SDKParticipates     bool                        `json:"sdk_participates"`
	SDKSuppressions     []SuppressionsFile          `json:"sdk_suppressions,omitempty"`

	SDKRepoTree       []string                      `json:"sdk_repo_tree,omitempty"`
	SDKRepoReadme     string                        `json:"sdk_repo_readme,omitempty"`

	SpecsCommit       string                        `json:"specs_commit"`
	HarnessCommit     string                        `json:"harness_commit"`
	SDKCommit         *string                       `json:"sdk_commit"`
	ModelIdentifier   string                        `json:"model_identifier"`
	PromptVersion     string                        `json:"prompt_version"`
}

// Hash returns a SHA-256 hex digest of the canonical JSON for the pack. Used
// for cache lookups so unchanged inputs reuse cached results.
func (p PromptPack) Hash() string {
	bytes, _ := json.Marshal(p)
	sum := sha256.Sum256(bytes)
	return hex.EncodeToString(sum[:])
}

func buildPromptPack(in packInputs) PromptPack {
	pack := PromptPack{
		SpecID:            in.spec.ID,
		SpecTitle:         in.spec.Title,
		SpecStatus:        in.spec.Status,
		SpecAppliesTo:     in.spec.AppliesTo,
		SpecReadmePath:    in.spec.ReadmePath,
		SpecRequirementCt: in.spec.RequirementCount,
		SpecSubSpecs:      in.spec.SubSpecs,
		SDKID:             in.sdkID,
		SDKName:           in.sdkName,
		SDKType:           in.sdkType,
		SDKLanguages:      in.languages,
		SDKRepoGitHub:     in.repoGitHub,
		SDKFeatures:       in.features,
		HarnessCapabilities: in.harness.Capabilities,
		SDKParticipates:   in.sdkHarness.Participates,
		SDKSuppressions:   in.sdkHarness.SuppressionsFiles,
		SpecsCommit:       in.specsCommit,
		HarnessCommit:     in.harnessCommit,
		SDKCommit:         in.sdkHarness.RepoCommit,
		ModelIdentifier:   in.modelIdentifier,
		PromptVersion:     in.promptVersion,
	}

	// Pick the right harness test groups for the SDK kind.
	switch canonicalSDKKind(in.sdkType) {
	case "client-sdk":
		pack.HarnessTestGroups = in.harness.TestGroups["client-side"]
	case "server-sdk":
		pack.HarnessTestGroups = in.harness.TestGroups["server-side"]
	}

	// Trim feature info to just the ones the SDK actually has, so the prompt
	// stays small and on-topic.
	if len(in.features) > 0 {
		pack.FeatureInfo = map[string]FeatureInfoLite{}
		for fid := range in.features {
			if info, ok := in.featureInfo[fid]; ok {
				pack.FeatureInfo[fid] = info
			}
		}
	}

	// Read spec README text. We don't fail the cell if the file went away
	// since the catalog ran; we just send what we have.
	if in.spec.ReadmePath != "" {
		path := filepath.Join(in.specsRepoPath, in.spec.ReadmePath)
		if bytes, err := os.ReadFile(path); err == nil {
			pack.SpecReadmeText = string(bytes)
		}
	}

	// SDK repo navigational signal: a depth-limited tree + README.
	_, name := splitOrgRepo(in.repoGitHub)
	if name != "" {
		repoDir := filepath.Join(in.sdkRepoRoot, name)
		pack.SDKRepoTree = repoTree(repoDir, 3, 200)
		readmePath := filepath.Join(repoDir, "README.md")
		if bytes, err := os.ReadFile(readmePath); err == nil {
			// Keep the first ~6KB to bound prompt size.
			text := string(bytes)
			if len(text) > 6*1024 {
				text = text[:6*1024] + "\n…[truncated]"
			}
			pack.SDKRepoReadme = text
		}
	}

	return pack
}

// repoTree returns up to maxEntries paths under root, depth-first, omitting
// noise directories. It's intentionally small: a navigational hint, not a
// full file listing.
func repoTree(root string, maxDepth, maxEntries int) []string {
	var out []string
	skipDirs := map[string]struct{}{
		".git": {}, "node_modules": {}, "vendor": {}, "build": {}, "target": {},
		".gradle": {}, ".idea": {}, "__pycache__": {}, ".venv": {}, "venv": {},
		"bin": {}, "dist": {},
	}
	rootDepth := strings.Count(filepath.Clean(root), string(os.PathSeparator))
	_ = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		rel, _ := filepath.Rel(root, path)
		if rel == "." {
			return nil
		}
		depth := strings.Count(filepath.Clean(path), string(os.PathSeparator)) - rootDepth
		if d.IsDir() {
			if _, skip := skipDirs[d.Name()]; skip {
				return filepath.SkipDir
			}
			if depth >= maxDepth {
				out = append(out, rel+"/")
				return filepath.SkipDir
			}
			out = append(out, rel+"/")
		} else {
			out = append(out, rel)
		}
		if len(out) >= maxEntries {
			return filepath.SkipAll
		}
		return nil
	})
	sort.Strings(out)
	return out
}

// ---------- cache ----------

type cellCache struct {
	dir string
	mu  sync.Mutex
}

func newCellCache(dir string) *cellCache { return &cellCache{dir: dir} }

func (c *cellCache) Get(hash string) (Cell, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	bytes, err := os.ReadFile(filepath.Join(c.dir, hash+".json"))
	if err != nil {
		return Cell{}, false
	}
	var cell Cell
	if err := json.Unmarshal(bytes, &cell); err != nil {
		return Cell{}, false
	}
	return cell, true
}

func (c *cellCache) Put(hash string, cell Cell) {
	c.mu.Lock()
	defer c.mu.Unlock()
	bytes, err := json.MarshalIndent(cell, "", "  ")
	if err != nil {
		return
	}
	_ = os.WriteFile(filepath.Join(c.dir, hash+".json"), bytes, 0o644)
}

// ---------- judge providers ----------

type Judge interface {
	Judge(ctx context.Context, pack PromptPack) (JudgeResponse, error)
	Model() string
	// Describe returns a human-readable summary of provider+model+region so
	// the run-startup log is self-explanatory when something is misconfigured
	// (e.g. wrong AWS_REGION for the bearer token).
	Describe() string
}

type JudgeResponse struct {
	State         string     `json:"state"`
	Confidence    string     `json:"confidence"`
	Rationale     string     `json:"rationale"`
	Evidence      []Evidence `json:"evidence"`
	NotesForHuman string     `json:"notes_for_human,omitempty"`
}

// defaultProvider picks the judge based on which credentials are present in
// the environment. Bedrock takes precedence: at LaunchDarkly we use Anthropic
// models through AWS Bedrock, so when the bearer token is set that's almost
// certainly what the operator means to use. The Anthropic-direct provider
// stays available for anyone running this tool outside the LD environment.
func defaultProvider() string {
	if os.Getenv("AWS_BEARER_TOKEN_BEDROCK") != "" {
		return "bedrock"
	}
	if os.Getenv("ANTHROPIC_API_KEY") != "" {
		return "anthropic"
	}
	return "noop"
}

// defaultModelFor picks a sensible default model id for the given provider.
// `GENSPECS_MODEL` overrides this if set.
func defaultModelFor(provider string) string {
	if v := os.Getenv("GENSPECS_MODEL"); v != "" {
		return v
	}
	switch provider {
	case "bedrock":
		// us.* is the cross-region inference profile (us-east-1 / us-east-2 /
		// us-west-2). Sonnet 4.5 is the same model the Anthropic-direct
		// branch uses below; keeping the choice aligned makes it easier to
		// compare runs across providers.
		return "us.anthropic.claude-sonnet-4-5-20250929-v1:0"
	case "anthropic":
		return "claude-sonnet-4-5-20250929"
	}
	return ""
}

func buildJudge(provider, model string) (Judge, error) {
	if model == "" {
		model = defaultModelFor(provider)
	}
	var inner Judge
	switch provider {
	case "bedrock":
		token := os.Getenv("AWS_BEARER_TOKEN_BEDROCK")
		if token == "" {
			return nil, fmt.Errorf("AWS_BEARER_TOKEN_BEDROCK not set (generate one at AWS console -> Bedrock -> API Keys -> Generate short-term API keys; tokens are valid 12 hours)")
		}
		region := os.Getenv("AWS_REGION")
		if region == "" {
			region = "us-east-1"
		}
		inner = &bedrockJudge{token: token, region: region, model: model}
	case "anthropic":
		key := os.Getenv("ANTHROPIC_API_KEY")
		if key == "" {
			return nil, fmt.Errorf("ANTHROPIC_API_KEY not set")
		}
		inner = &anthropicJudge{apiKey: key, model: model}
	case "noop":
		// noop never errors — no retry layer needed.
		return &noopJudge{}, nil
	default:
		return nil, fmt.Errorf("unknown provider %q (try 'bedrock', 'anthropic', or 'noop')", provider)
	}
	return &retryingJudge{inner: inner, maxAttempts: 4}, nil
}

// ---------- retry layer ----------

// retryableHTTPError is what the HTTP-backed providers return for non-2xx
// responses. The wrapper inspects it (and `Retry-After` if set) to decide
// whether to back off and retry.
type retryableHTTPError struct {
	StatusCode int
	Status     string
	Body       string
	Provider   string
	RetryAfter time.Duration // 0 if header absent or unparseable
}

func (e *retryableHTTPError) Error() string {
	return fmt.Sprintf("%s %s: %s", e.Provider, e.Status, truncate(e.Body, 240))
}

// retryable reports whether this HTTP status code is worth retrying. 408
// (request timeout), 429 (rate limit), and 5xx (server-side) are. Everything
// else (mainly 4xx) usually means a config error that won't fix itself.
func (e *retryableHTTPError) retryable() bool {
	return e.StatusCode == http.StatusRequestTimeout ||
		e.StatusCode == http.StatusTooManyRequests ||
		(e.StatusCode >= 500 && e.StatusCode < 600)
}

// isRetryable classifies any error returned from a provider's Judge() call
// into "worth a retry" vs "give up". Network-layer errors (timeouts, EOFs,
// connection resets) are always retryable; HTTP errors are retryable per
// retryableHTTPError.retryable().
func isRetryable(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		// Caller cancelled or hit a deadline — respect that, don't retry.
		return false
	}
	var hErr *retryableHTTPError
	if errors.As(err, &hErr) {
		return hErr.retryable()
	}
	// net.Error covers DNS, dial, read timeouts.
	var nErr interface{ Timeout() bool }
	if errors.As(err, &nErr) && nErr.Timeout() {
		return true
	}
	if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
		return true
	}
	// Fallback: catch the common transient error strings the Go runtime
	// produces for TCP-level failures that don't always satisfy the
	// interfaces above (notably "connection reset by peer").
	msg := strings.ToLower(err.Error())
	for _, marker := range []string{
		"connection reset",
		"connection refused",
		"broken pipe",
		"i/o timeout",
		"tls handshake",
		"unexpected eof",
		"no such host", // DNS hiccups during network changes
	} {
		if strings.Contains(msg, marker) {
			return true
		}
	}
	return false
}

// retryingJudge wraps another Judge with bounded exponential-backoff retries
// for transient failures. Backoff is base * 2^(attempt-1) with ±25% jitter,
// capped at 16s. If the underlying error includes a Retry-After hint
// (typical for 429s), that overrides the backoff schedule.
type retryingJudge struct {
	inner       Judge
	maxAttempts int
}

func (r *retryingJudge) Model() string    { return r.inner.Model() }
func (r *retryingJudge) Describe() string { return r.inner.Describe() }

func (r *retryingJudge) Judge(ctx context.Context, pack PromptPack) (JudgeResponse, error) {
	const baseDelay = 750 * time.Millisecond
	const maxDelay = 16 * time.Second
	var lastErr error
	for attempt := 1; attempt <= r.maxAttempts; attempt++ {
		resp, err := r.inner.Judge(ctx, pack)
		if err == nil {
			return resp, nil
		}
		lastErr = err
		if !isRetryable(err) || attempt == r.maxAttempts {
			return JudgeResponse{}, err
		}

		delay := baseDelay << (attempt - 1)
		if delay > maxDelay {
			delay = maxDelay
		}
		// ±25% jitter so concurrent retriers don't all hit the API in lockstep.
		jitter := time.Duration((rand.Float64() - 0.5) * float64(delay) * 0.5)
		delay += jitter

		// Honor server-supplied Retry-After (only on 429s in practice).
		var hErr *retryableHTTPError
		if errors.As(err, &hErr) && hErr.RetryAfter > 0 && hErr.RetryAfter < 60*time.Second {
			delay = hErr.RetryAfter
		}

		fmt.Fprintf(os.Stderr, "  retry %d/%d for %s/%s after %s: %v\n",
			attempt, r.maxAttempts-1, pack.SDKID, pack.SpecID, delay.Round(100*time.Millisecond), err)

		select {
		case <-ctx.Done():
			return JudgeResponse{}, ctx.Err()
		case <-time.After(delay):
		}
	}
	return JudgeResponse{}, lastErr
}

func parseRetryAfter(h string) time.Duration {
	h = strings.TrimSpace(h)
	if h == "" {
		return 0
	}
	if secs, err := strconv.Atoi(h); err == nil && secs >= 0 {
		return time.Duration(secs) * time.Second
	}
	if t, err := http.ParseTime(h); err == nil {
		return time.Until(t)
	}
	return 0
}

// noopJudge returns 'unknown' for every cell. Useful for running the pipeline
// end-to-end without LLM credentials so the schema/storage paths can be
// validated.
type noopJudge struct{}

func (n *noopJudge) Model() string    { return "none" }
func (n *noopJudge) Describe() string { return "provider=noop" }
func (n *noopJudge) Judge(_ context.Context, _ PromptPack) (JudgeResponse, error) {
	return JudgeResponse{
		State:      StateUnknown,
		Confidence: ConfidenceLow,
		Rationale:  "No judge provider configured; cell left unknown.",
	}, nil
}

// bedrockJudge calls Anthropic models via Amazon Bedrock's Runtime
// InvokeModel HTTP endpoint, authenticated with the short-term bearer token
// AWS hands out via "Bedrock -> API Keys -> Generate short-term API keys".
//
// LaunchDarkly setup notes (per #proj-building-with-ai and verified
// 2026-05-12):
//   - Generate the bearer token from the AWS **Development** account, not
//     SDK. The PowerUser role in SDK does not have bedrock:CallWithBearerToken
//     in its identity-based policy.
//   - Bearer tokens are scoped to the account+region they were generated in.
//     us-east-2 is known-good in the Development account; us-east-1 may also
//     work depending on which inference profiles are enabled. Set AWS_REGION
//     to match the region you generated the token in.
//   - Tokens expire after 12 hours.
//
// The wire body is the same Anthropic Messages API shape, with two
// differences:
//   - `anthropic_version` is the Bedrock literal "bedrock-2023-05-31".
//   - The model id is in the URL path, not the body.
//
// See https://docs.anthropic.com/en/api/claude-on-amazon-bedrock for details.
type bedrockJudge struct {
	token  string
	region string
	model  string
}

func (b *bedrockJudge) Model() string { return b.model }
func (b *bedrockJudge) Describe() string {
	return fmt.Sprintf("provider=bedrock region=%s model=%s", b.region, b.model)
}

func (b *bedrockJudge) Judge(ctx context.Context, pack PromptPack) (JudgeResponse, error) {
	systemPrompt := judgeSystemPrompt
	userPrompt, err := renderUserPrompt(pack)
	if err != nil {
		return JudgeResponse{}, err
	}

	body, err := json.Marshal(map[string]any{
		"anthropic_version": "bedrock-2023-05-31",
		"max_tokens":        1024,
		"system":            systemPrompt,
		"messages": []map[string]any{
			{"role": "user", "content": userPrompt},
		},
	})
	if err != nil {
		return JudgeResponse{}, err
	}

	endpoint := fmt.Sprintf(
		"https://bedrock-runtime.%s.amazonaws.com/model/%s/invoke",
		b.region, url.PathEscape(b.model),
	)
	httpResp, err := postJSON(ctx, endpoint, body, map[string]string{
		"authorization": "Bearer " + b.token,
		"content-type":  "application/json",
		"accept":        "application/json",
	})
	if err != nil {
		return JudgeResponse{}, err
	}
	defer httpResp.Body.Close()
	respBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return JudgeResponse{}, err
	}
	if httpResp.StatusCode/100 != 2 {
		return JudgeResponse{}, &retryableHTTPError{
			StatusCode: httpResp.StatusCode,
			Status:     httpResp.Status,
			Body:       string(respBytes),
			Provider:   "bedrock",
			RetryAfter: parseRetryAfter(httpResp.Header.Get("Retry-After")),
		}
	}

	// Bedrock returns the same envelope shape as the Anthropic Messages API.
	var envelope struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}
	if err := json.Unmarshal(respBytes, &envelope); err != nil {
		return JudgeResponse{}, fmt.Errorf("decoding bedrock response: %w", err)
	}
	var text strings.Builder
	for _, c := range envelope.Content {
		if c.Type == "text" {
			text.WriteString(c.Text)
		}
	}
	return parseJudgeJSON(text.String())
}

type anthropicJudge struct {
	apiKey string
	model  string
}

func (a *anthropicJudge) Model() string    { return a.model }
func (a *anthropicJudge) Describe() string { return "provider=anthropic model=" + a.model }

func (a *anthropicJudge) Judge(ctx context.Context, pack PromptPack) (JudgeResponse, error) {
	systemPrompt := judgeSystemPrompt
	userPrompt, err := renderUserPrompt(pack)
	if err != nil {
		return JudgeResponse{}, err
	}

	body, err := json.Marshal(map[string]any{
		"model":      a.model,
		"max_tokens": 1024,
		"system":     systemPrompt,
		"messages": []map[string]any{
			{"role": "user", "content": userPrompt},
		},
	})
	if err != nil {
		return JudgeResponse{}, err
	}

	const url = "https://api.anthropic.com/v1/messages"
	httpResp, err := postJSON(ctx, url, body, map[string]string{
		"x-api-key":         a.apiKey,
		"anthropic-version": "2023-06-01",
		"content-type":      "application/json",
	})
	if err != nil {
		return JudgeResponse{}, err
	}
	defer httpResp.Body.Close()
	respBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return JudgeResponse{}, err
	}
	if httpResp.StatusCode/100 != 2 {
		return JudgeResponse{}, &retryableHTTPError{
			StatusCode: httpResp.StatusCode,
			Status:     httpResp.Status,
			Body:       string(respBytes),
			Provider:   "anthropic",
			RetryAfter: parseRetryAfter(httpResp.Header.Get("Retry-After")),
		}
	}

	var envelope struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}
	if err := json.Unmarshal(respBytes, &envelope); err != nil {
		return JudgeResponse{}, fmt.Errorf("decoding anthropic response: %w", err)
	}
	var text strings.Builder
	for _, c := range envelope.Content {
		if c.Type == "text" {
			text.WriteString(c.Text)
		}
	}

	return parseJudgeJSON(text.String())
}

func parseJudgeJSON(s string) (JudgeResponse, error) {
	// Pull out the first JSON object in the response. The system prompt asks
	// the model to return a single JSON object; tolerate code fences and
	// surrounding prose.
	s = stripCodeFence(s)
	start := strings.Index(s, "{")
	end := strings.LastIndex(s, "}")
	if start < 0 || end < start {
		return JudgeResponse{}, fmt.Errorf("no JSON object in response: %q", truncate(s, 200))
	}
	candidate := s[start : end+1]
	var resp JudgeResponse
	if err := json.Unmarshal([]byte(candidate), &resp); err != nil {
		return JudgeResponse{}, fmt.Errorf("decoding judge JSON: %w (was: %s)", err, truncate(candidate, 200))
	}
	return resp, nil
}

func stripCodeFence(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "```") {
		s = strings.TrimPrefix(s, "```json")
		s = strings.TrimPrefix(s, "```")
	}
	s = strings.TrimSuffix(s, "```")
	return strings.TrimSpace(s)
}

const judgeSystemPrompt = `You are a senior LaunchDarkly SDK engineer evaluating, at a high level, whether a given LaunchDarkly SDK supports a given top-level engineering spec.

You will be given:
- The spec README text and metadata.
- The SDK's metadata (type, languages, GitHub repo).
- The SDK's existing entries in sdk-meta features.json (with feature definitions). Many features map closely to specs; use this as evidence when relevant.
- The sdk-test-harness capability constants (with doc comments) and the harness's top-level test groups for this SDK kind.
- The SDK's testharness-suppressions file(s), if any. Empty/missing files mean different things — see the 'sdk_participates' field.
- A depth-limited directory listing of the SDK repo and its README.

Pick exactly one of: "supported", "partial", "not-supported", "not-applicable".
- "supported" means the SDK clearly implements the bulk of the spec's normative requirements.
- "partial" means it implements some but is missing or differs on others (suppressions in the harness for the relevant area are a strong signal here).
- "not-supported" means there is no meaningful implementation evident.
- "not-applicable" means the spec does not apply to this SDK (only use this if the spec's applies-to or its content makes that clear; a deterministic filter already removed obvious cases).

Also pick a confidence: "high", "medium", or "low".

Return a single JSON object, with no surrounding prose or markdown, matching this shape:

{
  "state": "supported|partial|not-supported|not-applicable",
  "confidence": "high|medium|low",
  "rationale": "one or two sentences (<=240 chars), citing the strongest evidence",
  "evidence": [
    {"kind": "feature", "id": "hooks"},
    {"kind": "harness_capability", "id": "evaluation-hooks"},
    {"kind": "harness_suppression", "path": "evaluation/all flags state/client not ready"},
    {"kind": "repo_path", "path": "src/main/java/com/launchdarkly/sdk/server/Hook.java"}
  ],
  "notes_for_human": "optional; only if you see something the team should know"
}

Rules:
- Be honest about uncertainty. If the only signal is a feature name match but no repo evidence, use "medium" confidence.
- Do not invent file paths or feature ids. Cite only items present in the inputs.
- Keep the rationale short and concrete.
- Output exactly one JSON object. No prose before or after.`

func renderUserPrompt(pack PromptPack) (string, error) {
	bytes, err := json.MarshalIndent(pack, "", "  ")
	if err != nil {
		return "", err
	}
	return "Evaluate this (SDK, spec) cell. Inputs are below as JSON.\n\n" + string(bytes), nil
}

func postJSON(ctx context.Context, url string, body []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return httpClient.Do(req)
}
