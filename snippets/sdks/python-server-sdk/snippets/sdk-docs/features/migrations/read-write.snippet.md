---
id: python-server-sdk/sdk-docs/features/migrations/read-write
sdk: python-server-sdk
kind: reference
lang: python
description: Migration read and write example for Python SDK v9.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
from ldclient import Context
from ldclient.migrations import Stage

context = Context.builder("example-context-key").build()


# this is the migration stage to use if the flag's migration stage
# is not available from LaunchDarkly
default_stage = Stage.OFF

migrator = builder.build()

# when you need to perform a read in your application
migrator.read(
    'example-migration-flag-key',
    context,
    default_stage
)

# when you need to perform a write in your application
migrator.write(
    'example-migration-flag-key',
    context,
    default_stage
)
```
