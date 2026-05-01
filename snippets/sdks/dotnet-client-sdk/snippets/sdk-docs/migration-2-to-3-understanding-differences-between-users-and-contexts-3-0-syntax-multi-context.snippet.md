---
id: dotnet-client-sdk/sdk-docs/migration-2-to-3-understanding-differences-between-users-and-contexts-3-0-syntax-multi-context
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "3.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
---

```csharp
var userContext = Context.New("example-user-key");

var orgContext = Context.New(ContextKind.Of("organization"), "example-organization-key");

var multiContext = Context.NewMulti(userContext, orgContext);
```
