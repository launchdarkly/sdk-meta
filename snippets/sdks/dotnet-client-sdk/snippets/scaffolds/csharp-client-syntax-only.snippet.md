---
id: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
sdk: dotnet-client-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Parse-only validator for .NET client SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: dotnet-client
  entrypoint: Program.cs
  # The harness reads requirements.txt as a list of `dotnet add package`
  # arguments — one per line.
  requirements: |
    LaunchDarkly.ClientSdk
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
        // Stub fields so the wrappee body's `client.BoolVariation(...)`,
        // `client.StringVariation(...)`, etc. resolve at compile time.
        // Typing `client` as `dynamic` makes the body forward-compatible
        // across .NET client SDK versions whose APIs differ slightly
        // (LdClient.Init signatures changed across v3 / v4 / v5). Never
        // invoked at runtime — Main below short-circuits.
        #pragma warning disable CS8625, CS0414, CS0649
        private static dynamic client = null;
        // Some config fragments reference a `context` binding (e.g.
        // LdClient.Init(config, context)); provide it as a stub.
        private static dynamic context = null;
        // Init fragments pass a `startWaitTime` the docs assume already
        // exists.
        private static System.TimeSpan startWaitTime = default;
        // Test-data fragments reference a `td` the docs assume an
        // earlier `TestData.DataSource()` call created. Typed as the
        // real TestData (not dynamic) so lambda arguments to
        // `VariationFunc(...)`-style builder calls keep compiling --
        // C# forbids lambdas in dynamically dispatched invocations.
        private static TestData td = null;
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
