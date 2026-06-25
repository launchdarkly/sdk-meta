---
id: python-server-sdk/sdk-docs/features/filedata/flags-from-files
sdk: python-server-sdk
kind: reference
lang: python
description: File data source configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
import ldclient
from ldclient.config import Config
from ldclient.integrations import Files

data_source_callback = Files.new_data_source(paths=["file1.json", "file2.json"],
    auto_update=True)

config = Config('YOUR_SDK_KEY', update_processor_class=data_source_callback, send_events=False)

ldclient.set_config(config)
client = ldclient.get()
```
