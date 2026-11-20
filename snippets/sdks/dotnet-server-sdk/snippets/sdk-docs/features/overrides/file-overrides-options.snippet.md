---
id: dotnet-server-sdk/sdk-docs/features/overrides/file-overrides-options
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: File-based flag override source with multiple files, duplicate key handling, and change detection for .NET (server-side) SDK v9.0+.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
var overrides = FileOverrides.Source()
    .FilePaths("/etc/launchdarkly/overrides.json", "/etc/launchdarkly/team-overrides.json")
    .DuplicateKeysHandling(FileOverrideTypes.DuplicateKeysHandling.Ignore)
    .ChangeDetection(FileOverrideTypes.ChangeDetection.Watching);

var config = Configuration.Builder("YOUR_SDK_KEY")
    .DataSystem(Components.DataSystem().Default().Overrides(overrides))
    .Build();
```
