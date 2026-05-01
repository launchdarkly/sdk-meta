---
id: dotnet-server-sdk/sdk-docs/migration-6-to-7-understanding-changes-to-private-attributes-7-0-syntax-all-attributes-marked-private
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "7.0 syntax, all attributes marked private in section \"Understanding changes to private attributes\""
---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .Events(
        Components.SendEvents()
           .AllAttributesPrivate(true)  // defaults to false
    )
    .Build();

var client = new LDClient(config);
```
