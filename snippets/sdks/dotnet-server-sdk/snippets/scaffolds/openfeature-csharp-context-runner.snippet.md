---
id: dotnet-server-sdk/scaffolds/openfeature-csharp-context-runner
sdk: dotnet-server-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Runs an OpenFeature "construct a context" fragment, which declares its
  own `context` local. The scaffold registers a real LaunchDarkly
  provider and binds `client` inside an async `Main`, but leaves
  `context` for the fragment to declare; afterward it evaluates a flag
  with the fragment's `context` and prints the success line. Separate
  from `openfeature-csharp-runner` because C# forbids re-declaring a
  local the scaffold already bound.
inputs:
  body:
    type: string
    description: The wrappee fragment; declares `context` and runs with `client` in scope.
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

{{ body }}

            await client.GetBooleanValueAsync(
                Environment.GetEnvironmentVariable("LAUNCHDARKLY_FLAG_KEY"), false, context);
            Console.WriteLine("feature flag evaluates to true");
        }
    }
}
```
