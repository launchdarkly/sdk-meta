---
id: python-server-sdk/sdk-docs/features/aimetrics/track-feedback
sdk: python-server-sdk
kind: reference
lang: python
description: Track output satisfaction rate for the Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
# Track your own output satisfaction rate.

# Pass in FeedbackKind.Positive or FeedbackKind.Negative.
tracker.track_feedback({"kind": FeedbackKind.Positive})
```
