---
id: dotnet-server-sdk/ai-configs/implementation
sdk: dotnet-server-sdk
kind: implementation
lang: csharp
file: dotnet-server-sdk/ai-configs/implementation.txt
description: Resolve an AI Config with a fallback for dotnet-server-sdk.
---

```csharp
var fallbackConfig = LdAiConfig.New()
  .SetModelName("my-default-model")
  .SetModelParam("temperature", LdValue.Of(0.8))
  .AddMessage("", Role.system)
  .SetModelProviderName("my-default-provider")
  .SetEnabled(true)
  .Build()

var tracker = aiClient.Config(
  "{{configKey}}",
  context,
  fallbackConfig,
  new Dictionary<string, object> {
    { "exampleCustomVariable", "exampleCustomValue" }
  }
);
```
