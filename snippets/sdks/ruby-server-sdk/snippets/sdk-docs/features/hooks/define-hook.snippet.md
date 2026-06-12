---
id: ruby-server-sdk/sdk-docs/features/hooks/define-hook
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Hook implementation and configuration for the Ruby SDK.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only-toplevel
---

```ruby
require 'ldclient-rb'

class ExampleHook
  include LaunchDarkly::Interfaces::Hooks::Hook

  def metadata
    LaunchDarkly::Interfaces::Hooks::Metadata.new('example-hook')
  end

  # Implement at least one of `before_evaluation`, `after_evaluation`

  # `before_evaluation` is called during the execution of a variation method
  # before the flag value has been determined

  # `after_evaluation` is called during the execution of a variation method
  # after the flag value has been determined
end

example_hook = ExampleHook.new

config = LaunchDarkly::Config.new(hooks: [example_hook])

client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY", config)
```
