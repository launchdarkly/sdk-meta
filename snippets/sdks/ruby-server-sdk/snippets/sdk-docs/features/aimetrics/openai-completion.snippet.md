---
id: ruby-server-sdk/sdk-docs/features/aimetrics/openai-completion
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Record metrics from an OpenAI operation with track_openai_metrics for the Ruby AI SDK.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
if ai_config.enabled
  # Wrap the OpenAI operation to record metrics.
  # When calling the OpenAI operation, use details from ai_config.
  # For instance, you can pass ai_config.model.name
  # and ai_config.messages to your specific OpenAI operation.
  #
  # CAUTION: If the call inside track_openai_metrics throws an exception,
  # the SDK will re-throw that exception.

  completion = ai_config.tracker.track_openai_metrics do
    openai_client.chat.completions.create(
      model: ai_config.model.name,
      messages: ai_config.messages.map(&:to_h)
    )
  end
else
  # Application path to take when the ai_config is disabled
end
```
