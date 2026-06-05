---
id: ruby-server-sdk/sdk-docs/features/config/app-config
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Application metadata configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
LaunchDarkly::Config.new({
  application: {
    id: "authentication-service",
    version: "abc123def456"
  }
})
```
