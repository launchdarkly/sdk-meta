---
id: ruby-server-sdk/sdk-docs/features/otel/tracing-hook-options
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: OpenTelemetry tracing hook with span and value options for the Ruby SDK.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
require 'ldclient-otel'

tracing_hook_options = LaunchDarkly::Otel::TracingHookOptions.new(add_spans: true, include_value: true)
hook = LaunchDarkly::Otel::TracingHook.new(tracing_hook_options)
config = LaunchDarkly::Config.new(hooks: [hook])

client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config);

```
