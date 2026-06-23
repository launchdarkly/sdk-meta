---
id: dotnet-client-sdk/scaffolds/csharp-client-syntax-only-v4
sdk: dotnet-client-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Parse-and-type-check validator for .NET client SDK doc fragments
  that target the v4.x API surface (e.g.
  `Configuration.Builder(mobileKey, ConfigurationBuilder.AutoEnvAttributes.Enabled)`
  with the v4-era component builders).

  Routes through the same `dotnet-client` validator container as the
  current-version scaffold, but pins `LaunchDarkly.ClientSdk` to a
  v4.x release via the `Package==Version` syntax the harness's
  requirements-reader recognizes, so v4.x-titled docs validate
  against the actual v4 overload set rather than whatever the
  floating scaffold currently resolves.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: dotnet-client
  entrypoint: Program.cs
  requirements: |
    LaunchDarkly.ClientSdk==4.0.0
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
        // `client` is `dynamic` so the doc fragments' init shapes
        // assign without a type cast — the scaffold is parse-only and
        // the body is never invoked.
        #pragma warning disable CS8625, CS0414, CS0649
        private static dynamic client = null;
        // Evaluation/init fragments pass an ambient `context`; the
        // docs assume it already exists.
        private static dynamic context = null;
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
