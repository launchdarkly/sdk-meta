---
id: python-server-sdk/sdk-docs/features/migrations/migration-variation
sdk: python-server-sdk
kind: reference
lang: python
description: Migration stage evaluation (migration_variation) for Python SDK v9.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
context = Context.builder("example-context-key").build()

stage, tracker = ldclient.get().migration_variation('example-migration-flag-key', context, Stage.OFF)
```
