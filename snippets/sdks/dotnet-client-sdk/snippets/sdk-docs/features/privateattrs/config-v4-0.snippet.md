---
id: dotnet-client-sdk/sdk-docs/features/privateattrs/config-v4-0
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Private attribute configuration for .NET client SDK v4.0.
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only-v4

---

```csharp
// All attributes marked private
var configAllPrivate = Configuration
  .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
  .Events(Components.SendEvents().AllAttributesPrivate(true))
  .Build();
LdClient client = LdClient.Init(configAllPrivate, context, TimeSpan.FromSeconds(10));

// Two attributes marked private
var configSomePrivate = Configuration
  .Builder("example-mobile-key", ConfigurationBuilder.AutoEnvAttributes.Enabled)
  .Events(Components.SendEvents().PrivateAttributes("email", "address"))
  .Build();
client = LdClient.Init(configSomePrivate, context, TimeSpan.FromSeconds(10));
```
