---
id: dotnet-server-sdk/sdk-docs/install-the-sdk-net-sdk-v6-x-and-later-c
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: ".NET SDK v6.x and later (C#) in section \"Install the SDK\""
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
using LaunchDarkly.Sdk;
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.Observability;


// LaunchDarkly.Sdk defines general types like Context, which are also used in the client-side .NET SDK.
// LaunchDarkly.Sdk.Server defines the LdClient and Configuration types for the server-side .NET SDK.
// LaunchDarkly.Observability defines the optional observability,
// which requires .NET (server-side) SDK version 8.10 or later.
```
