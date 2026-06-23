---
id: dotnet-client-sdk/sdk-docs/features/privateattrs/config-v3-0
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Private attribute configuration for .NET client SDK v3.0.
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only-v3

---

```csharp
// All attributes marked private
var configAllPrivate = Configuration.Builder("example-mobile-key")
  .Events(Components.SendEvents().AllAttributesPrivate(true))
  .Build();
LdClient client = LdClient.Init(configAllPrivate, context, TimeSpan.FromSeconds(10));

// Two attributes marked private
var configSomePrivate = Configuration.Builder("example-mobile-key")
  .Events(Components.SendEvents().PrivateAttributes("email", "address"))
  .Build();
client = LdClient.Init(configSomePrivate, context, TimeSpan.FromSeconds(10));
```
