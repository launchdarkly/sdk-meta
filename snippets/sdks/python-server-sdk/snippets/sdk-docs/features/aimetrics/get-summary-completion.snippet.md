---
id: python-server-sdk/sdk-docs/features/aimetrics/get-summary-completion
sdk: python-server-sdk
kind: reference
lang: python
description: Retrieve automatically recorded metrics with get_summary in completion mode for the Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
tracker.get_summary()

# recorded metrics available in tracker.get_summary().duration, .feedback,
# .success, .usage, and .time_to_first_token
```
