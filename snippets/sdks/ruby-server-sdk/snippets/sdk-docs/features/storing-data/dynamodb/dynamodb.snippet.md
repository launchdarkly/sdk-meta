---
id: ruby-server-sdk/sdk-docs/features/storing-data/dynamodb/dynamodb
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: DynamoDB feature store configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
require 'ldclient-rb'

store = LaunchDarkly::Integrations::DynamoDB.new_feature_store('my-table',
  { expiration: 30 })

config = LaunchDarkly::Config.new(
  feature_store: store
)
client = LaunchDarkly::LDClient.new(sdk_key, config)
```
