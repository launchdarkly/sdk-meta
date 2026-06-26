---
id: python-server-sdk/sdk-docs/features/aimetrics/bedrock-completion
sdk: python-server-sdk
kind: reference
lang: python
description: Record metrics from a Bedrock Converse command in completion mode for the Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
if config.enabled:
    # Pass in the result of the Bedrock Converse command.
    # When calling the Bedrock Converse command, use details from config.
    # For instance, you can pass config.model.name
    # and config.messages[0].content to your specific Bedrock Converse command.

    completion = tracker.track_bedrock_converse_metrics(
        client.converse(
            modelId=config.model.name,
            messages=map_messages_to_conversation(config.messages)
        )
    )
else:
    # Application path to take when the config is disabled
    pass
```
