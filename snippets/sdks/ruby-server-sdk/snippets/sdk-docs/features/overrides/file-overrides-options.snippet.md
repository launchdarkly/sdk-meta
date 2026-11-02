---
id: ruby-server-sdk/sdk-docs/features/overrides/file-overrides-options
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: File-based flag override source with multiple files, duplicate key handling, and change detection for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
overrides = LaunchDarkly::Integrations::FileData.override_source(
  paths: ["/etc/launchdarkly/overrides.json", "/etc/launchdarkly/team-overrides.json"],
  duplicate_keys_handling: :ignore,
  change_detection: :watching
)

config = LaunchDarkly::Config.new(
  data_system_config: LaunchDarkly::DataSystem.default.overrides(overrides).build
)
```
