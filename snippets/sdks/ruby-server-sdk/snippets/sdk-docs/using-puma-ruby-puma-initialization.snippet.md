---
id: ruby-server-sdk/sdk-docs/using-puma-ruby-puma-initialization
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby Puma initialization in section \"Using Puma\""
---

```ruby
client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY")

on_worker_boot do
  client.postfork
end
```
