package validate

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Runner is the parsed contents of a validator's runner.yaml. Each
// validators/languages/<runtime>/ directory carries one.
type Runner struct {
	// Mode controls how the harness runs:
	//   "docker" — the Go validator builds the Dockerfile and runs
	//              harness/run.sh inside the resulting container with the
	//              staged snippet bind-mounted at /snippet.
	//   "native" — the Go validator execs harness/run.sh on the host with
	//              the staged snippet path passed as $SNIPPET_DIR.
	Mode string `yaml:"mode"`

	// RunsOn is a hint for the CI workflow's matrix.runs-on. Not read by
	// the Go validator; included so the workflow file and the runner
	// descriptor stay in sync visually.
	RunsOn string `yaml:"runs-on"`

	// ImagePrefix is the Docker image-tag prefix used for `mode: docker`.
	// The full tag is `<image-prefix>:<content-hash>`.
	ImagePrefix string `yaml:"image-prefix"`

	// Batch opts the validator into batch mode. When true the Go runner
	// stages every matching snippet at once and invokes the harness a
	// handful of times (one per worker shard) rather than once per
	// snippet, handing each invocation a manifest of staged snippets to
	// loop over inside a single warm workspace. The harness's run.sh must
	// understand the batch contract (read $SNIPPET_BATCH manifest, loop,
	// aggregate pass/fail). Validators that haven't been migrated leave
	// this false and keep the one-invocation-per-snippet path.
	Batch bool `yaml:"batch"`
}

func loadRunner(validatorsDir, runtime string) (*Runner, string, error) {
	dir := filepath.Join(validatorsDir, "languages", runtime)
	rp := filepath.Join(dir, "runner.yaml")
	raw, err := os.ReadFile(rp)
	if err != nil {
		return nil, "", fmt.Errorf("validator runner.yaml not found for runtime %q at %s: %w", runtime, rp, err)
	}
	var r Runner
	dec := yaml.NewDecoder(bytes.NewReader(raw))
	dec.KnownFields(true)
	if err := dec.Decode(&r); err != nil {
		return nil, "", fmt.Errorf("%s: %w", rp, err)
	}
	if r.Mode != "docker" && r.Mode != "native" {
		return nil, "", fmt.Errorf("%s: mode must be `docker` or `native`, got %q", rp, r.Mode)
	}
	if r.Mode == "docker" && r.ImagePrefix == "" {
		return nil, "", fmt.Errorf("%s: mode=docker requires image-prefix", rp)
	}
	return &r, dir, nil
}
