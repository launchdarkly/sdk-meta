---
id: dotnet-server-sdk/sdk-docs/features/testdata/configure
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Test data source configuration for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
using LaunchDarkly.Sdk.Server.Integrations;

var td = TestData.DataSource();
// You can set any initial flag states here with td.Update

var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataSource(td)
    .Build();
var client = new LdClient(config);
```
