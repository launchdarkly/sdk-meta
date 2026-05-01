---
id: dotnet-server-sdk/sdk-docs/migration-6-to-7-understanding-changes-to-private-attributes-7-0-syntax-attribute-marked-private-for-one-context
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: "7.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```csharp

var context = Context.Builder("example-context-key")
    .Name("Sandy")
    .Set("email", "sandy@example.com")
    .Private("email")
    .Build();

```
