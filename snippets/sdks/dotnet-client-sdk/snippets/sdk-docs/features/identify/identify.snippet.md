---
id: dotnet-client-sdk/sdk-docs/features/identify/identify
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Identify example for the .NET client SDK v3.0+.
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only
---

```csharp
var updatedContext = Context.Builder("example-context-key")
    .Set("email", "sandy@example.com")
    .Build();

// Synchronous method
client.Identify(updatedContext, TimeSpan.FromSeconds(5));

// Asynchronous method
await client.IdentifyAsync(updatedContext);
```
