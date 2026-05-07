---
id: ruby-server-sdk/sdk-docs/using-passenger-ruby-passenger-initialization-with-rails
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby Passenger initialization with Rails in section \"Using Passenger\""
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
if defined?(PhusionPassenger)
  PhusionPassenger.on_event(:starting_worker_process) do |forked|
    Rails.configuration.client.postfork
  end
end
```
