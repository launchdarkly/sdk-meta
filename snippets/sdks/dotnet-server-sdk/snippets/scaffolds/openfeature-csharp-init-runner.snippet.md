---
id: dotnet-server-sdk/scaffolds/openfeature-csharp-init-runner
sdk: dotnet-server-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Runs an OpenFeature "initialize the provider" fragment end-to-end
  against a real LaunchDarkly environment. The fragment builds a
  LaunchDarkly provider, registers it with the OpenFeature API, and
  binds a `client`; this scaffold runs it inside an async `Main`, then
  uses `client` to evaluate a flag and print the success line. The
  fragment's `YOUR_SDK_KEY` literal is replaced with the real key via
  `validation.placeholders`.
inputs:
  body:
    type: string
    description: The wrappee init fragment; registers the provider and binds `client`.
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
{{ body }}

            await client.GetBooleanValueAsync(
                Environment.GetEnvironmentVariable("LAUNCHDARKLY_FLAG_KEY"), false,
                EvaluationContext.Builder().Set("targetingKey", "example-user-key").Build());
            Console.WriteLine("feature flag evaluates to true");
        }
    }
}
```
