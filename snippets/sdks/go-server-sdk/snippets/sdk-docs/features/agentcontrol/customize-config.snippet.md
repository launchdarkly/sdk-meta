---
id: go-server-sdk/sdk-docs/features/agentcontrol/customize-config
sdk: go-server-sdk
kind: reference
lang: go
description: Customize an AgentControl config for Go AI.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
defaultValue := ldai.NewAICompletionConfigDefault().Disabled() // used when the config can't be evaluated

cfg := aiClient.CompletionConfig("example-config-key", context, defaultValue, map[string]interface{}{"exampleCustomVariable": "exampleCustomValue"})
tracker := cfg.CreateTracker()

if cfg.Enabled() {

  // Send a request to your AI provider, using details from the customized cfg

} else {

  // Application path to take when the cfg is disabled

}

```
