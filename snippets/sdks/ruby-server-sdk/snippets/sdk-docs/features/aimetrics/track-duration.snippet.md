---
id: ruby-server-sdk/sdk-docs/features/aimetrics/track-duration
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Track duration manually for the Ruby AI SDK.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
# Track your own start and stop time.

# Set duration to the time (in ms) that your AI model generation takes.
# The duration may include network latency, depending on how you calculate it.

ai_config.tracker.track_duration(duration)
```
