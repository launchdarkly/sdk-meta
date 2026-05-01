---
id: go-server-sdk/sdk-docs/migration-4-to-5-using-the-user-builder-4-16-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "4.16+ syntax in section \"Using the user builder\""
---

```go
import (
  "gopkg.in/launchdarkly/go-server-sdk.v1/ldvalue"
  ld "gopkg.in/launchdarkly/go-server-sdk.v4"
)

// 4.16+: create a simple user with just a key
user := ld.NewUser("key")

// 4.16+: set email and country
user := ld.NewUserBuilder("key").
    Email("sandy@example.com").
    Country("us").
    Build()

// 4.16+: set custom attribute "team" to "admin"
user := ld.NewUserBuilder("key").
    Custom("team", ldvalue.String("admin")).
    Build()
```
