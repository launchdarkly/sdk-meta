---
id: dotnet-server-sdk/sdk-docs/features/flagchanges/flag-changes
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Flag change subscription example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only-members

---

```csharp
void LogWheneverAnyFlagChanges(LdClient client) {
    client.FlagTracker.FlagChanged += (sender, eventArgs) =>
    {
        Console.WriteLine("Flag \"{0}\" has changed", eventArgs.Key);
    };
}
```
