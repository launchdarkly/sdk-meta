---
id: ruby-server-sdk/sdk-docs/features/migrations/stage-switch
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Per-stage migration structure for Ruby SDK v8.0.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
# define the combination of reads and writes from the new and old systems
# that should occur at each migration stage

case stage
when LaunchDarkly::Migrations::STAGE_OFF
when LaunchDarkly::Migrations::STAGE_DUALWRITE
when LaunchDarkly::Migrations::STAGE_SHADOW
when LaunchDarkly::Migrations::STAGE_LIVE
when LaunchDarkly::Migrations::STAGE_RAMPDOWN
when LaunchDarkly::Migrations::STAGE_COMPLETE
else
  # throw an error
end
```
