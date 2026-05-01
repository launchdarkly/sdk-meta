---
id: ruby-server-sdk/sdk-docs/migration-6-to-7-understanding-changes-to-private-attributes-7-0-syntax-attribute-marked-private-for-one-context
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "7.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```ruby
context = LaunchDarkly::LDContext.create({key: "key", name: "Sandy", email: "sandy@example.com", _meta: {privateAttributes: ["email"]}})
```
