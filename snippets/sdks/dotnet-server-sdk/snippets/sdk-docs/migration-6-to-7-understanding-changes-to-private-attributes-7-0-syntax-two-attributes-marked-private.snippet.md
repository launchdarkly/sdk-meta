---
id: dotnet-server-sdk/sdk-docs/migration-6-to-7-understanding-changes-to-private-attributes-7-0-syntax-two-attributes-marked-private
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "7.0 syntax, two attributes marked private in section \"Understanding changes to private attributes\""
---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .Events(
        Components.SendEvents()
          .PrivateAttributes("email", "address")
    )
    .Build();

var client = new LDClient(config);
```
