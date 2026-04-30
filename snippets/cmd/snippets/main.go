package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/launchdarkly/sdk-meta/snippets"
	"github.com/launchdarkly/sdk-meta/snippets/internal/adapters/ldapplication"
	"github.com/launchdarkly/sdk-meta/snippets/internal/validate"
	"github.com/launchdarkly/sdk-meta/snippets/internal/version"
)

// resolveSDKsFS returns the embedded sdks tree when sdksFlag is empty, or
// os.DirFS(sdksFlag) otherwise. The embedded form is what release artifacts
// ship with — consumers don't need to check out sdk-meta to render. Authors
// developing locally pass `--sdks=./sdks` to point at their working tree.
func resolveSDKsFS(sdksFlag string) fs.FS {
	if sdksFlag == "" {
		return snippets.SDKsFS()
	}
	return os.DirFS(sdksFlag)
}

// repeatableString implements flag.Value so a flag like --entrypoint can
// be passed multiple times and accumulate into a slice. Mirrors the
// stringSliceFlag idiom widely used in Go CLIs that stick to the stdlib
// `flag` package.
type repeatableString []string

func (r *repeatableString) String() string     { return strings.Join(*r, ",") }
func (r *repeatableString) Set(s string) error { *r = append(*r, s); return nil }

const usage = `snippets — LaunchDarkly SDK snippet generator

usage:
  snippets render --target=ld-application --entrypoint=<dir> [--entrypoint=<dir2> ...] [--sdks=./sdks]
      Walks each --entrypoint directory in the consumer checkout, finds files
      that contain SDK_SNIPPET:RENDER markers, and rewrites each marked
      region from the snippet sources. --entrypoint may be passed multiple
      times. Authors run this after editing a snippet; the consumer's
      sync action runs it on every release.

  snippets verify --target=ld-application --entrypoint=<dir> [--entrypoint=<dir2> ...] [--sdks=./sdks]
      Read-only counterpart to render, used by CI in the consumer repo.
      Fails if the rendered bytes would drift from what's on disk, or if
      a marker's hash does not match its current region's content. Never
      writes; never executes any snippet code.

  snippets validate --sdk=<sdk-id> [--sdks=./sdks] [--validators=./validators]
      Builds the SDK's per-language validator (Docker image or native harness),
      stages each runnable snippet with concrete input values, and runs it
      against a real LaunchDarkly environment. Exercises the snippet code end
      to end; requires LAUNCHDARKLY_SDK_KEY (or _MOBILE_KEY / _CLIENT_SIDE_ID)
      and LAUNCHDARKLY_FLAG_KEY in the env.

  snippets version
      Print the snippets generator version.
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
	switch os.Args[1] {
	case "render":
		runRender(os.Args[2:])
	case "verify":
		runVerify(os.Args[2:])
	case "validate":
		runValidate(os.Args[2:])
	case "version":
		fmt.Println(version.Version)
	case "-h", "--help", "help":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand %q\n\n%s", os.Args[1], usage)
		os.Exit(2)
	}
}

func runRender(args []string) {
	fset := flag.NewFlagSet("render", flag.ExitOnError)
	target := fset.String("target", "", "adapter target: `ld-application`")
	var entrypoints repeatableString
	fset.Var(&entrypoints, "entrypoint", "directory in the consumer checkout to walk for marker files (repeatable)")
	sdks := fset.String("sdks", "", "path to a sdks/ directory (default: embedded)")
	_ = fset.Parse(args)

	if *target != "ld-application" || len(entrypoints) == 0 {
		fmt.Fprintf(os.Stderr, "render: --target=ld-application and at least one --entrypoint are required\n")
		os.Exit(2)
	}
	changed, err := ldapplication.Render(resolveSDKsFS(*sdks), entrypoints)
	if err != nil {
		fmt.Fprintf(os.Stderr, "render failed: %v\n", err)
		os.Exit(1)
	}
	if len(changed) == 0 {
		fmt.Println("no changes")
		return
	}
	for _, p := range changed {
		fmt.Println("rewrote", p)
	}
}

func runVerify(args []string) {
	fset := flag.NewFlagSet("verify", flag.ExitOnError)
	target := fset.String("target", "", "adapter target: `ld-application`")
	var entrypoints repeatableString
	fset.Var(&entrypoints, "entrypoint", "directory in the consumer checkout to walk for marker files (repeatable)")
	sdks := fset.String("sdks", "", "path to a sdks/ directory (default: embedded)")
	_ = fset.Parse(args)

	if *target != "ld-application" || len(entrypoints) == 0 {
		fmt.Fprintf(os.Stderr, "verify: --target=ld-application and at least one --entrypoint are required\n")
		os.Exit(2)
	}
	if err := ldapplication.Verify(resolveSDKsFS(*sdks), entrypoints); err != nil {
		fmt.Fprintf(os.Stderr, "verify failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("ok")
}

func runValidate(args []string) {
	fset := flag.NewFlagSet("validate", flag.ExitOnError)
	sdk := fset.String("sdk", "", "sdk id to validate (required)")
	sdks := fset.String("sdks", "", "path to a sdks/ directory (default: embedded)")
	validators := fset.String("validators", "./validators", "path to the validators/ directory")
	_ = fset.Parse(args)

	if *sdk == "" {
		fmt.Fprintf(os.Stderr, "validate: --sdk is required\n")
		os.Exit(2)
	}
	if err := validate.Run(validate.Config{SDKsFS: resolveSDKsFS(*sdks), ValidatorsDir: *validators, SDK: *sdk}); err != nil {
		fmt.Fprintf(os.Stderr, "validate failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("ok")
}
