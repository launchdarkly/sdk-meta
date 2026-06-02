---
id: ruby-server-sdk/sdk-docs/features/config/app-config
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Application metadata configuration example for Ruby.
---

```ruby
LaunchDarkly::Config.new({
  application: {
    id: "authentication-service",
    version: "abc123def456"
  }
})
```
