package main

// Types in this file mirror the JSON schemas under sdk-meta/schemas/. Keep them
// in sync. Where the schema specifies a fixed enum, callers are responsible for
// passing the correct string; we don't enforce it at the type level so the JSON
// surface stays straightforward.

// ---------- specs.json ----------

type SpecsProduct struct {
	GeneratedAt  string          `json:"generated_at"`
	SourceCommit string          `json:"source_commit"`
	Specs        map[string]Spec `json:"specs"`
}

type Spec struct {
	ID               string   `json:"id"`
	Title            string   `json:"title"`
	Status           string   `json:"status"`
	Description      string   `json:"description"`
	AppliesTo        []string `json:"applies_to"`
	Path             string   `json:"path"`
	ReadmePath       string   `json:"readme_path,omitempty"`
	LatestVersion    *string  `json:"latest_version"`
	Versions         []string `json:"versions,omitempty"`
	SubSpecs         []string `json:"sub_specs,omitempty"`
	RequirementCount int      `json:"requirement_count"`
}

// ---------- harness_signals.json ----------

type HarnessProduct struct {
	GeneratedAt string                       `json:"generated_at"`
	Harness     HarnessInfo                  `json:"harness"`
	SDKs        map[string]SDKHarnessSignal  `json:"sdks"`
}

type HarnessInfo struct {
	Commit       string                  `json:"commit"`
	Capabilities []HarnessCapability     `json:"capabilities"`
	TestGroups   map[string][]TestGroup  `json:"test_groups"`
}

type HarnessCapability struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Doc   string `json:"doc,omitempty"`
}

type TestGroup struct {
	Name string `json:"name"`
	Func string `json:"func"`
}

type SDKHarnessSignal struct {
	Participates           bool                  `json:"participates"`
	Repo                   string                `json:"repo,omitempty"`
	RepoCommit             *string               `json:"repo_commit"`
	CurrentRef             *string               `json:"current_ref,omitempty"`
	SuppressionsFiles      []SuppressionsFile    `json:"suppressions_files,omitempty"`
	SuppressionsByTopGroup map[string]int        `json:"suppressions_by_top_group,omitempty"`
}

type SuppressionsFile struct {
	Path     string             `json:"path"`
	Variant  *string            `json:"variant"`
	Lines    []string           `json:"lines"`
	Comments []SuppressionGroup `json:"comments"`
}

type SuppressionGroup struct {
	Comment    string   `json:"comment"`
	AppliesTo  []string `json:"applies_to,omitempty"`
}

// ---------- spec_support.json ----------

type SpecSupportProduct struct {
	GeneratedAt   string                          `json:"generated_at"`
	SpecsCommit   string                          `json:"specs_commit"`
	HarnessCommit string                          `json:"harness_commit"`
	Model         string                          `json:"model"`
	PromptVersion string                          `json:"prompt_version"`
	SDKs          map[string]map[string]Cell      `json:"sdks"`
}

// Cell is intentionally slim. Audit metadata (source, evidence list,
// judged-at, judged-against commits) is captured in the on-disk judge cache
// (tool/specs/.judge-cache/), not in the public product. The cache lets us
// trace any decision back to its inputs without bloating spec_support.json.
type Cell struct {
	State         string  `json:"state"`
	Confidence    string  `json:"confidence,omitempty"`
	Rationale     string  `json:"rationale,omitempty"`
	NotesForHuman *string `json:"notes_for_human,omitempty"`
}

// Evidence is what the LLM cites as it reasons about a cell. We still ask
// for it in the prompt (chain-of-thought tends to improve answer quality),
// but we don't persist it in the public product. Kept as a type because the
// judge response decoder uses it.
type Evidence struct {
	Kind string  `json:"kind"`
	ID   *string `json:"id,omitempty"`
	Path *string `json:"path,omitempty"`
	Note *string `json:"note,omitempty"`
}

// Cell states.
const (
	StateSupported     = "supported"
	StatePartial       = "partial"
	StateNotSupported  = "not-supported"
	StateNotApplicable = "not-applicable"
	StateUnknown       = "unknown"
)

// Confidence levels.
const (
	ConfidenceHigh   = "high"
	ConfidenceMedium = "medium"
	ConfidenceLow    = "low"
	ConfidenceNA     = "n/a"
)
