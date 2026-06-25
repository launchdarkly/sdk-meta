---
id: go-server-sdk/sdk-docs/features/contextconfig/scoped-client-update
sdk: go-server-sdk
kind: reference
lang: go
description: Updating contexts in a scoped client for Go SDK v7.13.4+.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
userContext := ldcontext.New("example-user-key")

scopedClient := ld.NewScopedClient(client, userContext)
scopedClient.CurrentContext() // returns the single "user" context

scopedClient.AddContext(ldcontext.NewWithKind("device", "example-device-key"))
scopedClient.CurrentContext() // returns a multi-context with "user" and "device" contexts

currentUserContext := scopedClient.CurrentContext().IndividualContextByKind("user")
updatedUserContext := ldcontext.NewBuilderFromContext(currentUserContext).
  Set("additionalAttribute", "attribute value").
  Build()
scopedClient.OverWriteContextByKind(updatedUserContext)

scopedClient.BoolVariation("example-flag-key", false) // evaluates the flag using a multi-context with updated "user" context and "device" context
```
