---
id: go-server-sdk/sdk-docs/migration-5-to-6-working-with-built-in-and-custom-attributes-5-x-syntax-user-with-attributes
sdk: go-server-sdk
kind: reference
lang: go
description: "5.x syntax, user with attributes in section \"Working with built-in and custom attributes\""
---

```go
user2 := lduser.NewUserBuilder("example-user-key").
    Name("Sandy").
    Email("sandy@example.com").
    Custom("groups", ldvalue.ArrayOf(
        ldvalue.String("Acme"), ldvalue.String("Global Health Services"))).
    Build()
```
