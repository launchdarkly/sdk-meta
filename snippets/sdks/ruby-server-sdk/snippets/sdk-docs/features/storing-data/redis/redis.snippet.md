---
id: ruby-server-sdk/sdk-docs/features/storing-data/redis/redis
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Redis feature store configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
require 'ldclient-rb'

store = LaunchDarkly::Integrations::Redis.new_feature_store(
  redis_url: 'redis://my-redis:6379',
  prefix: 'my-key-prefix',
  expiration: 30
)

config = LaunchDarkly::Config.new(
  feature_store: store
)
client = LaunchDarkly::LDClient.new(sdk_key, config)
```
