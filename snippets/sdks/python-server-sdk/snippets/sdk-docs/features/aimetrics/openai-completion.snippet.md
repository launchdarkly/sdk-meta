---
id: python-server-sdk/sdk-docs/features/aimetrics/openai-completion
sdk: python-server-sdk
kind: reference
lang: python
description: Record metrics from an OpenAI operation in completion mode for the Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
if config.enabled:
    # Pass in the result of the OpenAI operation.
    # When calling the OpenAI operation, use details from config.
    # For instance, you can pass config.model.name
    # and config.messages[0].content to your specific OpenAI operation.
    #
    # CAUTION: If the call inside of track_openai_metrics throws an exception,
    # the SDK will re-throw that exception

    messages = [] if config.messages is None else config.messages
    completion = tracker.track_openai_metrics(
        lambda:
          openai_client.chat.completions.create(
              model=config.model.name,
              messages=[message.to_dict() for message in messages],
          )
    )
else:
    # Application path to take when the config is disabled
    pass
```
