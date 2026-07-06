---
id: dotnet-server-sdk/sdk-docs/openfeature/initialize-the-provider
sdk: dotnet-server-sdk
kind: reference
lang: csharp
file: dotnet-server-sdk/sdk-docs/openfeature/initialize-the-provider.cs
description: ".NET (server-side) OpenFeature provider in section \"Initialize the provider\""
validation:
  scaffold: dotnet-server-sdk/scaffolds/openfeature-csharp-init-runner
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
---

```csharp
var config = Configuration.Builder("YOUR_SDK_KEY")
    .StartWaitTime(TimeSpan.FromSeconds(10))
    .Build();

var provider = new Provider(config);

await OpenFeature.Api.Instance.SetProviderAsync(provider);

var client = OpenFeature.Api.Instance.GetClient();
```
