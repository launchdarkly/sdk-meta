---
id: dotnet-server-sdk/ai-configs/initialize
sdk: dotnet-server-sdk
kind: initialize
lang: csharp
file: dotnet-server-sdk/ai-configs/initialize.txt
description: Initialize the LaunchDarkly client and AI Configs client for dotnet-server-sdk.
---

```csharp
var baseClient = new LdClient(Configuration.Builder(
      "{{sdkkey}}"
    )
    .StartWaitTime(TimeSpan.FromSeconds(5)).Build());
var aiClient = new LdAiClient(new LdClientAdapter(baseClient));
```
