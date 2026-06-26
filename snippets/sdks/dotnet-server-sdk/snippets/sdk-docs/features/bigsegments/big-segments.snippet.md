---
id: dotnet-server-sdk/sdk-docs/features/bigsegments/big-segments
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Big segments Redis store configuration example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.Sdk.Server.Integrations;
var config = Configuration.Builder("YOUR_SDK_KEY")
    .BigSegments(
        Components.BigSegments(
            Redis.BigSegmentStore()
                .HostAndPort("your-redis", 6379)
                .Prefix("example-client-side-id")
        )
        .ContextCacheSize(2000)
        .ContextCacheTime(TimeSpan.FromSeconds(30))
    )
    .Build();
var client = new LdClient(config);
```
