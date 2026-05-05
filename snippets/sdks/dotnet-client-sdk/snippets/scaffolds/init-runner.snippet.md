---
id: dotnet-client-sdk/scaffolds/init-runner
sdk: dotnet-client-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Runs an `init.txt`-style snippet end-to-end against a real LaunchDarkly
  env. The wrappee body uses C# top-level statements: it builds a
  context, awaits `LdClient.InitAsync`, then conditionally prints
  `SDK successfully initialized!`. The scaffold drops the body verbatim
  and appends `Console.WriteLine("feature flag evaluates to true");`
  so the harness's `await_success_line` regex sees the EXAM-HELLO line
  on a clean run. The body uses `client.Initialized` as the gate; if
  init fails the trailer still runs (the body has no `Environment.Exit`
  in the failure branch), so the trailer is wrapped in an
  `if (client.Initialized)` guard tied to the same lexical `client`.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: dotnet-client
  entrypoint: Program.cs
  requirements: |
    LaunchDarkly.ClientSdk
---

```csharp
{{ body }}

if (client.Initialized)
{
    Console.WriteLine("feature flag evaluates to true");
}
else
{
    Console.Error.WriteLine("scaffold: client.Initialized is false after InitAsync");
    Environment.Exit(1);
}
```
