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

type Cell struct {
	State         string         `json:"state"`
	Confidence    string         `json:"confidence,omitempty"`
	Source        string         `json:"source"`
	Rationale     string         `json:"rationale,omitempty"`
	Evidence      []Evidence     `json:"evidence,omitempty"`
	JudgedAt      *string        `json:"judged_at,omitempty"`
	JudgedAgainst *JudgedAgainst `json:"judged_against,omitempty"`
	NotesForHuman *string        `json:"notes_for_human,omitempty"`
}

type Evidence struct {
	Kind string  `json:"kind"`
	ID   *string `json:"id,omitempty"`
	Path *string `json:"path,omitempty"`
	Note *string `json:"note,omitempty"`
}

type JudgedAgainst struct {
	SpecsCommit   string  `json:"specs_commit"`
	HarnessCommit string  `json:"harness_commit"`
	SDKCommit     *string `json:"sdk_commit"`
	Model         string  `json:"model"`
	PromptVersion string  `json:"prompt_version"`
}

// Cell states.
const (
	StateSupported     = "supported"
	StatePartial       = "partial"
	StateNotSupported  = "not-supported"
	StateNotApplicable = "not-applicable"
	StateUnknown       = "unknown"
)

// Cell sources.
const (
	SourceAppliesTo   = "applies_to"
	SourceLLMJudge    = "llm_judge"
	SourceJudgeFailed = "judge_failed"
)

// Confidence levels.
const (
	ConfidenceHigh   = "high"
	ConfidenceMedium = "medium"
	ConfidenceLow    = "low"
	ConfidenceNA     = "n/a"
)
