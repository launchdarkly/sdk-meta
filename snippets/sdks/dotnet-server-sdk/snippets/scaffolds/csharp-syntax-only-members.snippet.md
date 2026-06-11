---
id: dotnet-server-sdk/scaffolds/csharp-syntax-only-members
sdk: dotnet-server-sdk
kind: scaffold
lang: csharp
file: Program.cs
description: |
  Class-member-scope sibling of `csharp-syntax-only`. That scaffold
  splices the body inside `Wrappee()`, which breaks for fragments
  that are themselves method declarations (C# local functions can't
  carry the doc fragments' plain method shape with comments between
  members). This variant splices the body at class scope instead.

  Same `dotnet-server` validator container, so the body compiles
  against the real `LaunchDarkly.ServerSdk` package.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced at class scope.
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

namespace LaunchDarklySnippet
{
    public class Program
    {
        public static void Main(string[] args)
        {
            System.Console.WriteLine("feature flag evaluates to true");
        }

{{ body }}
    }
}
```
