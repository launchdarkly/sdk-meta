---
id: python-server-sdk/sdk-docs/features/config/migration-config
sdk: python-server-sdk
kind: reference
lang: python
description: Migration configuration example for the Python SDK v9 — read/write methods, execution order, latency/error tracking.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
from ldclient import Result, MigratorBuilder, ExecutionOrder

builder = MigratorBuilder(ldclient.get())

builder.read(lambda _: Result.success("read old"), lambda _: Result.success("read new"), lambda lhs, rhs: lhs == rhs)
builder.write(lambda _: Result.success("write old"), lambda _: Result.success("write new"))

builder.read_execution_order(ExecutionOrder.PARALLEL)
# could also use ExecutionOrder.SERIAL, ExecutionOrder.RANDOM

builder.track_latency(True) # defaults to True
builder.track_errors(True)  # defaults to True

result = builder.build()

```
