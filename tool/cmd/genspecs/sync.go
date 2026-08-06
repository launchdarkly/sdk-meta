package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

const (
	specsRepo   = "launchdarkly/sdk-specs"
	harnessRepo = "launchdarkly/sdk-test-harness"
)

type syncResult struct {
	Repo       string
	LocalPath  string
	Status     string // cloned | updated | unchanged | dirty-skipped | branch-skipped | clone-failed | fetch-failed
	HeadCommit string
	CurrentRef string // empty when on default branch
	Err        error
}

func runSync(args []string) error {
	fs := flag.NewFlagSet("sync-repos", flag.ExitOnError)
	reposJSON := fs.String("repos-json", "products/repos.json", "Path to products/repos.json.")
	root := fs.String("sdk-repos-root", defaultReposRoot(), "Directory under which SDK repos live.")
	concurrency := fs.Int("concurrency", 8, "Maximum concurrent git operations.")
	fullClone := fs.Bool("full-clone", false, "If set, do a full clone instead of --depth=1 for missing repos.")
	jsonOut := fs.Bool("json", false, "Print a machine-readable summary to stdout.")
	if err := fs.Parse(args); err != nil {
		return err
	}

	repos, err := readReposJSON(*reposJSON)
	if err != nil {
		return fmt.Errorf("reading repos.json: %w", err)
	}

	// Deduplicate by GitHub org/repo since several SDKs share monorepos.
	repoSet := map[string]struct{}{
		specsRepo:   {},
		harnessRepo: {},
	}
	for _, r := range repos {
		repoSet[r] = struct{}{}
	}
	allRepos := make([]string, 0, len(repoSet))
	for r := range repoSet {
		allRepos = append(allRepos, r)
	}
	sort.Strings(allRepos)

	if err := os.MkdirAll(*root, 0o755); err != nil {
		return fmt.Errorf("creating repos root %s: %w", *root, err)
	}

	results := make([]syncResult, len(allRepos))
	sem := make(chan struct{}, *concurrency)
	var wg sync.WaitGroup
	for i, repo := range allRepos {
		wg.Add(1)
		i, repo := i, repo
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			results[i] = syncOne(repo, *root, *fullClone)
		}()
	}
	wg.Wait()

	// Print per-repo line and tally.
	var tally = map[string]int{}
	for _, r := range results {
		tally[r.Status]++
		short := r.HeadCommit
		if len(short) > 8 {
			short = short[:8]
		}
		extra := ""
		if r.CurrentRef != "" {
			extra = " (on " + r.CurrentRef + ")"
		}
		if r.Err != nil {
			extra += " — " + r.Err.Error()
		}
		fmt.Fprintf(os.Stderr, "  %-12s %-40s %s%s\n", r.Status, r.Repo, short, extra)
	}

	fmt.Fprintf(os.Stderr, "\nsync-repos complete. ")
	fmt.Fprintln(os.Stderr, summarizeTally(tally))

	if *jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(results); err != nil {
			return err
		}
	}

	for _, r := range results {
		if r.Status == "clone-failed" {
			return fmt.Errorf("at least one repo failed to clone (%s)", r.Repo)
		}
	}
	return nil
}

