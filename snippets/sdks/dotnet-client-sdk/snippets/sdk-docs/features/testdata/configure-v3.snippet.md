---
id: dotnet-client-sdk/sdk-docs/features/testdata/configure-v3
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Test data source configuration for .NET (client-side) SDK v3.0.
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only-v3
---

```csharp
using LaunchDarkly.Sdk.Client.Integrations;

var td = TestData.DataSource();
// You can set any initial flag states here with td.Update

var config = Configuration.Builder("example-mobile-key")
    .DataSource(td)
    .Build();
var client = LdClient.Init(config, context, startWaitTime);
```
