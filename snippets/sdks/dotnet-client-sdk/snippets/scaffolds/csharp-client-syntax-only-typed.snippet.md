---
id: dotnet-client-sdk/scaffolds/csharp-client-syntax-only-typed
sdk: dotnet-client-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Parse-only validator for .NET client SDK doc fragments that need a
  statically typed `client`.

  The default csharp-client-syntax-only scaffold types `client` as
  `dynamic` so one stub serves several SDK eras, but C# rejects a
  lambda used as an operand of a dynamically dispatched operation
  (error CS1977) — event-handler fragments such as
  `client.DataSourceStatusProvider.StatusChanged += (sender, status) => …`
  cannot compile through a `dynamic` receiver. This variant types
  `client` as the real `LdClient` so `+=` binds against the actual
  `event EventHandler<DataSourceStatus>` member and the lambda's
  parameter types are inferred.
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
        // Statically typed stub so event-subscription bodies bind
        // against the real LdClient surface. Never invoked at
        // runtime — Main below short-circuits.
        #pragma warning disable CS8625, CS0414, CS0649
        private static LdClient client = null;
        #pragma warning restore CS8625, CS0414, CS0649

        // Caller-supplied helper the flag-changes body assumes exists.
        private static void DoSomethingWithNewFlagValue(bool newValue) {}

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
