---
id: ruby-server-sdk/sdk-docs/features/contextconfig/multi-context
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Multi-context example for Ruby SDK v7.0.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
multi_context = LaunchDarkly::LDContext.create_multi([
    LaunchDarkly::LDContext.with_key("example-user-key"),
    LaunchDarkly::LDContext.with_key("example-device-key", "device"),
])
```
