---
id: ruby-server-sdk/sdk-docs/features/privateattrs/context-ai
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Marking context attributes private in the context object for the Ruby AI SDK.
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
    groups: ["Acme", "Global Health Services"],
    _meta: {
      privateAttributes: ['email']
    }
})
```
