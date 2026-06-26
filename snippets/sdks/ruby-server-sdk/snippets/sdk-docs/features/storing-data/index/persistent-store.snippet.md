---
id: ruby-server-sdk/sdk-docs/features/storing-data/index/persistent-store
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Persistent feature store configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
require 'ldclient-rb'

store = SomeKindOfFeatureStore.new(storeOptions)
config = LaunchDarkly::Config.new(
  feature_store: store
)
client = LaunchDarkly::LDClient.new(sdk_key, config)
```
