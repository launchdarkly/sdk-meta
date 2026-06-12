---
id: ruby-server-sdk/sdk-docs/features/bigsegments/big-segments
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Big segments Redis store configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
store = LaunchDarkly::Integrations::Redis.new_big_segment_store(
  redis_url: 'redis://your-redis:6379',
  prefix: 'example-client-side-id'
)

config = LaunchDarkly::Config.new(
  big_segments: LaunchDarkly::BigSegmentsConfig.new(store: store)
)

client = LaunchDarkly::LDClient.new(sdk_key, config)
```
