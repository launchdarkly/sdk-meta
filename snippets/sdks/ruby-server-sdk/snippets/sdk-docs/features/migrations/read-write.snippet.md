---
id: ruby-server-sdk/sdk-docs/features/migrations/read-write
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Migration read and write example for Ruby SDK v8.0.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
context = LaunchDarkly::LDContext.create({key: "example-user-key", kind:"user"})

# this is the migration stage to use if the flag's migration stage
# is not available from LaunchDarkly
default_stage = LaunchDarkly::Migrations::STAGE_OFF

read_result = migrator.read("example-migration-flag-key", context, default_stage, payload)

write_result = migrator.write("example-migration-flag-key", context, default_stage, payload)
```
