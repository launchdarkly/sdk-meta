---
id: ruby-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode-datasystem
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Daemon mode DataSystem configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
require 'ldclient-rb'

store = SomeKindOfFeatureStore.new(storeOptions)

config = LaunchDarkly::Config.new(
  data_system_config: LaunchDarkly::DataSystem.daemon(store).build
)

client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config)
```
