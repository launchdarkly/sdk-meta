---
id: dotnet-client-sdk/scaffolds/csharp-client-syntax-only-v3
sdk: dotnet-client-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Parse-and-type-check validator for .NET client SDK doc fragments
  that target the v3.x API surface (e.g.
  `LdClient.InitAsync(mobileKey, context)`, `Context.New(...)`).

  Routes through the same `dotnet-client` validator container as the
  current-version scaffold, but pins `LaunchDarkly.ClientSdk` to a
  v3.x release via the `Package==Version` syntax the harness's
  requirements-reader recognizes. The current-version
  `csharp-client-syntax-only` scaffold floats to the latest
  ClientSdk, whose overload set no longer matches the v3.x
  invocations the docs cover.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: dotnet-client
  entrypoint: Program.cs
  requirements: |
    LaunchDarkly.ClientSdk==3.1.0
---

```csharp
using LaunchDarkly.Sdk;
using LaunchDarkly.Sdk.Client;
using System;
using LaunchDarkly.Sdk.Client.Integrations;
// USING_LIFT_MARKER

namespace LaunchDarklySnippet
{
    public class Program
    {
        // `client` is `dynamic` so v3 init shapes (which return
        // `LdClient` directly from `LdClient.InitAsync(string, Context)`)
        // assign without a type cast — v3, v4, and v5 returned subtly
        // different shapes, and the scaffold is parse-only.
        #pragma warning disable CS8625, CS0414, CS0649
        private static dynamic client = null;
        #pragma warning restore CS8625, CS0414, CS0649

        public static void Main(string[] args)
        {
            System.Console.WriteLine("feature flag evaluates to true");
        }

        #pragma warning disable CS0162
        private async System.Threading.Tasks.Task Wrappee()
        {
            try {
{{ body }}
            } catch (System.Exception) { /* never reached */ }
            await System.Threading.Tasks.Task.CompletedTask;
        }
        #pragma warning restore CS0162
    }
}
```
