---
id: ruby-server-sdk/sdk-docs/migration-6-to-7-understanding-differences-between-users-and-contexts-7-0-syntax-multi-context
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "7.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
---

```ruby
multi_context = LaunchDarkly::LDContext.create_multi([
    LaunchDarkly::LDContext.with_key("example-user-key"),
    LaunchDarkly::LDContext.with_key("example-device-key", "device"),
])
```
