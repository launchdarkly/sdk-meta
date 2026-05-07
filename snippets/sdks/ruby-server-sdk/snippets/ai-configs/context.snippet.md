---
id: ruby-server-sdk/ai-configs/context
sdk: ruby-server-sdk
kind: context
lang: ruby
file: ruby-server-sdk/ai-configs/context.txt
description: Build an evaluation context for ruby-server-sdk AI Configs.
---

```ruby
context = LaunchDarkly::LDContext.create({
  key: 'context-key-123abc',
  kind: 'user',
  firstName: 'Sandy',
  lastName: 'Smith',
  email: 'sandy@example.com',
  groups: ['Google', 'Microsoft']
})
```
