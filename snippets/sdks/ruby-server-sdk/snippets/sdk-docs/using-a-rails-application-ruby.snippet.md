---
id: ruby-server-sdk/sdk-docs/using-a-rails-application-ruby
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby in section \"Using a Rails application\""
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
Rails.configuration.client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY")
```
