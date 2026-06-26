---
id: ruby-server-sdk/sdk-docs/features/datasaving/file-bootstrap
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Data saving mode with file-based bootstrap and live updates for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
require 'ldclient-rb'

config = LaunchDarkly::Config.new(
  data_system_config: LaunchDarkly::DataSystem.custom
    .initializers([
      LaunchDarkly::Integrations::FileData.data_source_v2(paths: ['flags.json']),
      LaunchDarkly::DataSystem.polling_ds_builder,
    ])
    .synchronizers([
      LaunchDarkly::DataSystem.streaming_ds_builder,
      LaunchDarkly::DataSystem.polling_ds_builder,
    ])
    .build
)

client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config)
```
