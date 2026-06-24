---
id: dotnet-server-sdk/sdk-docs/features/filedata/flags-from-files
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: File data source configuration example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
using LaunchDarkly.Sdk.Server;
using LaunchDarkly.Sdk.Server.Integrations;

var config = Configuration.Builder("sdk key")
    .DataSource(
        FileData.DataSource()
            .FilePaths("file1.json", "file2.json")
            .AutoUpdate(true)
    )
    .Events(Components.NoEvents)
    .Build();

var client = new LdClient(config);
```
