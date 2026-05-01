---
id: dotnet-server-sdk/scaffolds/csharp-syntax-only
sdk: dotnet-server-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Parse-only validator for .NET server SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: dotnet-server
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
