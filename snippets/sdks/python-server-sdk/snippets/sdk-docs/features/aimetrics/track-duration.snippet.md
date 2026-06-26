---
id: python-server-sdk/sdk-docs/features/aimetrics/track-duration
sdk: python-server-sdk
kind: reference
lang: python
description: Track duration manually for the Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
# Track your own start and stop time.

# Set duration to the time (in ms) that your AI model generation takes.
# The duration may include network latency, depending on how you calculate it.

tracker.track_duration(duration)
```
