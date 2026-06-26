---
id: ruby-server-sdk/sdk-docs/features/aimetrics/track-time-to-first-token
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Track time to first token for the Ruby AI SDK.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
# Track the time it takes to generate the first token

# Pass in the time (in ms) until your first token is generated
# This may include network latency, depending on how you calculate it

ai_config.tracker.track_time_to_first_token(1000)
```
