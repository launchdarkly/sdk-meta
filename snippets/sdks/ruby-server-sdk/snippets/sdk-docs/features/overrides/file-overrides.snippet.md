---
id: ruby-server-sdk/sdk-docs/features/overrides/file-overrides
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: File-based flag override source configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
require 'ldclient-rb'

overrides = LaunchDarkly::Integrations::FileData.override_source(
  paths: ["/etc/launchdarkly/overrides.json"]
)

config = LaunchDarkly::Config.new(
  data_system_config: LaunchDarkly::DataSystem.default.overrides(overrides).build
)

client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config)
```
