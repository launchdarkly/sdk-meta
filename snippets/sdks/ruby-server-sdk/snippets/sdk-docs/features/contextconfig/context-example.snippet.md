---
id: ruby-server-sdk/sdk-docs/features/contextconfig/context-example
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Context example for Ruby SDK v7.0.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

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
