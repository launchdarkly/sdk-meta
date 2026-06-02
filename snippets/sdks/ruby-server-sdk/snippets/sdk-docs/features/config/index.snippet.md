---
id: ruby-server-sdk/sdk-docs/features/config/index
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: SDK configuration example for Ruby.
---

```ruby
config = LaunchDarkly::Config.new({connect_timeout: 1, read_timeout: 2})
client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config)
```
