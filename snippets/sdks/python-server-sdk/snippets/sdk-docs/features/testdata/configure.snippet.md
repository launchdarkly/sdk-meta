---
id: python-server-sdk/sdk-docs/features/testdata/configure
sdk: python-server-sdk
kind: reference
lang: python
description: Test data source configuration for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
from ldclient.integrations.test_data import TestData
from ldclient import LDClient, Config

td = TestData.data_source()
# You can set any initial flag states here with td.update
client = LDClient(config=Config(sdk_key, update_processor_class = td))
```
