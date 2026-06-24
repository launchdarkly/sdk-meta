---
id: ruby-server-sdk/sdk-docs/features/offlinemode/offline-mode
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Offline mode example for Ruby SDK v7.0.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
config = LaunchDarkly::Config.new({offline: true})
client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config)
client.variation("any.feature.flag", context, false) # will always return the default value (false)
```
