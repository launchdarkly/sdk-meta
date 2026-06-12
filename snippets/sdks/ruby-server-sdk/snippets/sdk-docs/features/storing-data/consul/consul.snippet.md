---
id: ruby-server-sdk/sdk-docs/features/storing-data/consul/consul
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Consul feature store configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
require 'ldclient-rb'

store = LaunchDarkly::Integrations::Consul.new_feature_store(
  { url: 'http://my-consul:8100', prefix: 'my-key-prefix', expiration: 30 })

config = LaunchDarkly::Config.new(
  feature_store: store
)
client = LaunchDarkly::LDClient.new(sdk_key, config)
```
