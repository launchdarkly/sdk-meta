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

namespace LaunchDarklySnippet
{
    public class Program
    {
        // Stub so the wrappee body's `client.BoolVariation(...)`,
        // `client.StringVariation(...)`, etc. resolve at compile time.
        // Never invoked at runtime — Main below short-circuits to print
        // the EXAM-HELLO line.
        #pragma warning disable CS8625
        private static LdClient client = null;
        #pragma warning restore CS8625

        public static void Main(string[] args)
        {
            System.Console.WriteLine("feature flag evaluates to true");
        }

        #pragma warning disable CS0162
        private void Wrappee()
        {
{{ body }}
        }
        #pragma warning restore CS0162
    }
}
```
