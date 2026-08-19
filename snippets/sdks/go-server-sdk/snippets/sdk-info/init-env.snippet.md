---
id: go-server-sdk/sdk-info/init-env
sdk: go-server-sdk
kind: init
lang: go
file: go-server-sdk/init-env.txt
description: Client initialization snippet for go-server-sdk reading the SDK key from an environment variable.
validation:
  scaffold: go-server-sdk/scaffolds/init-runner
---

```go
package main

import (
	"fmt"
	"os"
	"time"

	ld "github.com/launchdarkly/go-server-sdk/v7"
)

func main() {
	// Reads your LaunchDarkly SDK key from the LAUNCHDARKLY_SDK_KEY
	// environment variable, so one build can run in every environment.
	ldClient, _ := ld.MakeClient(os.Getenv("LAUNCHDARKLY_SDK_KEY"), 5*time.Second)
	if ldClient.Initialized() {
		fmt.Printf("SDK successfully initialized!")
	} else {
		fmt.Printf("SDK failed to initialize")
		os.Exit(1)
	}

	// For onboarding purposes only we flush events as soon as
	// possible so we quickly detect your connection.
	// You don't have to do this in practice because events are automatically flushed.
	ldClient.Flush()
}
```
