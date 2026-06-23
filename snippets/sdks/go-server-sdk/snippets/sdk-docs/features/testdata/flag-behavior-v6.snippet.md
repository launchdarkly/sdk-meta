---
id: go-server-sdk/sdk-docs/features/testdata/flag-behavior-v6
sdk: go-server-sdk
kind: reference
lang: go
description: Configuring test data flag behavior for Go SDK v6.0.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
    "github.com/launchdarkly/go-sdk-common/v3/ldvalue"
)

// This flag is true for the context with the key "example-context-key" and kind of "organization",
// and false for everyone else
td.Update(td.Flag("flag-key-456def").
    VariationForKey("organization", "example-context-key", true).
    FallthroughVariation(false));


// This flag returns the string variation "green" for contexts that have the
// attribute "admin" with a value of true, and "red" for everyone else.
td.Update(td.Flag("flag-key-789ghi").
    Variations(ldvalue.String("red"), ldvalue.String("green")).
    FallthroughVariationIndex(0).
    IfMatch("admin", ldvalue.Bool(true)).
    ThenReturnIndex(1));
```
