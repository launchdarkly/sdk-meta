---
id: dotnet-client-sdk/sdk-docs/features/flagchanges/flag-changes
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: Flag change subscription example for .NET (client-side).
validation:
  scaffold: dotnet-client-sdk/scaffolds/csharp-client-syntax-only-typed

---

```csharp
client.FlagTracker.FlagValueChanged += (sender, eventArgs) => {
    if (eventArgs.Key == "key-for-flag-i-am-watching") {
        DoSomethingWithNewFlagValue(eventArgs.NewValue.AsBool);
    }
};
```
