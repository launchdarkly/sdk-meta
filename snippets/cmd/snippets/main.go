package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/launchdarkly/sdk-meta/snippets/internal/adapters/ldapplication"
	"github.com/launchdarkly/sdk-meta/snippets/internal/validate"
	"github.com/launchdarkly/sdk-meta/snippets/internal/version"
)

const usage = `snippets — LaunchDarkly SDK snippet generator

usage:
  snippets render   --target=ld-application --out=<app-checkout> [--sdks=./sdks]
  snippets verify   --target=ld-application --out=<app-checkout> [--sdks=./sdks]
  snippets validate --sdk=<sdk-id> [--sdks=./sdks] [--validators=./validators]
  snippets version

First-pass support: target=ld-application, sdk=python-server-sdk.
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
	fs := flag.NewFlagSet("render", flag.ExitOnError)
	target := fs.String("target", "", "adapter target: `ld-application`")
	out := fs.String("out", "", "path to the consumer checkout")
	sdks := fs.String("sdks", "./sdks", "path to the sdks/ directory")
	_ = fs.Parse(args)

	if *target != "ld-application" || *out == "" {
		fmt.Fprintf(os.Stderr, "render: --target=ld-application and --out are required\n")
		os.Exit(2)
	}
	changed, err := ldapplication.Render(*sdks, *out)
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
	fs := flag.NewFlagSet("verify", flag.ExitOnError)
	target := fs.String("target", "", "adapter target: `ld-application`")
	out := fs.String("out", "", "path to the consumer checkout")
	sdks := fs.String("sdks", "./sdks", "path to the sdks/ directory")
	_ = fs.Parse(args)

	if *target != "ld-application" || *out == "" {
		fmt.Fprintf(os.Stderr, "verify: --target=ld-application and --out are required\n")
		os.Exit(2)
	}
	if err := ldapplication.Verify(*sdks, *out); err != nil {
		fmt.Fprintf(os.Stderr, "verify failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("ok")
}

func runValidate(args []string) {
	fs := flag.NewFlagSet("validate", flag.ExitOnError)
	sdk := fs.String("sdk", "", "sdk id to validate (required)")
	sdks := fs.String("sdks", "./sdks", "path to the sdks/ directory")
	validators := fs.String("validators", "./validators", "path to the validators/ directory")
	_ = fs.Parse(args)

	if *sdk == "" {
		fmt.Fprintf(os.Stderr, "validate: --sdk is required\n")
		os.Exit(2)
	}
	if err := validate.Run(validate.Config{SDKsDir: *sdks, ValidatorsDir: *validators, SDK: *sdk}); err != nil {
		fmt.Fprintf(os.Stderr, "validate failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("ok")
}
