---
id: dotnet-server-sdk/sdk-docs/features/testdata/flag-behavior-v7
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Configuring test data flag behavior for .NET (server-side) SDK v7.0.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
// This flag is true for the context with the key "example-context-key" and kind of "organization",
// and false for everyone else
td.Update(td.Flag("flag-key-456def")
    .VariationForKey(ContextKind.Of("organization"), "example-context-key", true)
    .FallthroughVariation(false));

// This flag returns the string variation "green" for contexts that have the
// attribute "admin" with a value of true, and "red" for everyone else.
td.Update(td.Flag("flag-key-789ghi")
    .Variations(LdValue.Of("red"), LdValue.Of("green"))
    .FallthroughVariation(0)
    .IfMatch("admin", LdValue.Of(true))
    .ThenReturn(1));
```
