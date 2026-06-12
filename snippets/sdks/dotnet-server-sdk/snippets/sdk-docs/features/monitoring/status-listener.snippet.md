---
id: dotnet-server-sdk/sdk-docs/features/monitoring/status-listener
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Data source status change handler for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only-typed

---

```csharp
client.DataSourceStatusProvider.StatusChanged +=
    (sender, status) => {
        Console.WriteLine("new status is: {0}", status);
    };
```
