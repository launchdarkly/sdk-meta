---
id: dotnet-server-sdk/sdk-docs/features/overrides/file-overrides-yaml
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: File-based flag override source reading a YAML file through YamlDotNet for .NET (server-side) SDK v9.0+.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only
---

```csharp
using YamlDotNet.Serialization;

var yaml = new DeserializerBuilder()
    .WithAttemptingUnquotedStringTypeDeserialization()
    .Build();

var overrides = FileOverrides.Source()
    .FilePaths("/etc/launchdarkly/overrides.yaml")
    .Parser(".yaml", content => yaml.Deserialize<object>(content));
```
