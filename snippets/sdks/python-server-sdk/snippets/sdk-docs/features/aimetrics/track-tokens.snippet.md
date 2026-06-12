---
id: python-server-sdk/sdk-docs/features/aimetrics/track-tokens
sdk: python-server-sdk
kind: reference
lang: python
description: Track token usage manually for the Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
# Track your own token usage.

# TokenUsage is provided by the LaunchDarkly AI SDK.
# Update the input, output, and total values
# with return values from your AI model generation.
tokens = TokenUsage(0, 0, 0)

tracker.track_tokens(tokens)
```
