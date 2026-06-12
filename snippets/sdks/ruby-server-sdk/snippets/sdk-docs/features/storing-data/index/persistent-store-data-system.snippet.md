---
id: ruby-server-sdk/sdk-docs/features/storing-data/index/persistent-store-data-system
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Persistent store configuration via the DataSystem builder for Ruby SDK 8.12.0+.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
require 'ldclient-rb'

store = SomeKindOfFeatureStore.new(storeOptions)

config = LaunchDarkly::Config.new(
  data_system_config: LaunchDarkly::DataSystem.persistent_store(store).build
)

client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config)
```
