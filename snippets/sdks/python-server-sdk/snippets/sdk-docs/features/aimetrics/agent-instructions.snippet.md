---
id: python-server-sdk/sdk-docs/features/aimetrics/agent-instructions
sdk: python-server-sdk
kind: reference
lang: python
description: Access instructions and record metrics in agent mode for the Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
if agent.enabled:
    # Retrieve instructions from the config and pass to your AI model
    result = example_model_api(agent.instructions)

    # Track metrics from the result
    agent.tracker.track_success()
else:
    # Application path to take when the agent is disabled
    pass
```
