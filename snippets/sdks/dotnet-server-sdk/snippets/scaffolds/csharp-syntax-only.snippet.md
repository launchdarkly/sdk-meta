---
id: dotnet-server-sdk/scaffolds/csharp-syntax-only
sdk: dotnet-server-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Parse-only validator for .NET server SDK doc fragments.

  C# requires `using` directives at the top of the file (or inside a
  namespace), not inside a method body. The `// USING_LIFT_MARKER`
  comment cues the harness pre-stage rewrite to lift any `using …;`
  lines from the wrappee body up to the marker, so doc fragments that
  show install-time `using LaunchDarkly.Sdk;` etc. can compile.

  C# has no local type declarations, so doc fragments that define a
  class alongside statements (e.g. a hook implementation followed by
  the configuration that registers it) cannot compile inside
  `Wrappee()`. The `// TYPE_LIFT_TARGET` comment cues a second harness
  pre-stage rewrite: brace-balanced type declarations found between
  the `// BODY_BEGIN` / `// BODY_END` markers are moved up to the
  target at namespace scope, where they compile as ordinary top-level
  types. Bodies without type declarations are untouched.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: dotnet-server
  entrypoint: Program.cs
  requirements: |
    LaunchDarkly.ServerSdk
    LaunchDarkly.Observability
    LaunchDarkly.ServerSdk.Ai
    LaunchDarkly.ServerSdk.Telemetry
    LaunchDarkly.ServerSdk.Redis
    LaunchDarkly.ServerSdk.DynamoDB
    LaunchDarkly.ServerSdk.Consul
---

```csharp
// USING_LIFT_MARKER
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
// Web-proxy fragments construct `NetworkCredential` unqualified while
// fully qualifying the other System.Net types; the docs assume the
// using directive is ambient, so provide it here.
using System.Net;
using LaunchDarkly.Sdk;
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.Sdk.Server.Migrations;
using LaunchDarkly.Sdk.Server.Integrations;
using LaunchDarkly.Sdk.Server.Ai;
using LaunchDarkly.Sdk.Server.Ai.Adapters;
using LaunchDarkly.Sdk.Server.Ai.Config;
using LaunchDarkly.Sdk.Server.Ai.Tracking;

namespace LaunchDarklySnippet
{
    // TYPE_LIFT_TARGET
    public class Program
    {
        // Stub fields the wrappee body refers to. Never used at runtime.
        // The body's `client.BoolVariation(...)` calls resolve through
        // these so the C# compiler is happy. v6 docs use `User` and
        // v6 overloads of variation methods, v7+ uses `Context` —
        // typing `client` as `dynamic` lets any overload resolve at
        // runtime so the same scaffold validates both API surfaces.
        // `aiClient` is dynamic so AI Config bodies that call
        // `aiClient.Config(...)` resolve without us pinning the
        // LdAiClient surface in the stub field type.
#pragma warning disable CS0414, CS0649
        private static dynamic client = null;
        // Docs use a `_client` field-naming convention in some fragments
        // (e.g. migration config); provide it alongside `client`.
        private static dynamic _client = null;
        private static dynamic aiClient = null;
        private static User user = null;
        private static Context context = default;
        // Evaluation fragments pass `myContext` to the variation
        // methods; the docs assume it already exists.
        private static Context myContext = default;
        // Persistent-store fragments reference a placeholder
        // `SomeDatabaseName` integration standing in for whichever
        // database package the reader installs; dynamic so its
        // `.DataStore()` call resolves without pinning a package.
        private static dynamic SomeDatabaseName = null;
        // AI metrics fragments call methods on the config's `tracker`
        // and read fields from a provider `response`; both come from
        // surrounding application code in the docs.
        private static dynamic response = null;
        // AI metrics flush fragments call Flush() on the underlying
        // LaunchDarkly client, which the docs name `baseClient`.
        private static dynamic baseClient = null;
        // Migration fragments reference an ambient migrator, payload,
        // op tracker, and the stage from a previous MigrationVariation
        // call; the docs assume they already exist.
        private static dynamic migration = null;
        private static dynamic payload = null;
        private static dynamic tracker = null;
        private static MigrationStage stage = default;
        // Test-data fragments reference a `td` the docs assume an
        // earlier `TestData.DataSource()` call created. Typed as the
        // real TestData (not dynamic) so lambda arguments to
        // `VariationFunc(...)`-style builder calls keep compiling --
        // C# forbids lambdas in dynamically dispatched invocations.
        private static TestData td = null;
        // Init fragments pass an `sdkKey` the docs assume already exists.
        private static string sdkKey = "";
        // The legacy aliasing fragment passes `newUser` /
        // `previousUser`; the docs assume earlier snippets created
        // them.
        private static User newUser = null;
        private static User previousUser = null;
        // The logging fragments pass an ILoggerFactory obtained from
        // ASP.NET Core dependency injection; the docs assume it
        // already exists. `dynamic` keeps the stub independent of the
        // Microsoft.Extensions.Logging.Abstractions type while still
        // letting `Logs.CoreLogging(loggerFactory)` compile.
        private static dynamic loggerFactory = null;
#pragma warning restore CS0414, CS0649

        public static void Main(string[] args)
        {
            System.Console.WriteLine("feature flag evaluates to true");
        }

        private void Wrappee()
        {
            try {
// BODY_BEGIN
{{ body }}
// BODY_END
            } catch (System.Exception) { /* never reached */ }
        }
    }
}
```
