---
id: dotnet-server-sdk/scaffolds/csharp-syntax-only-typed
sdk: dotnet-server-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Parse-only validator for .NET server SDK doc fragments that need a
  statically typed `client`.

  The default csharp-syntax-only scaffold types `client` as `dynamic`
  so one stub serves both the v6 `User` and v7+ `Context` API
  surfaces, but C# rejects a lambda used as an operand of a
  dynamically dispatched operation (error CS1977) — event-handler
  fragments such as
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
  runtime: dotnet-server
  entrypoint: Program.cs
  requirements: |
    LaunchDarkly.ServerSdk
---

```csharp
// USING_LIFT_MARKER
using System;
using LaunchDarkly.Sdk;
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.Sdk.Server.Migrations;

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

        public static void Main(string[] args)
        {
            System.Console.WriteLine("feature flag evaluates to true");
        }

        private void Wrappee()
        {
            try {
{{ body }}
            } catch (System.Exception) { /* never reached */ }
        }
    }
}
```
