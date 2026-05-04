---
id: go-server-sdk/sdk-info/init
sdk: go-server-sdk
kind: init
lang: go
file: go-server-sdk/init.txt
description: Client initialization snippet for go-server-sdk.
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
	// This is your LaunchDarkly SDK key.
	// Never hardcode your SDK key in production.
	ldClient, _ := ld.MakeClient("YOUR_SDK_KEY", 5*time.Second)
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
