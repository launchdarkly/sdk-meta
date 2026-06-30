package validate

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// runBatches validates batch-mode units a handful of times in warm
// workspaces instead of once per snippet. Units are grouped by runtime and
// build-affecting env (so snippets that need a different SDK version or
// build flag don't share a workspace), each group's image is built once,
// then the group is partitioned across up to `jobs` concurrent harness
// invocations. Each invocation receives a manifest of staged snippets and
// loops over them inside a single warm project (see the batch harness
// contract in CLAUDE.md / each run.sh).
func runBatches(cfg Config, units []*resolvedUnit) error {
	groups, order := groupUnits(units)
	jobs := cfg.Jobs
	if jobs <= 0 {
		jobs = runtime.NumCPU()
	}
	if jobs < 1 {
		jobs = 1
	}
	for _, key := range order {
		g := groups[key]
		// Native validators run directly on the host (no container
		// isolation), so concurrent shards would contend on shared state —
		// the Homebrew download lock, a single iOS Simulator runtime, the
		// shared SwiftPM/DerivedData caches. Run native groups single-shard;
		// one warm workspace resolving dependencies once is also the optimal
		// shape there. Docker groups keep the worker pool (each shard is an
		// isolated container).
		effJobs := jobs
		if g[0].runner.Mode == "native" {
			effJobs = 1
		} else {
			fmt.Printf("--- building %s validator image (%d snippets, %d-way) ---\n",
				g[0].runtime, len(g), min(jobs, len(g)))
			if err := buildImage(cfg, g[0].runner, g[0].runnerDir, os.Stdout); err != nil {
				return err
			}
		}
		if err := runGroup(cfg, g, effJobs); err != nil {
			return err
		}
	}
	return nil
}

// groupUnits buckets units by (runtime, build-affecting env). Units in one
// bucket share a warm workspace, so their env must be identical — the
// bucket key is the canonical encoding of extraEnv. SNIPPET_CHECK is part
// of extraEnv, so parse-kind and runtime-kind units never share a bucket.
// The returned order is deterministic for stable logs.
func groupUnits(units []*resolvedUnit) (map[string][]*resolvedUnit, []string) {
	groups := map[string][]*resolvedUnit{}
	var order []string
	for _, u := range units {
		key := u.runtime + "\x00" + canonicalEnv(u.extraEnv)
		if _, ok := groups[key]; !ok {
			order = append(order, key)
		}
		groups[key] = append(groups[key], u)
	}
	sort.Strings(order)
	return groups, order
}

// canonicalEnv encodes a map as a sorted, NUL-joined k=v string so two
// equal maps always produce the same bucket key.
func canonicalEnv(m map[string]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for _, k := range keys {
		b.WriteString(k)
		b.WriteByte('=')
		b.WriteString(m[k])
		b.WriteByte(0)
	}
	return b.String()
}

// runGroup partitions one same-config group across up to `jobs` shards and
// runs them concurrently, each as a single warm-workspace harness
// invocation. Shard output is captured and flushed in order so concurrent
// logs stay readable. The first non-nil shard error is returned (after all
// shards finish, since fail-fast across an already-running batch buys
// nothing).
func runGroup(cfg Config, g []*resolvedUnit, jobs int) error {
	nShards := min(jobs, len(g))
	shards := make([][]*resolvedUnit, nShards)
	for i, u := range g {
		shards[i%nShards] = append(shards[i%nShards], u)
	}

	var wg sync.WaitGroup
	errs := make([]error, nShards)
	outs := make([]bytes.Buffer, nShards)
	for i := range shards {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			errs[i] = runShard(cfg, shards[i], &outs[i])
		}(i)
	}
	wg.Wait()

	for i := range outs {
		_, _ = os.Stdout.Write(outs[i].Bytes())
	}
	for _, e := range errs {
		if e != nil {
			return e
		}
	}
	return nil
}

// runShard stages every unit in the shard under one batch directory, writes
// a manifest, and runs the harness once over the whole set.
//
// Manifest format: one TSV line per snippet, `<relpath>\t<label>`, where
// relpath is the path under /snippet (or $SNIPPET_DIR) to the snippet's
// entry file (`<index>/<entrypoint>`). The harness loops over these in a
// single warm workspace and exits non-zero if any snippet fails.
func runShard(cfg Config, units []*resolvedUnit, out io.Writer) error {
	if len(units) == 0 {
		return nil
	}
	batchDir, err := os.MkdirTemp("", "snippets-batch-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(batchDir)

	var manifest strings.Builder
	for i, u := range units {
		staged, err := u.stage()
		if err != nil {
			return err
		}
		sub := filepath.Join(batchDir, strconv.Itoa(i))
		if err := os.Rename(staged, sub); err != nil {
			os.RemoveAll(staged)
			return fmt.Errorf("stage %s into batch: %w", u.snippet.Frontmatter.ID, err)
		}
		fmt.Fprintf(&manifest, "%s\t%s\n", filepath.Join(strconv.Itoa(i), u.entrypoint), u.label)
	}
	if err := os.WriteFile(filepath.Join(batchDir, "manifest.tsv"), []byte(manifest.String()), 0o644); err != nil {
		return err
	}

	first := units[0]
	fmt.Fprintf(out, "--- validate %s batch (%d snippets, runtime=%s) ---\n",
		first.snippet.Frontmatter.SDK, len(units), first.runtime)

	switch first.runner.Mode {
	case "docker":
		tag, err := validatorImageTag(cfg.ValidatorsDir, first.runnerDir, first.runner.ImagePrefix)
		if err != nil {
			return err
		}
		return runContainerBatch(tag, batchDir, first.env, first.extraEnv, out)
	case "native":
		return runNativeBatch(first.runnerDir, batchDir, first.env, first.extraEnv, out)
	default:
		return fmt.Errorf("validator runtime %q: unknown mode %q", first.runtime, first.runner.Mode)
	}
}

// runContainerBatch runs the harness once with the batch dir mounted at
// /snippet. SNIPPET_BATCH points the harness at the manifest; the harness
// loops rather than reading a single SNIPPET_ENTRYPOINT.
func runContainerBatch(tag, batchDir string, env envInputs, extraEnv map[string]string, out io.Writer) error {
	args := []string{"run", "--rm",
		"-v", batchDir + ":/snippet:ro",
		"-e", "SNIPPET_BATCH=/snippet/manifest.tsv",
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
		return fmt.Errorf("batch validation failed: %w", err)
	}
	return nil
}

// runNativeBatch execs the harness once on the host with the batch dir as
// $SNIPPET_DIR and $SNIPPET_BATCH pointing at the manifest.
func runNativeBatch(runnerDir, batchDir string, env envInputs, extraEnv map[string]string, out io.Writer) error {
	script := filepath.Join(runnerDir, "harness", "run.sh")
	if _, err := os.Stat(script); err != nil {
		return fmt.Errorf("native validator run.sh not found at %s: %w", script, err)
	}
	cmd := exec.Command("/bin/sh", script)
	cmd.Stdout = out
	cmd.Stderr = out
	cmd.Env = append(os.Environ(),
		"SNIPPET_DIR="+batchDir,
		"SNIPPET_BATCH="+filepath.Join(batchDir, "manifest.tsv"),
	)
	cmd.Env = append(cmd.Env, envForRun(env)...)
	for k, v := range extraEnv {
		cmd.Env = append(cmd.Env, k+"="+v)
	}
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("native batch validator failed: %w", err)
	}
	return nil
}
