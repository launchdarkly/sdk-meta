---
id: go-server-sdk/sdk-docs/migration-4-to-5-using-the-value-type-for-json-values-before-4-16
sdk: go-server-sdk
kind: reference
lang: go
description: "before 4.16 in section \"Using the Value type for JSON values\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v4"
)

// before 4.16: get a JSON flag variation whose default value is "default",
// and check if the result is a string or a number
result, _ := client.JsonVariation(flagKey, user, "default")
if s, ok := result.(string); ok {
    DoSomethingWithString(s)
} else if n, ok := result.(float64); ok {
    DoSomethingWithNumberAsInt(int(n))
}

// before 4.16: set a user's custom attribute "teams" to an
// array of ["admin", "foosball"]
user := ld.NewUser("key")
custom := make(map[string]interface{})
custom["team"] = []interface{}{"admin", "foosball"}
```
