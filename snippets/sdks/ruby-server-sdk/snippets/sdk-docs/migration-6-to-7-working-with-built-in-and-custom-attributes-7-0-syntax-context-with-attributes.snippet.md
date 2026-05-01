---
id: ruby-server-sdk/sdk-docs/migration-6-to-7-working-with-built-in-and-custom-attributes-7-0-syntax-context-with-attributes
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "7.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```ruby
context = LaunchDarkly::LDContext.create({
    key: "example-user-key",
    kind: "user",
    firstName: "Sandy",
    lastName: "Smith",
    email: "sandy@example.com",
    groups: ["Acme", "Global Health Services"]
})
```
