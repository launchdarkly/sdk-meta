---
id: ruby-server-sdk/sdk-docs/using-spring-ruby-spring-initialization
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby Spring initialization in section \"Using Spring\""
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY")

Spring.after_fork do
  client.postfork
end
```
