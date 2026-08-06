// Command genspecs builds the spec-support pipeline that lives alongside
// sdk-meta's existing data products.
//
// Subcommands:
//
//	sync-repos   Clone (shallow) any repo from products/repos.json that isn't
//	             present under --sdk-repos-root, plus sdk-specs and
//	             sdk-test-harness. Fast-forward existing clones when safe.
//	catalog      Walk sdk-specs and emit products/specs.json.
//	harness      Read sdk-test-harness + each SDK's testharness-suppressions
//	             files and emit products/harness_signals.json.
//	judge        Apply the cheap applies-to filter, then call the configured
//	             LLM provider for every remaining (sdk, spec) cell. Emits
//	             products/spec_support.json.
//	html         Render the matrix and per-SDK HTML views into _site/.
package main

import (
	"fmt"
	"os"
)

const promptVersion = "v1"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	var err error
	switch cmd {
	case "sync-repos":
		err = runSync(args)
	case "catalog":
		err = runCatalog(args)
	case "harness":
		err = runHarness(args)
	case "judge":
		err = runJudge(args)
	case "html":
		err = runHTML(args)
	case "-h", "--help", "help":
		usage()
		return
	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand %q\n\n", cmd)
		usage()
		os.Exit(2)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `Usage: genspecs <subcommand> [flags]

Subcommands:
  sync-repos   Clone missing SDK repos and fast-forward existing ones.
  catalog      Generate products/specs.json from the sdk-specs repo.
  harness      Generate products/harness_signals.json from the sdk-test-harness
               repo and each SDK's testharness-suppressions files.
  judge        Generate products/spec_support.json by classifying each
               (sdk, spec) pair, using an LLM where the applies-to filter
               doesn't trivially decide it.
  html         Render _site/spec-support.html and _site/spec-support-by-sdk.html.

Run "genspecs <subcommand> -h" for per-subcommand flags.`)
}
