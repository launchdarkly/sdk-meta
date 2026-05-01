---
id: go-server-sdk/sdk-docs/migration-4-to-5-using-the-value-type-for-json-values-4-16-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "4.16+ syntax in section \"Using the Value type for JSON values\""
---

```go
import (
  "gopkg.in/launchdarkly/go-sdk-common.v1/ldvalue"
  ld "gopkg.in/launchdarkly/go-server-sdk.v4"
)

// 4.16+: get a JSON flag variation whose default value is "default",
// and check if the result is a string or a number
result, _ := client.JSONVariation(flagKey, user, ldvalue.String("default"))
if result.Type() == ldvalue.StringType {
    DoSomethingWithString(result.StringValue())
} else if result.Type() == ldvalue.NumberType {
    DoSomethingWIthNumberAsInt(result.IntValue())
}

// 4.16+: set a user's custom attribute "teams" to an
// array of ["admin", "foosball"]
user := ld.NewUserBuilder("key").
    Custom("teams", ldvalue.ArrayBuild().
        Add(ldvalue.String("admin")).Add(ldvalue.String("foosball")).Build()).
    Build()
```
