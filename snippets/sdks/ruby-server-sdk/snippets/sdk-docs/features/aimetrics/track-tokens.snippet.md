---
id: ruby-server-sdk/sdk-docs/features/aimetrics/track-tokens
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Track token usage manually for the Ruby AI SDK.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
# Track your own token usage.

# First, create an instance of TokenUsage.
# Update the input, output, and total values
# with return values from your AI model generation.
tokens = LaunchDarkly::Server::AI::TokenUsage.new(total: 300, input: 200, output: 100)

ai_config.tracker.track_tokens(tokens)
```
