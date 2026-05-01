---
id: dotnet-client-sdk/sdk-docs/migration-2-to-3-referencing-properties-of-an-attribute-object-3-0-syntax-context-with-object-attributes
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "3.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
---

```csharp
var address = LdValue.ObjectBuilder("address")
  .Add("street", "Main St")
  .Add("city", "Springfield")
  .Build();

var context = Context.Builder("example-context-key")
  .Set("address", address)
  .Build();
```
