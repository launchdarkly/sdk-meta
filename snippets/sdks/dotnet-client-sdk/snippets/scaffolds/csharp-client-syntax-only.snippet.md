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
---

```csharp
namespace LaunchDarklySnippet
{
    public class Program
    {
        public static void Main(string[] args)
        {
            System.Console.WriteLine("feature flag evaluates to true");
        }

        private void Wrappee()
        {
{{ body }}
        }
    }
}
```
