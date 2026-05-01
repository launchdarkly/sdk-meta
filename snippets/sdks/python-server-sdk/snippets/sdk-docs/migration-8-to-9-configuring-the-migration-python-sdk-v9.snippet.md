---
id: python-server-sdk/sdk-docs/migration-8-to-9-configuring-the-migration-python-sdk-v9
sdk: python-server-sdk
kind: reference
lang: python
description: "Python SDK v9 in section \"Configuring the migration\""
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
