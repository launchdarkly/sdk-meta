---
id: dotnet-server-sdk/scaffolds/init-runner
sdk: dotnet-server-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Runs an `init.txt`-style snippet end-to-end against a real LaunchDarkly
  env. The wrappee body uses C# top-level statements that include
  `WebApplication.CreateBuilder(args)` (gonfalon's docs target ASP.NET
  Core hosting) plus a regular `LdClient` initialization.

  Layout:
    - Program.cs (this scaffold) — the snippet body verbatim, with the
      literal SDK-key placeholder substituted by `validation.placeholders`.
      We append a final `Console.WriteLine("feature flag evaluates to true")`
      after the body so the harness's `await_success_line` regex sees the
      EXAM-HELLO line on a clean run. The body's failure branch calls
      `Environment.Exit(1)` so a failed init never reaches the trailer.
    - HelloDotNet.csproj (companion) — overrides the harness's default
      console-only `.csproj` with `Microsoft.NET.Sdk.Web`, so
      `WebApplication.CreateBuilder` resolves at compile time.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: dotnet-server
  entrypoint: Program.cs
  companions:
    - dotnet-server-sdk/scaffolds/init-runner-csproj
  requirements: |
    LaunchDarkly.ServerSdk
---

```csharp
{{ body }}

Console.WriteLine("feature flag evaluates to true");
```
