---
id: ruby-server-sdk/sdk-docs/features/datasaving/standard-setup
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Data saving mode standard setup for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
require 'ldclient-rb'

config = LaunchDarkly::Config.new(
  data_system_config: LaunchDarkly::DataSystem.default.build
)

client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config)
```
