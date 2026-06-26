---
id: ruby-server-sdk/sdk-docs/features/anonymous/anonymous-ai
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Anonymous context example for the Ruby AI SDK.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
context = LaunchDarkly::LDContext.create({ key: "example-context-key", anonymous: true })
```
