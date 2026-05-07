---
id: ruby-server-sdk/sdk-docs/using-puma-ruby-puma-initialization
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby Puma initialization in section \"Using Puma\""
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY")

on_worker_boot do
  client.postfork
end
```
