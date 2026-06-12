---
id: ruby-server-sdk/sdk-docs/features/otel/tracing-hook
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: OpenTelemetry tracing hook configuration for the Ruby SDK.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
require 'ldclient-otel'

config = LaunchDarkly::Config.new(hooks: [LaunchDarkly::Otel::TracingHook.new])

client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config);

```
