---
id: dotnet-server-sdk/sdk-docs/migration-6-to-7-referencing-properties-of-an-attribute-object-7-0-syntax-context-with-object-attributes
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "7.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
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