func syncOne(repo, root string, fullClone bool) syncResult {
	res := syncResult{Repo: repo}
	_, name := splitOrgRepo(repo)
	res.LocalPath = filepath.Join(root, name)

	if _, err := os.Stat(filepath.Join(res.LocalPath, ".git")); errors.Is(err, os.ErrNotExist) {
		// Missing — clone fresh.
		args := []string{"clone"}
		if !fullClone {
			args = append(args, "--depth=1")
		}
		args = append(args, fmt.Sprintf("https://github.com/%s.git", repo), res.LocalPath)
		if err := runGit(context.Background(), root, args...); err != nil {
			res.Status = "clone-failed"
			res.Err = err
			return res
		}
		res.Status = "cloned"
	} else if err != nil {
		res.Status = "clone-failed"
		res.Err = fmt.Errorf("statting repo: %w", err)
		return res
	} else {
		// Present — fetch and try to ff-pull, but never blow away WIP.
		if err := runGit(context.Background(), res.LocalPath, "fetch", "--tags", "--quiet", "origin"); err != nil {
			res.Status = "fetch-failed"
			res.Err = err
			// Don't bail; we can still record HEAD.
		} else {
			defaultBranch, dbErr := defaultBranchName(res.LocalPath)
			currentBranch, _ := currentBranchName(res.LocalPath)
			dirty := workingTreeDirty(res.LocalPath)

			switch {
			case dirty:
				res.Status = "dirty-skipped"
			case dbErr != nil || currentBranch == "" || currentBranch != defaultBranch:
				res.Status = "branch-skipped"
				if currentBranch != "" && currentBranch != defaultBranch {
					res.CurrentRef = currentBranch
				}
			default:
				before, _ := headCommit(res.LocalPath)
				if err := runGit(context.Background(), res.LocalPath, "pull", "--ff-only", "--quiet"); err != nil {
					res.Status = "fetch-failed"
					res.Err = err
				} else {
					after, _ := headCommit(res.LocalPath)
					if after != before {
						res.Status = "updated"
					} else {
						res.Status = "unchanged"
					}
				}
			}
		}
	}

	res.HeadCommit, _ = headCommit(res.LocalPath)
	if res.CurrentRef == "" {
		current, _ := currentBranchName(res.LocalPath)
		def, _ := defaultBranchName(res.LocalPath)
		if current != "" && current != def {
			res.CurrentRef = current
		}
	}
	return res
}

func defaultReposRoot() string {
	if v := os.Getenv("SDK_REPOS_ROOT"); v != "" {
		return v
	}
	if home, err := os.UserHomeDir(); err == nil {
		// Mirror the convention used by the rest of the LD tooling.
		candidate := filepath.Join(home, "code", "launchdarkly")
		return candidate
	}
	return "."
}

func readReposJSON(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var raw map[string]struct {
		GitHub string `json:"github"`
	}
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return nil, err
	}
	seen := map[string]struct{}{}
	var out []string
	for _, v := range raw {
		if v.GitHub == "" {
			continue
		}
		if _, ok := seen[v.GitHub]; ok {
			continue
		}
		seen[v.GitHub] = struct{}{}
		out = append(out, v.GitHub)
	}
	sort.Strings(out)
	return out, nil
}

// readReposJSONFull returns the repos.json contents as sdk_id -> github
// org/repo string. The harness and judge steps both need to know which sdk a
// given local checkout belongs to.
func readReposJSONFull(path string) (map[string]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var raw map[string]struct {
		GitHub string `json:"github"`
	}
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return nil, err
	}
	out := make(map[string]string, len(raw))
	for sdkID, v := range raw {
		out[sdkID] = v.GitHub
	}
	return out, nil
}

func splitOrgRepo(s string) (string, string) {
	parts := strings.SplitN(s, "/", 2)
	if len(parts) != 2 {
		return "", s
	}
	return parts[0], parts[1]
}

func runGit(ctx context.Context, dir string, args ...string) error {
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = dir
	cmd.Stdin = nil
	cmd.Stdout = io.Discard
	var stderr strings.Builder
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			msg = err.Error()
		}
		return fmt.Errorf("git %s: %s", strings.Join(args, " "), msg)
	}
	return nil
}

func gitOutput(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func headCommit(dir string) (string, error) {
	return gitOutput(dir, "rev-parse", "HEAD")
}

func currentBranchName(dir string) (string, error) {
	out, err := gitOutput(dir, "symbolic-ref", "--short", "HEAD")
	if err != nil {
		return "", err
	}
	return out, nil
}

func defaultBranchName(dir string) (string, error) {
	if out, err := gitOutput(dir, "symbolic-ref", "--short", "refs/remotes/origin/HEAD"); err == nil {
		// origin/main → main
		_, branch := splitOrgRepo(out)
		return branch, nil
	}
	// Fallback to common defaults.
	for _, candidate := range []string{"main", "master", "v2"} {
		if _, err := gitOutput(dir, "rev-parse", "--verify", "refs/remotes/origin/"+candidate); err == nil {
			return candidate, nil
		}
	}
	return "", fmt.Errorf("could not determine default branch for %s", dir)
}

func workingTreeDirty(dir string) bool {
	out, err := gitOutput(dir, "status", "--porcelain")
	if err != nil {
		return false
	}
	return out != ""
}

func summarizeTally(tally map[string]int) string {
	keys := make([]string, 0, len(tally))
	for k := range tally {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%d", k, tally[k]))
	}
	return strings.Join(parts, " ")
}
