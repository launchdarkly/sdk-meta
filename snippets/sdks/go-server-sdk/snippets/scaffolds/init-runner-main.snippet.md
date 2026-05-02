---
id: go-server-sdk/scaffolds/init-runner-main
sdk: go-server-sdk
kind: scaffold
lang: go
file: main.go
description: |
  Runner half of the Go init scaffold pair. Staged as a companion of
  `init-runner` so it lands at `main.go` in the staging dir while the
  wrappee body sits at `wrappee/init.go`. `validation.entrypoint:
  main.go` on the parent scaffold makes the harness run this file.
---

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("go", "run", "wrappee/init.go")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "scaffold: wrappee exited with error: %v\n", err)
		fmt.Fprintf(os.Stderr, "wrappee stdout: %s\n", out)
		os.Exit(1)
	}
	if !strings.Contains(string(out), "SDK successfully initialized") {
		fmt.Fprintf(os.Stderr, "scaffold: wrappee did not print 'SDK successfully initialized'\n")
		fmt.Fprintf(os.Stderr, "wrappee stdout: %s\n", out)
		os.Exit(1)
	}
	fmt.Println("feature flag evaluates to true")
}
```
