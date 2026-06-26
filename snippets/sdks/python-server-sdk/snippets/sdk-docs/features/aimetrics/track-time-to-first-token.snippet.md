---
id: python-server-sdk/sdk-docs/features/aimetrics/track-time-to-first-token
sdk: python-server-sdk
kind: reference
lang: python
description: Track time to first token for the Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
# Track the time it takes to generate the first token.

# Pass in the time (in ms) until your first token is generated.
# This may include network latency, depending on how you calculate it.

tracker.track_time_to_first_token(1000)
```
