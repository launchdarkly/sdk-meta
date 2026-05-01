---
id: go-server-sdk/sdk-docs/migration-4-to-5-using-the-user-builder-5-0-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "5.0 syntax in section \"Using the user builder\""
---

```go
import (
  "gopkg.in/launchdarkly/go-sdk-common.v2/lduser"
  "gopkg.in/launchdarkly/go-sdk-common.v2/ldvalue"
)

// 5.0: create a simple user with just a key
user := lduser.NewUser("key")

// 5.0: set email and country
user := lduser.NewUserBuilder("key").
    Email("sandy@example.com").
    Country("us").
    Build()

// 5.0: set custom attribute "team" to "admin"
user := lduser.NewUserBuilder("key").
    Custom("team", ldvalue.String("admin")).
    Build()
```
