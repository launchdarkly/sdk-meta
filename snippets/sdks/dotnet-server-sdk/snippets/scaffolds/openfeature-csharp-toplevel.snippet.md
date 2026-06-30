---
id: dotnet-server-sdk/scaffolds/openfeature-csharp-toplevel
sdk: dotnet-server-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Resolves an OpenFeature provider doc fragment that is a set of `using`
  directives. The harness lifts the body's `using …;` lines to file
  scope at the USING_LIFT_MARKER, so the C# compiler resolves each
  namespace against the OpenFeature package and the LaunchDarkly
  provider baked into the validator image. A `using` that names a
  namespace the packages don't ship fails compilation. The body does
  not connect to LaunchDarkly.
inputs:
  body:
    type: string
    description: The wrappee's using directives; lifted to file scope by the harness.
validation:
  runtime: dotnet-server
  entrypoint: Program.cs
---

```csharp
// USING_LIFT_MARKER
using System;

namespace LaunchDarklySnippet
{
    public class Program
    {
        public static void Main(string[] args)
        {
{{ body }}
            Console.WriteLine("feature flag evaluates to true");
        }
    }
}
```
