---
id: ruby-server-sdk/sdk-docs/features/datasaving/relay-proxy-fallback
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Data saving mode with Relay Proxy and LaunchDarkly API fallback for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
require 'ldclient-rb'

relay_uri = "http://my-relay-proxy:8030"

config = LaunchDarkly::Config.new(
  data_system_config: LaunchDarkly::DataSystem.custom
    .initializers([
      LaunchDarkly::DataSystem.polling_ds_builder.base_uri(relay_uri),
      LaunchDarkly::DataSystem.polling_ds_builder,
    ])
    .synchronizers([
      LaunchDarkly::DataSystem.streaming_ds_builder.base_uri(relay_uri),
      LaunchDarkly::DataSystem.streaming_ds_builder,
      LaunchDarkly::DataSystem.polling_ds_builder,
    ])
    .build
)

client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config)
```
