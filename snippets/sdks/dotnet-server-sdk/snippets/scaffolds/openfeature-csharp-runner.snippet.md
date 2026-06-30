---
id: dotnet-server-sdk/scaffolds/openfeature-csharp-runner
sdk: dotnet-server-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Runs an OpenFeature provider doc fragment that assumes a registered
  provider and a bound `provider`, `client`, and `context` already
  exist — the "evaluate a context" and "access the LaunchDarkly client"
  fragments. The scaffold registers a real LaunchDarkly provider inside
  an async `Main`, binds those names, runs the fragment, then evaluates
  a flag and prints the success line. C# forbids re-declaring a local
  the scaffold already bound, so fragments that declare their own
  `context` use the `openfeature-csharp-context-runner` variant.
inputs:
  body:
    type: string
    description: The wrappee fragment, run with `provider`, `client`, and `context` in scope.
validation:
  runtime: dotnet-server
  entrypoint: Program.cs
---

```csharp
using System;
using System.Threading.Tasks;
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.OpenFeature.ServerProvider;
using OpenFeature.Model;

namespace LaunchDarklySnippet
{
    public class Program
    {
        public static async Task Main(string[] args)
        {
            var config = Configuration.Builder(Environment.GetEnvironmentVariable("LAUNCHDARKLY_SDK_KEY"))
                .StartWaitTime(TimeSpan.FromSeconds(10))
                .Build();
            var provider = new Provider(config);
            await OpenFeature.Api.Instance.SetProviderAsync(provider);
            var client = OpenFeature.Api.Instance.GetClient();
            var context = EvaluationContext.Builder().Set("targetingKey", "example-user-key").Build();

{{ body }}

            await client.GetBooleanValueAsync(
                Environment.GetEnvironmentVariable("LAUNCHDARKLY_FLAG_KEY"), false, context);
            Console.WriteLine("feature flag evaluates to true");
        }
    }
}
```
