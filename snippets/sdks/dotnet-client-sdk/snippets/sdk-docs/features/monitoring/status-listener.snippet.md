---
id: dotnet-client-sdk/sdk-docs/features/monitoring/status-listener
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Data source status change handler for .NET (client-side).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only-typed

---

```csharp
client.DataSourceStatusProvider.StatusChanged +=
    (sender, status) => {
        Console.WriteLine("new status is: {0}", status);
    };
```
