package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

var suppressionsRE = regexp.MustCompile(`^testharness-suppressions(?:-([a-z0-9]+))?\.txt$`)

func runHarness(args []string) error {
	fs := flag.NewFlagSet("harness", flag.ExitOnError)
	harnessRepoPath := fs.String("harness-repo", filepath.Join(defaultReposRoot(), "sdk-test-harness"), "Path to the local sdk-test-harness checkout.")
	root := fs.String("sdk-repos-root", defaultReposRoot(), "Directory under which SDK repos live.")
	reposJSON := fs.String("repos-json", "products/repos.json", "Path to products/repos.json.")
	out := fs.String("out", "products/harness_signals.json", "Output path for the harness_signals JSON.")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if _, err := os.Stat(*harnessRepoPath); err != nil {
		return fmt.Errorf("could not find sdk-test-harness at %s (did you run `genspecs sync-repos`?): %w", *harnessRepoPath, err)
	}

	commit, err := headCommit(*harnessRepoPath)
	if err != nil {
		return fmt.Errorf("getting sdk-test-harness HEAD: %w", err)
	}

	caps, err := parseCapabilities(*harnessRepoPath)
	if err != nil {
		return fmt.Errorf("parsing capabilities: %w", err)
	}

	groups, err := parseTestGroups(*harnessRepoPath)
	if err != nil {
		return fmt.Errorf("parsing test groups: %w", err)
	}

	sdkRepos, err := readReposJSONFull(*reposJSON)
	if err != nil {
		return fmt.Errorf("reading repos.json: %w", err)
	}

	sdkSignals := make(map[string]SDKHarnessSignal, len(sdkRepos))
	for sdkID, repo := range sdkRepos {
		signal := buildSDKSignal(*root, repo)
		sdkSignals[sdkID] = signal
	}

	product := HarnessProduct{
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Harness: HarnessInfo{
			Commit:       commit,
			Capabilities: caps,
			TestGroups:   groups,
		},
		SDKs: sdkSignals,
	}

	return writeJSON(*out, product)
}

// parseCapabilities walks sdk-test-harness/servicedef/ for *.go files and
// extracts every constant named Capability* — value, name, and the doc comment
// (if any).
func parseCapabilities(harnessRepoPath string) ([]HarnessCapability, error) {
	dir := filepath.Join(harnessRepoPath, "servicedef")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var caps []HarnessCapability
	fset := token.NewFileSet()
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		path := filepath.Join(dir, e.Name())
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("parsing %s: %w", path, err)
		}
		for _, decl := range file.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.CONST {
				continue
			}
			for _, spec := range gen.Specs {
				vs, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				for i, name := range vs.Names {
					if !strings.HasPrefix(name.Name, "Capability") {
						continue
					}
					if i >= len(vs.Values) {
						continue
					}
					value, ok := stringLit(vs.Values[i])
					if !ok {
						continue
					}
					doc := commentText(vs.Doc)
					if doc == "" {
						doc = commentText(gen.Doc)
					}
					caps = append(caps, HarnessCapability{
						Name:  name.Name,
						Value: value,
						Doc:   doc,
					})
				}
			}
		}
	}
	sort.Slice(caps, func(i, j int) bool { return caps[i].Name < caps[j].Name })
	return caps, nil
}

func stringLit(e ast.Expr) (string, bool) {
	bl, ok := e.(*ast.BasicLit)
	if !ok || bl.Kind != token.STRING {
		return "", false
	}
	s, err := strconvUnquote(bl.Value)
	if err != nil {
		return "", false
	}
	return s, true
}

// strconvUnquote is a tiny inline alternative to strconv.Unquote so we don't
// have to widen the imports footprint.
func strconvUnquote(s string) (string, error) {
	if len(s) < 2 {
		return "", fmt.Errorf("not a string literal: %q", s)
	}
	if s[0] != '"' || s[len(s)-1] != '"' {
		return "", fmt.Errorf("not a double-quoted literal: %q", s)
	}
	// We assume no escape sequences in capability/test-group names.
	return s[1 : len(s)-1], nil
}

func commentText(group *ast.CommentGroup) string {
	if group == nil {
		return ""
	}
	var lines []string
	for _, c := range group.List {
		text := strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(c.Text, "//"), "/*"))
		text = strings.TrimSuffix(text, "*/")
		text = strings.TrimSpace(text)
		if text != "" {
			lines = append(lines, text)
		}
	}
	return strings.Join(lines, " ")
}

// parseTestGroups looks at sdk-test-harness/sdktests/testsuite_entry_point.go
// for the doAll{ServerSide,ClientSide,PHP}Tests funcs and lists the t.Run()
// invocations within them. The group "name" is the string literal arg, and
// "func" is the source text of the second arg (function reference or
// closure expression).
func parseTestGroups(harnessRepoPath string) (map[string][]TestGroup, error) {
	path := filepath.Join(harnessRepoPath, "sdktests", "testsuite_entry_point.go")
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}

	wanted := map[string]string{
		"doAllServerSideTests": "server-side",
		"doAllClientSideTests": "client-side",
		"doAllPHPTests":        "php",
	}

	groups := make(map[string][]TestGroup, len(wanted))
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		key, ok := wanted[fn.Name.Name]
		if !ok {
			continue
		}
		var found []TestGroup
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			if sel.Sel.Name != "Run" {
				return true
			}
			if len(call.Args) < 2 {
				return true
			}
			name, ok := stringLit(call.Args[0])
			if !ok {
				return true
			}
			fnText := exprToString(fset, call.Args[1])
			found = append(found, TestGroup{Name: name, Func: fnText})
			return true
		})
		groups[key] = found
	}
	return groups, nil
}

