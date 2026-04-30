---
id: python-server-sdk/sdk-docs/postfork
sdk: python-server-sdk
kind: reference
lang: python
description: Worker-based-server initialization with postfork() reinitialization (Python SDK v9.11+).
validation:
  # Forking semantics aren't reproducible in the validator harness; fall
  # back to a parse-only check so we still catch syntax issues.
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
# 1. Create the client before forking.
ldclient.set_config(Config("YOUR_SDK_KEY"))
client = ldclient.get()

# 2. From the newly forked process, reinitialize the client by calling `postfork`.
# Examples for specific servers are shown below.
client.postfork()
```
