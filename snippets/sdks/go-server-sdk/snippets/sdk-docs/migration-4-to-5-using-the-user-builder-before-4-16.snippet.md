---
id: go-server-sdk/sdk-docs/migration-4-to-5-using-the-user-builder-before-4-16
sdk: go-server-sdk
kind: reference
lang: go
description: "before 4.16 in section \"Using the user builder\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v4"
)

// before 4.16: create a simple user with just a key
user := ld.NewUser("key")

// before 4.16: set email and country
user := ld.NewUser("key")
email := "sandy@example.com"
country := "us"
user.Email = &email
user.Country = &country

// before 4.16: set custom attribute "team" to "admin"
user := ld.NewUser("key")
custom := make(map[string]interface{})
custom["team"] = "admin"
```
