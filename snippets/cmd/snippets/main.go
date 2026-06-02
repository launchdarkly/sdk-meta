package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/launchdarkly/sdk-meta/snippets"
	"github.com/launchdarkly/sdk-meta/snippets/internal/adapters/ldapplication"
	"github.com/launchdarkly/sdk-meta/snippets/internal/adapters/lddocs"
	"github.com/launchdarkly/sdk-meta/snippets/internal/adapters/rawfiles"
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
  snippets render --target=<target> --entrypoint=<dir> [--entrypoint=<dir2> ...] [--sdks=./sdks]
      Walks each --entrypoint directory in the consumer checkout, finds files
      that contain SDK_SNIPPET:RENDER markers, and rewrites each marked
      region from the snippet sources. --entrypoint may be passed multiple
      times. --target selects the adapter:
        ld-application — rewrites JSX children inside marked components
                         (gonfalon, the LaunchDarkly app)
        ld-docs        — rewrites fenced code-block bodies in MDX
                         (ld-docs-private, ld-docs)

  snippets render --target=raw-files --manifest=<path> [--consumer=<dir>] [--sdks=./sdks]
      Reads a YAML manifest enumerating (snippet-id, output-path) pairs and
      writes each rendered body to <consumer>/<manifest.out>/<entry.path>.
      Used by consumers that import snippet text via Vite's "?raw" import
      pattern (e.g. gonfalon's packages/sdk-info/) where the marker-driven
      flow doesn't apply. --consumer defaults to the manifest's directory.

  snippets verify --target=<target> --entrypoint=<dir> [--entrypoint=<dir2> ...] [--sdks=./sdks]
      Read-only counterpart to render, used by CI in the consumer repo.
      Fails if the rendered bytes would drift from what's on disk, or if
      a marker's hash does not match its current region's content. Never
      writes; never executes any snippet code.

  snippets validate --sdk=<sdk-id> [--snippet=<id>] [--snippet-skip=<id>]
                    [--group=<sdk-info|sdk-docs>] [--sdks=./sdks] [--validators=./validators]
      Builds the SDK's per-language validator (Docker image or native harness),
      stages each runnable snippet with concrete input values, and runs it
      against a real LaunchDarkly environment. Exercises the snippet code end
      to end; requires LAUNCHDARKLY_SDK_KEY (or _MOBILE_KEY / _CLIENT_SIDE_ID)
      and LAUNCHDARKLY_FLAG_KEY in the env. Pass --snippet=<id> to validate
      a single snippet (useful while developing scaffolds). --group filters
      to one snippet group (the middle segment of the snippet id), which is
      how CI splits one SDK across multiple matrix rows.

  snippets image-tag --runtime=<runtime> [--validators=./validators]
      Print the deterministic Docker image tag the validate subcommand
      will build for the given validator runtime (e.g. python,
      cpp-server). The tag is <image-prefix>:<content-hash-16-hex>,
      computed over validators/shared/ plus validators/languages/<runtime>/.
      Empty output (no error) for runtimes whose mode is native. Used by
      CI to pre-build images with shared layer caches (docker buildx
      --cache-from type=gha) under the same tag the validator will later
      look up — keeps cold-cache CI builds out of the validate step.

  snippets list-runners --sdk=<sdk-id> [--group=<group>] [--sdks=./sdks]
                        [--validators=./validators]
      Print one runtime name per line — every distinct validator runtime
      that this SDK's bound snippets will exercise under the given group
      filter. Resolves scaffold-bound snippets the same way validate
      does (scaffold's runtime, falling back to the wrappee's lang). CI's
      per-row pre-build step uses this to enumerate the Dockerfiles it
      needs to warm before validate runs.

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
	case "image-tag":
		runImageTag(os.Args[2:])
	case "list-runners":
		runListRunners(os.Args[2:])
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
	target := fset.String("target", "", "adapter target: `ld-application`, `ld-docs`, or `raw-files`")
	var entrypoints repeatableString
	fset.Var(&entrypoints, "entrypoint", "directory in the consumer checkout to walk for marker files (repeatable; ld-application and ld-docs only)")
	manifest := fset.String("manifest", "", "path to a raw-files manifest YAML (raw-files only)")
	consumer := fset.String("consumer", "", "consumer-checkout root that the manifest's `out:` resolves against (default: manifest's directory)")
	sdks := fset.String("sdks", "", "path to a sdks/ directory (default: embedded)")
	_ = fset.Parse(args)

	switch *target {
	case "ld-application":
		if len(entrypoints) == 0 {
			fmt.Fprintf(os.Stderr, "render: --target=ld-application requires at least one --entrypoint\n")
			os.Exit(2)
		}
		changed, err := ldapplication.Render(resolveSDKsFS(*sdks), entrypoints)
		if err != nil {
			fmt.Fprintf(os.Stderr, "render failed: %v\n", err)
			os.Exit(1)
		}
		printRenderResult(changed)
	case "ld-docs":
		if len(entrypoints) == 0 {
			fmt.Fprintf(os.Stderr, "render: --target=ld-docs requires at least one --entrypoint\n")
			os.Exit(2)
		}
		changed, err := lddocs.Render(resolveSDKsFS(*sdks), entrypoints)
		if err != nil {
			fmt.Fprintf(os.Stderr, "render failed: %v\n", err)
			os.Exit(1)
		}
		printRenderResult(changed)
	case "raw-files":
		if *manifest == "" {
			fmt.Fprintf(os.Stderr, "render: --target=raw-files requires --manifest\n")
			os.Exit(2)
		}
		written, err := rawfiles.Render(resolveSDKsFS(*sdks), *manifest, *consumer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "render failed: %v\n", err)
			os.Exit(1)
		}
		if len(written) == 0 {
			fmt.Println("no files written")
			return
		}
		for _, p := range written {
			fmt.Println("wrote", p)
		}
	default:
		fmt.Fprintf(os.Stderr, "render: --target must be `ld-application`, `ld-docs`, or `raw-files` (got %q)\n", *target)
		os.Exit(2)
	}
}

