package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// Spec READMEs follow the format described in
// sdk-specs/specs/SPEC-specification-for-specs/README.md:
//
//   | id | status | title | description | applies-to |
//   |----|--------|-------|-------------|------------|
//   | XX | DRAFT  | ...   | ...         | a, b       |
//
// followed by the body containing `### Requirement N.N.N` headings.

var (
	requirementRE = regexp.MustCompile(`(?m)^#{2,4} (?:Conditional )?Requirement [0-9.]+\b`)
	versionDirRE  = regexp.MustCompile(`^v[0-9]+(\.[0-9]+)?$`)
	subSpecDirRE  = regexp.MustCompile(`^([A-Z]+)-`)
	// Recognized status keywords. Anything outside this set is preserved as-is
	// (the spec catalog publishes whatever the README declares — we don't
	// silently drop it).
	statusSet = map[string]struct{}{
		"DRAFT": {}, "ACCEPTED": {}, "DEPRECATED": {}, "SUPERSEDED": {},
		"APPROVED": {}, "CURRENT": {},
	}
	// Some specs use `all-sdks` as a shorthand for both client and server.
	// Other specs use `client-side`/`server-side` (matching the sdk-meta types
	// vocabulary rather than the spec vocabulary). We accept both spellings.
	// Tokens that aren't in this map (e.g. "events", "error" found in some
	// experimental specs) are dropped rather than passed through.
	appliesToExpansions = map[string][]string{
		"client-sdk":  {"client-sdk"},
		"server-sdk":  {"server-sdk"},
		"relay-proxy": {"relay-proxy"},
		"client-side": {"client-sdk"},
		"server-side": {"server-sdk"},
		"all-sdks":    {"client-sdk", "server-sdk"},
	}
	// Strips a leading version prefix like `v1:` or `v2:` from a status entry.
	statusVersionPrefixRE = regexp.MustCompile(`^v[0-9]+(?:\.[0-9]+)?:`)
)

func runCatalog(args []string) error {
	fs := flag.NewFlagSet("catalog", flag.ExitOnError)
	specsRepoPath := fs.String("specs-repo", filepath.Join(defaultReposRoot(), "sdk-specs"), "Path to the local sdk-specs checkout.")
	out := fs.String("out", "products/specs.json", "Output path for the catalog JSON.")
	if err := fs.Parse(args); err != nil {
		return err
	}

	specsDir := filepath.Join(*specsRepoPath, "specs")
	if _, err := os.Stat(specsDir); err != nil {
		return fmt.Errorf("could not find specs dir at %s (did you run `genspecs sync-repos`?): %w", specsDir, err)
	}

	commit, err := headCommit(*specsRepoPath)
	if err != nil {
		return fmt.Errorf("getting sdk-specs HEAD: %w", err)
	}

	entries, err := os.ReadDir(specsDir)
	if err != nil {
		return err
	}

	specs := map[string]Spec{}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		dirName := e.Name()
		m := subSpecDirRE.FindStringSubmatch(dirName)
		if m == nil {
			continue
		}
		specID := m[1]
		spec, err := buildSpec(*specsRepoPath, specsDir, dirName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: could not build catalog entry for %s: %v\n", dirName, err)
			continue
		}
		// The directory prefix is the canonical id used elsewhere (other
		// specs cross-reference each other by it). README-declared ids that
		// drift from the dir name (e.g. legacy `SDKC-LEGACY-ARCO` for
		// directory `ARCO-architecture-overview`) are intentionally ignored.
		spec.ID = specID
		// Title fallback: derive from the directory name if the README didn't
		// supply one. e.g. "RPENDPOINTS-relay-proxy-endpoints" -> "Relay
		// Proxy Endpoints". The schema requires a non-empty title since the
		// HTML viewer uses it as the human label, and we'd rather degrade
		// gracefully than reject the spec.
		if spec.Title == "" {
			spec.Title = titleFromDirName(dirName, specID)
		}
		specs[spec.ID] = spec
	}

	product := SpecsProduct{
		GeneratedAt:  time.Now().UTC().Format(time.RFC3339),
		SourceCommit: commit,
		Specs:        specs,
	}

	return writeJSON(*out, product)
}

