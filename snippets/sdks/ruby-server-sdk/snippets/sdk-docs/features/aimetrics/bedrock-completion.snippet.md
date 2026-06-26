---
id: ruby-server-sdk/sdk-docs/features/aimetrics/bedrock-completion
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Record metrics from a Bedrock Converse command with track_bedrock_converse_metrics for the Ruby AI SDK.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
if ai_config.enabled
  # Wrap the Bedrock Converse command to record metrics.
  # When calling the Bedrock Converse command, use details from ai_config.
  # For instance, you can pass ai_config.model.name
  # and ai_config.messages to your specific Bedrock Converse command.

  completion = ai_config.tracker.track_bedrock_converse_metrics do
    bedrock_client.converse(
      map_converse_arguments(
        ai_config.model.name,
        ai_config.messages
      )
    )
  end
else
  # Application path to take when the ai_config is disabled
end
```
