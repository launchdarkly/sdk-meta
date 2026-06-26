---
id: python-server-sdk/sdk-docs/features/aimetrics/get-summary-agent
sdk: python-server-sdk
kind: reference
lang: python
description: Retrieve automatically recorded metrics with get_summary in agent mode for the Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
agent.tracker.get_summary()

# recorded metrics available in agent.tracker.get_summary().duration, .feedback,
# .success, .usage, and .time_to_first_token

```
