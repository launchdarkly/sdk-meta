---
id: ruby-server-sdk/sdk-docs/features/hooks/add-hook
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Adding a hook to an existing client for the Ruby SDK.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config)

client.add_hook(example_hook)

```
