---
id: dotnet-server-sdk/sdk-docs/features/agentcontrol/customize-config
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Customize an AgentControl config for .NET AI.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
var fallbackConfig = LdAiCompletionConfigDefault.New()
  .SetEnabled(false)
  .Build();

var config = aiClient.Config(
  "example-config-key",
  context,
  fallbackConfig,
  new Dictionary<string, object> {
    { "exampleCustomVariable", "exampleCustomValue" }
  }
);

if (config.Enabled == true) {

  // Send a request to your AI provider, using the details from the customized config

} else {

  // Application path to take when the config is disabled

}

```
