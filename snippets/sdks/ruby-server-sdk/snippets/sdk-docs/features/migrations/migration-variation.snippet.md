---
id: ruby-server-sdk/sdk-docs/features/migrations/migration-variation
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Migration stage evaluation (migration_variation) for Ruby SDK v8.0.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
context = LaunchDarkly::LDContext.create({key: "example-user-key", kind:"user"})

stage, tracker = client.migration_variation(
  "example-migration-flag-key",
  context,
  LaunchDarkly::Migrations::STAGE_OFF
)
```
