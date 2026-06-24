---
id: dotnet-server-sdk/sdk-docs/features/flagchanges/flag-value-changes-v7
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Flag value change subscription example for .NET (server-side) SDK v7.0.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only-members

---

```csharp
void LogWheneverOneFlagChangesForOneContext(LdClient client, string flagKey, Context context) {
    client.FlagTracker.FlagChanged += client.FlagTracker.FlagValueChangeHandler(
        flagKey,
        context,
        (sender, eventArgs) =>
        {
            Console.WriteLine(
                "Flag \"{0}\" for context \"{1}\" has changed from {2} to {3}",
                flagKey,
                context.Key,
                eventArgs.OldValue,
                eventArgs.NewValue
            );
        });
}
```
