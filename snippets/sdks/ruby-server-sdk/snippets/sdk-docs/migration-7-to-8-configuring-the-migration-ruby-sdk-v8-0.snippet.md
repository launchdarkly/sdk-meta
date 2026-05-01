---
id: ruby-server-sdk/sdk-docs/migration-7-to-8-configuring-the-migration-ruby-sdk-v8-0
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby SDK v8.0 in section \"Configuring the migration\""
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