func buildSpec(repoRoot, specsDir, dirName string) (Spec, error) {
	dir := filepath.Join(specsDir, dirName)

	// Discover versions and sub-specs by listing the immediate subdirs.
	entries, err := os.ReadDir(dir)
	if err != nil {
		return Spec{}, err
	}
	var versions []string
	var subSpecs []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()
		if versionDirRE.MatchString(name) {
			versions = append(versions, name)
		} else if m := subSpecDirRE.FindStringSubmatch(name); m != nil {
			subSpecs = append(subSpecs, m[1])
		}
	}
	sort.Strings(versions)
	sort.Strings(subSpecs)

	var latestVersion *string
	if len(versions) > 0 {
		v := versions[len(versions)-1]
		latestVersion = &v
	}

	// Pick the canonical README to parse for the header table.
	readmePath, err := pickReadme(dir, latestVersion)
	if err != nil {
		return Spec{}, err
	}

	header, body, err := parseReadme(readmePath)
	if err != nil {
		return Spec{}, fmt.Errorf("parsing %s: %w", readmePath, err)
	}

	// Count requirements across this dir (root README + version dirs + sub-spec dirs, recursively).
	count := countRequirements(body)
	for _, v := range versions {
		count += sumRequirementsInDir(filepath.Join(dir, v))
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if subSpecDirRE.MatchString(e.Name()) {
			count += sumRequirementsInDir(filepath.Join(dir, e.Name()))
		}
	}

	relPath, _ := filepath.Rel(repoRoot, dir)
	relReadme, _ := filepath.Rel(repoRoot, readmePath)

	appliesTo := header.appliesTo
	if appliesTo == nil {
		appliesTo = []string{}
	}

	spec := Spec{
		ID:               header.id,
		Title:            header.title,
		Status:           header.status,
		Description:      header.description,
		AppliesTo:        appliesTo,
		Path:             relPath,
		ReadmePath:       relReadme,
		LatestVersion:    latestVersion,
		Versions:         versions,
		SubSpecs:         subSpecs,
		RequirementCount: count,
	}
	return spec, nil
}

// pickReadme returns the README that best describes the spec. We prefer the
// top-level README.md when it has a parsable header table; otherwise we fall
// through to the latest versioned README (specs like DATASYSTEM and TDS keep
// only a folder-level explanation at the top and put the real spec in v1/v2).
func pickReadme(dir string, latestVersion *string) (string, error) {
	candidates := []string{filepath.Join(dir, "README.md")}
	if latestVersion != nil {
		candidates = append(candidates, filepath.Join(dir, *latestVersion, "README.md"))
	}

	// Walk for any README.md as a final fallback.
	var walkFound string
	_ = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Base(path) == "README.md" {
			walkFound = path
			return filepath.SkipAll
		}
		return nil
	})
	if walkFound != "" {
		candidates = append(candidates, walkFound)
	}

	var firstExisting string
	for _, c := range candidates {
		if _, err := os.Stat(c); err != nil {
			continue
		}
		if firstExisting == "" {
			firstExisting = c
		}
		bytes, err := os.ReadFile(c)
		if err != nil {
			continue
		}
		if hasHeaderTable(string(bytes)) {
			return c, nil
		}
	}
	if firstExisting != "" {
		return firstExisting, nil
	}
	return "", fmt.Errorf("no README.md found under %s", dir)
}

func hasHeaderTable(content string) bool {
	header := parseHeaderTable(content)
	return header.id != "" || header.title != ""
}

type readmeHeader struct {
	id          string
	status      string
	title       string
	description string
	appliesTo   []string
}

func parseReadme(path string) (readmeHeader, string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return readmeHeader{}, "", err
	}
	content := string(bytes)
	header := parseHeaderTable(content)
	return header, content, nil
}

