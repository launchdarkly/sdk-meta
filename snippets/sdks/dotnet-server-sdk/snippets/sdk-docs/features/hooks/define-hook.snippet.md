---
id: dotnet-server-sdk/sdk-docs/features/hooks/define-hook
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Hook implementation and configuration for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
using LaunchDarkly.Sdk.Server.Hooks;

public class ExampleHook : Hook {

  public ExampleHook() : base("example-hook") { }

  // Implement at least one of `BeforeEvaluation`, `AfterEvaluation`

  // `BeforeEvaluation` is called during the execution of a variation method
  // before the flag value has been determined

  // `AfterEvaluation` is called during the execution of a variation method
  // after the flag value has been determined

}

var exampleHook = new ExampleHook();

var config = Configuration.Builder("YOUR_SDK_KEY")
  .Hooks(Components.Hooks()
    .Add(exampleHook)
  ).Build();

var client = new LdClient(config);
```
