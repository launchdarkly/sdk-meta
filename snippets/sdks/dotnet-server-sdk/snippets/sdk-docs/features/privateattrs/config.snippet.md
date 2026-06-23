---
id: dotnet-server-sdk/sdk-docs/features/privateattrs/config
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Private attribute configuration for .NET server SDK v7.0.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
// All attributes marked as private
var config = Configuration.Builder("YOUR_SDK_KEY")
    .Events(
        Components.SendEvents()
           .AllAttributesPrivate(true)  // defaults to false
    )
    .Build();

var client = new LdClient(config);

// Two attributes marked as private
config = Configuration.Builder("YOUR_SDK_KEY")
    .Events(
        Components.SendEvents()
          .PrivateAttributes("email", "address")
    )
    .Build();

client = new LdClient(config);
```
