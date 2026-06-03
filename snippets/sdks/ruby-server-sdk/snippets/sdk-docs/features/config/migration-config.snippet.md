---
id: ruby-server-sdk/sdk-docs/features/config/migration-config
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Migration configuration example for the Ruby SDK v8 — read/write methods, execution order, latency/error tracking.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
builder = LaunchDarkly::Migrations::MigratorBuilder.new(@client)
builder.read(
  ->(_payload) { LaunchDarkly::Result.success('old value') },
  ->(_payload) { LaunchDarkly::Result.success('new value') },
  ->(lhs, rhs) { lhs == rhs }
)

builder.write(
  ->(_payload) { LaunchDarkly::Result.success('old value') },
  ->(_payload) { LaunchDarkly::Result.success('new value') }
)
builder.read_execution_order(builder.EXECUTION_PARALLEL) # or .EXECUTION_SERIAL, or .EXECUTION_RANDOM
builder.track_latency(true) # defaults to true
builder.track_errors(true)  # defaults to true
migrator = builder.build
```