func printRenderResult(changed []string) {
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
	target := fset.String("target", "", "adapter target: `ld-application` or `ld-docs`")
	var entrypoints repeatableString
	fset.Var(&entrypoints, "entrypoint", "directory in the consumer checkout to walk for marker files (repeatable)")
	sdks := fset.String("sdks", "", "path to a sdks/ directory (default: embedded)")
	_ = fset.Parse(args)

	if len(entrypoints) == 0 {
		fmt.Fprintf(os.Stderr, "verify: at least one --entrypoint is required\n")
		os.Exit(2)
	}
	var err error
	switch *target {
	case "ld-application":
		err = ldapplication.Verify(resolveSDKsFS(*sdks), entrypoints)
	case "ld-docs":
		err = lddocs.Verify(resolveSDKsFS(*sdks), entrypoints)
	default:
		fmt.Fprintf(os.Stderr, "verify: --target must be `ld-application` or `ld-docs`\n")
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "verify failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("ok")
}

func runValidate(args []string) {
	fset := flag.NewFlagSet("validate", flag.ExitOnError)
	sdk := fset.String("sdk", "", "sdk id to validate (required)")
	snippet := fset.String("snippet", "", "snippet id to validate (optional; restricts to one snippet)")
	snippetSkip := fset.String("snippet-skip", "", "snippet id to skip (optional; useful for splitting one SDK across CI rows)")
	group := fset.String("group", "", "snippet group to validate (optional; e.g. `sdk-info` or `sdk-docs`)")
	sdks := fset.String("sdks", "", "path to a sdks/ directory (default: embedded)")
	validators := fset.String("validators", "./validators", "path to the validators/ directory")
	_ = fset.Parse(args)

	if *sdk == "" {
		fmt.Fprintf(os.Stderr, "validate: --sdk is required\n")
		os.Exit(2)
	}
	if err := validate.Run(validate.Config{
		SDKsFS:        resolveSDKsFS(*sdks),
		ValidatorsDir: *validators,
		SDK:           *sdk,
		Snippet:       *snippet,
		SnippetSkip:   *snippetSkip,
		Group:         *group,
	}); err != nil {
		fmt.Fprintf(os.Stderr, "validate failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("ok")
}

// runImageTag prints the deterministic Docker tag for a given validator
// runtime. CI uses this to pre-build images with `docker buildx build
// --cache-from type=gha --load -t $(snippets image-tag ...)`, and the
// validator's subsequent `docker build -t <same-tag>` finds the layers in
// the local Docker daemon's cache — cold-cache cost moves out of the
// validate step.
func runImageTag(args []string) {
	fset := flag.NewFlagSet("image-tag", flag.ExitOnError)
	runtime := fset.String("runtime", "", "validator runtime name (required; matches a directory under validators/languages/)")
	validators := fset.String("validators", "./validators", "path to the validators/ directory")
	_ = fset.Parse(args)
	if *runtime == "" {
		fmt.Fprintf(os.Stderr, "image-tag: --runtime is required\n")
		os.Exit(2)
	}
	tag, err := validate.ImageTag(*validators, *runtime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "image-tag failed: %v\n", err)
		os.Exit(1)
	}
	// Native validators have no Docker image. Print nothing (no newline)
	// so CI's shell can branch on an empty result.
	if tag == "" {
		return
	}
	fmt.Println(tag)
}

// runListRunners prints, one per line, every distinct validator runtime
// the SDK's bound snippets will exercise under the given group filter.
// CI's per-row pre-build step loops over the output to warm caches before
// `validate` runs.
func runListRunners(args []string) {
	fset := flag.NewFlagSet("list-runners", flag.ExitOnError)
	sdk := fset.String("sdk", "", "sdk id (required)")
	group := fset.String("group", "", "snippet group (optional; e.g. sdk-info or sdk-docs)")
	sdks := fset.String("sdks", "", "path to a sdks/ directory (default: embedded)")
	validators := fset.String("validators", "./validators", "path to the validators/ directory")
	_ = fset.Parse(args)
	if *sdk == "" {
		fmt.Fprintf(os.Stderr, "list-runners: --sdk is required\n")
		os.Exit(2)
	}
	runners, err := validate.ListRunners(resolveSDKsFS(*sdks), *validators, *sdk, *group)
	if err != nil {
		fmt.Fprintf(os.Stderr, "list-runners failed: %v\n", err)
		os.Exit(1)
	}
	for _, r := range runners {
		fmt.Println(r)
	}
}