func exprToString(fset *token.FileSet, e ast.Expr) string {
	if e == nil {
		return ""
	}
	pos := fset.Position(e.Pos())
	end := fset.Position(e.End())
	if pos.Filename == "" || pos.Filename != end.Filename {
		return ""
	}
	bytes, err := os.ReadFile(pos.Filename)
	if err != nil {
		return ""
	}
	if pos.Offset < 0 || end.Offset > len(bytes) {
		return ""
	}
	return strings.TrimSpace(string(bytes[pos.Offset:end.Offset]))
}

func buildSDKSignal(root, repo string) SDKHarnessSignal {
	signal := SDKHarnessSignal{Repo: repo}
	if repo == "" {
		return signal
	}
	_, name := splitOrgRepo(repo)
	repoDir := filepath.Join(root, name)
	if _, err := os.Stat(repoDir); err != nil {
		return signal // missing locally — participates=false, repo_commit=null
	}
	if commit, err := headCommit(repoDir); err == nil {
		signal.RepoCommit = &commit
	}
	if cur, err := currentBranchName(repoDir); err == nil {
		def, _ := defaultBranchName(repoDir)
		if cur != "" && def != "" && cur != def {
			signal.CurrentRef = &cur
		}
	}

	files := findSuppressionsFiles(repoDir)
	if len(files) == 0 && !hasContractTestsDir(repoDir) {
		return signal
	}
	signal.Participates = true

	bucket := map[string]int{}
	for _, p := range files {
		rel, _ := filepath.Rel(repoDir, p)
		variant := classifyVariant(filepath.Base(p))
		lines, comments := parseSuppressionsFile(p)
		if lines == nil {
			lines = []string{}
		}
		if comments == nil {
			comments = []SuppressionGroup{}
		}
		signal.SuppressionsFiles = append(signal.SuppressionsFiles, SuppressionsFile{
			Path:     rel,
			Variant:  variant,
			Lines:    lines,
			Comments: comments,
		})
		for _, line := range lines {
			top := topGroup(line)
			if top != "" {
				bucket[top]++
			}
		}
	}
	if len(bucket) > 0 {
		signal.SuppressionsByTopGroup = bucket
	}
	// Sort files for deterministic output.
	sort.Slice(signal.SuppressionsFiles, func(i, j int) bool {
		return signal.SuppressionsFiles[i].Path < signal.SuppressionsFiles[j].Path
	})
	return signal
}

func findSuppressionsFiles(repoDir string) []string {
	var out []string
	_ = filepath.WalkDir(repoDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			// Skip well-known noise dirs.
			base := filepath.Base(path)
			if base == ".git" || base == "node_modules" || base == "vendor" || base == "build" || base == "target" || base == ".gradle" {
				return filepath.SkipDir
			}
			return nil
		}
		if suppressionsRE.MatchString(d.Name()) {
			out = append(out, path)
		}
		return nil
	})
	return out
}

func hasContractTestsDir(repoDir string) bool {
	candidates := []string{"contract-tests", "contract_tests", "contracttests", "ContractTests", "test-service"}
	for _, c := range candidates {
		if _, err := os.Stat(filepath.Join(repoDir, c)); err == nil {
			return true
		}
		// One level deeper for monorepos.
		matches, _ := filepath.Glob(filepath.Join(repoDir, "*", c))
		if len(matches) > 0 {
			return true
		}
	}
	return false
}

func classifyVariant(filename string) *string {
	m := suppressionsRE.FindStringSubmatch(filename)
	if m == nil || m[1] == "" {
		return nil
	}
	v := m[1]
	return &v
}

// parseSuppressionsFile preserves comment groups by associating each `#`
// comment with the suppression entries that immediately follow it (until the
// next blank line or comment). This makes the LLM judge step richer: many
// suppressions files have human-written explanations like "# Roku does not
// yet support REPORT verb" that we want to keep.
func parseSuppressionsFile(path string) ([]string, []SuppressionGroup) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, nil
	}
	var lines []string
	var groups []SuppressionGroup
	var pending *SuppressionGroup
	for _, raw := range strings.Split(string(bytes), "\n") {
		trimmed := strings.TrimRight(raw, "\r\n\t ")
		if trimmed == "" {
			pending = nil
			continue
		}
		stripped := strings.TrimLeft(trimmed, " \t")
		if strings.HasPrefix(stripped, "#") {
			text := strings.TrimSpace(strings.TrimPrefix(stripped, "#"))
			pending = &SuppressionGroup{Comment: text}
			groups = append(groups, *pending)
			// keep pointer into slice, see below
			pending = &groups[len(groups)-1]
			continue
		}
		lines = append(lines, trimmed)
		if pending != nil {
			pending.AppliesTo = append(pending.AppliesTo, trimmed)
		}
	}
	return lines, groups
}

func topGroup(line string) string {
	if line == "" {
		return ""
	}
	if i := strings.Index(line, "/"); i >= 0 {
		return line[:i]
	}
	return line
}
