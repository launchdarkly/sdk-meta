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
fallbackValue := ldai.NewConfig().Build() // by default, the Config is disabled

cfg, tracker := aiClient.Config("example-config-key", context, fallbackValue, map[string]interface{}{"exampleCustomVariable": "exampleCustomValue"})

if cfg.Enabled() {

  // Send a request to your AI provider, using details from the customized cfg

} else {

  // Application path to take when the cfg is disabled

}

```
