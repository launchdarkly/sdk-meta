---
id: ruby-server-sdk/sdk-docs/migration-7-to-8-reading-and-writing-during-the-migration-ruby-sdk-v8-0
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby SDK v8.0 in section \"Reading and writing during the migration\""
---

```ruby
context = LaunchDarkly::LDContext.create({key: "example-user-key", kind:"user"})

# this is the migration stage to use if the flag's migration stage
# is not available from LaunchDarkly
default_stage = LaunchDarkly::Migrations::STAGE_OFF

read_result = migrator.read("example-migration-flag-key", context, default_stage, payload)

write_result = migrator.write("example-migration-flag-key", context, default_stage, payload)

```
