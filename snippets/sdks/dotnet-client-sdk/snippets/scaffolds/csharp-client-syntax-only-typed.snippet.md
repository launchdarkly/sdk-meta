---
id: dotnet-client-sdk/scaffolds/csharp-client-syntax-only-typed
sdk: dotnet-client-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Real-typed sibling of `csharp-client-syntax-only`. That scaffold
  stubs `client` as `dynamic`, which breaks for event-pattern bodies:
  C# forbids using a lambda expression as an operand of a dynamically
  dispatched operation, so
  `client.FlagTracker.FlagValueChanged += (sender, eventArgs) => ...`
  cannot compile against a dynamic stub. This variant types the stub
  as the real `LdClient` so the event subscription type-checks against
  the SDK's `IFlagTracker` surface.
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
// USING_LIFT_MARKER

namespace LaunchDarklySnippet
{
    public class Program
    {
        // Real-typed client stub so the body's
        // `client.FlagTracker.FlagValueChanged += ...` resolves through
        // the SDK's actual interfaces. Never invoked at runtime — Main
        // below short-circuits.
        #pragma warning disable CS0649
        private static LdClient client;
        #pragma warning restore CS0649

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
