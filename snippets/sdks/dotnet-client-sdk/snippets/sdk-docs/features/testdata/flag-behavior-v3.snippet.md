---
id: dotnet-client-sdk/sdk-docs/features/testdata/flag-behavior-v3
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Configuring test data flag behavior for .NET (client-side) SDK v3.0.
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only-v3
---

```csharp
// This flag is true for the context key "example-context-key" and false for everyone else
td.Update(td.Flag("flag-key-456def")
    .Variation(false)
    .VariationForKey(ContextKind.Of("user"), "example-context-key", true));

// This flag returns the string variation "green" for contexts where the custom
// attribute "admin" has a value of true, and "red" for all other contexts.
td.Update(td.Flag("flag-key-789ghi")
    .Variations(LdValue.Of("red"), LdValue.Of("green"))
    .VariationFunc(context =>
        context.GetValue("admin").AsBool ? 1 : 0));
```
