---
id: go-server-sdk/ai-configs/initialize
sdk: go-server-sdk
kind: initialize
lang: go
file: go-server-sdk/ai-configs/initialize.txt
description: Initialize the LaunchDarkly client and AI Configs client for go-server-sdk.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
client, _ := ld.MakeClient("{{sdkkey}}", 5*time.Second) // This is your SDK key
aiClient, _ := ldai.NewClient(client)
```
