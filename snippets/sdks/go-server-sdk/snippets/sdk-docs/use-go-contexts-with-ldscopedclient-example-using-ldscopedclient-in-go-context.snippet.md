---
id: go-server-sdk/sdk-docs/use-go-contexts-with-ldscopedclient-example-using-ldscopedclient-in-go-context
sdk: go-server-sdk
kind: reference
lang: go
description: "Example: Using LDScopedClient in Go context in section \"Use Go contexts with LDScopedClient\""
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
  func LDScopedClientMiddleware(client *LDClient) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
          scopedClient := NewScopedClient(client, ldcontext.New("example-context-key"))
          ctx := GoContextWithScopedClient(r.Context(), scopedClient)
          next.ServeHTTP(w, r.WithContext(ctx))
      })
    }
  }

  func requestLogic(r *http.Request) {
    featureFlagEnabled := MustGetScopedClient(r.Context()).BoolVariation("example-flag-key", false)
    // use featureFlagEnabled...
  }

```
