---
id: ruby-server-sdk/sdk-docs/using-passenger-ruby-passenger-initialization
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby Passenger initialization in section \"Using Passenger\""
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY")

if defined?(PhusionPassenger)
  PhusionPassenger.on_event(:starting_worker_process) do |forked|
    client.postfork
  end
end
```