// parseHeaderTable scans for the first markdown table whose header row contains
// id|status|title|description|applies-to (in any order) and returns the first
// data row's contents.
func parseHeaderTable(content string) readmeHeader {
	lines := strings.Split(content, "\n")
	for i := 0; i < len(lines); i++ {
		row := strings.TrimSpace(lines[i])
		if !strings.HasPrefix(row, "|") {
			continue
		}
		// Header row?
		cols := splitMarkdownRow(row)
		hasID, hasTitle := false, false
		colIdx := map[string]int{}
		for idx, c := range cols {
			lc := strings.ToLower(strings.TrimSpace(c))
			colIdx[lc] = idx
			switch lc {
			case "id":
				hasID = true
			case "title":
				hasTitle = true
			}
		}
		// `id` and `title` are the minimum we need to recognize this as a
		// spec header. Some specs (RPENDPOINTS) omit status/applies-to.
		if !(hasID && hasTitle) {
			continue
		}
		// Next line should be the separator (|---|---|...). Then the data row.
		if i+2 >= len(lines) {
			break
		}
		dataLine := strings.TrimSpace(lines[i+2])
		if !strings.HasPrefix(dataLine, "|") {
			break
		}
		dataCols := splitMarkdownRow(dataLine)
		get := func(name string) string {
			if idx, ok := colIdx[name]; ok && idx < len(dataCols) {
				return strings.TrimSpace(dataCols[idx])
			}
			return ""
		}
		h := readmeHeader{
			id:          get("id"),
			status:      normalizeStatus(get("status")),
			title:       get("title"),
			description: get("description"),
		}
		// applies-to may be multiple comma-delimited tags. Some specs use
		// `all-sdks` as shorthand; expand it to its constituents. Unknown
		// tags are dropped so the LLM judge step doesn't have to defend
		// against arbitrary vocabulary.
		raw := get("applies-to")
		seen := map[string]struct{}{}
		if raw != "" {
			for _, t := range strings.Split(raw, ",") {
				tag := strings.ToLower(strings.TrimSpace(t))
				if tag == "" {
					continue
				}
				expansions, ok := appliesToExpansions[tag]
				if !ok {
					continue
				}
				for _, e := range expansions {
					if _, dup := seen[e]; dup {
						continue
					}
					seen[e] = struct{}{}
					h.appliesTo = append(h.appliesTo, e)
				}
			}
			sort.Strings(h.appliesTo)
		}
		return h
	}
	return readmeHeader{}
}

// normalizeStatus picks one canonical status from a (possibly multi-version)
// raw status field. Examples seen in the wild:
//
//	"DRAFT"                       -> "DRAFT"
//	"v1:DRAFT"                    -> "DRAFT"
//	"v1:ACCEPTED, v2:DRAFT"       -> "ACCEPTED" (highest version wins)
//	"APPROVED"                    -> "APPROVED"
//
// Unknown tokens fall through unchanged so we don't lose information.
func normalizeStatus(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}
	parts := strings.Split(raw, ",")
	type entry struct {
		version int
		status  string
	}
	var entries []entry
	for _, p := range parts {
		token := strings.TrimSpace(p)
		ver := 0
		if m := statusVersionPrefixRE.FindString(token); m != "" {
			fmt.Sscanf(m, "v%d:", &ver)
			token = strings.TrimPrefix(token, m)
		}
		token = strings.ToUpper(strings.TrimSpace(token))
		entries = append(entries, entry{version: ver, status: token})
	}
	if len(entries) == 0 {
		return raw
	}
	// Pick the highest-versioned entry; ties prefer the one declared first.
	best := entries[0]
	for _, e := range entries[1:] {
		if e.version > best.version {
			best = e
		}
	}
	if _, ok := statusSet[best.status]; ok {
		return best.status
	}
	return best.status
}

func titleFromDirName(dir, id string) string {
	rest := strings.TrimPrefix(dir, id)
	rest = strings.TrimPrefix(rest, "-")
	if rest == "" {
		return id
	}
	parts := strings.Split(rest, "-")
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, " ")
}

func splitMarkdownRow(row string) []string {
	row = strings.TrimSpace(row)
	row = strings.TrimPrefix(row, "|")
	row = strings.TrimSuffix(row, "|")
	cols := strings.Split(row, "|")
	for i, c := range cols {
		cols[i] = strings.TrimSpace(c)
	}
	return cols
}

func countRequirements(body string) int {
	return len(requirementRE.FindAllString(body, -1))
}

func sumRequirementsInDir(dir string) int {
	total := 0
	_ = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Base(path) != "README.md" {
			return nil
		}
		bytes, rdErr := os.ReadFile(path)
		if rdErr != nil {
			return nil
		}
		total += countRequirements(string(bytes))
		return nil
	})
	return total
}

func writeJSON(path string, v any) error {
	if dir := filepath.Dir(path); dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	bytes = append(bytes, '\n')
	return os.WriteFile(path, bytes, 0o644)
}
