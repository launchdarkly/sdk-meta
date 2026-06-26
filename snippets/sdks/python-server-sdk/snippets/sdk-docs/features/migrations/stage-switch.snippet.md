---
id: python-server-sdk/sdk-docs/features/migrations/stage-switch
sdk: python-server-sdk
kind: reference
lang: python
description: Per-stage migration structure for Python SDK v9.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
# define the combination of reads and writes from the new and old systems
# that should occur at each migration stage

if stage == Stage.OFF:
  pass
elif stage == Stage.DUALWRITE:
  pass
elif stage == Stage.SHADOW:
  pass
elif stage == Stage.LIVE:
  pass
elif stage == Stage.RAMPDOWN:
  pass
elif stage == Stage.COMPLETE:
  pass
else:
  # throw an error
  pass
```
