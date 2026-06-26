---
id: dotnet-server-sdk/sdk-docs/features/storing-data/dynamodb/dynamodb
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: DynamoDB feature store configuration example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.Sdk.Server.Integrations;
var config = Configuration.Builder(sdkKey)
    .DataStore(
        Components.PersistentDataStore(
            DynamoDB.DataStore("my-table")
        ).CacheSeconds(30)
    )
    .Build();
var client = new LdClient(config);
```
