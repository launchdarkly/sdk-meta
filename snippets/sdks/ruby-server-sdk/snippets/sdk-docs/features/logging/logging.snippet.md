---
id: ruby-server-sdk/sdk-docs/features/logging/logging
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Custom logger configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
log = ::Logger.new($stdout)
log.level = ::Logger::DEBUG
config = LaunchDarkly::Config.new({logger: log})
client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config)
```
