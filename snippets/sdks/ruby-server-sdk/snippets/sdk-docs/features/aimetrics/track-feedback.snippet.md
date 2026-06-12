---
id: ruby-server-sdk/sdk-docs/features/aimetrics/track-feedback
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Track output satisfaction rate for the Ruby AI SDK.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
# Track your own output satisfaction rate

# Pass in kind: :positive or kind: :negative
ai_config.tracker.track_feedback(kind: :positive)
```
