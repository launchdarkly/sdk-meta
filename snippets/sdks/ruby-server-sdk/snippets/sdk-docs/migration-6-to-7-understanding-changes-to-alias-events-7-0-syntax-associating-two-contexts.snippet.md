---
id: ruby-server-sdk/sdk-docs/migration-6-to-7-understanding-changes-to-alias-events-7-0-syntax-associating-two-contexts
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "7.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```ruby
context1 = LaunchDarkly::LDContext.with_key("example-user-key")
context2 = LaunchDarkly::LDContext.with_key("example-device-key", "device")
multi_context = LaunchDarkly::LDContext.create_multi([context1, context2])
client.identify(multi_context)
```
