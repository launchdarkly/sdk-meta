---
id: go-server-sdk/sdk-docs/use-go-contexts-with-ldscopedclient-add-ldscopedclient-to-go-context
sdk: go-server-sdk
kind: reference
lang: go
description: "Add LDScopedClient to Go context in section \"Use Go contexts with LDScopedClient\""
---

```go
  scopedClient := ld.NewScopedClient(client, ldContext)
  ctx := ld.GoContextWithScopedClient(context.Background(), scopedClient)
  otherFunction(ctx)
  // LDScopedClient is in beta and may change without notice.
```
