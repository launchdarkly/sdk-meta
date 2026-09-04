---
id: python-server-sdk/sdk-docs/init-async
sdk: python-server-sdk
kind: reference
lang: python
description: Create and start the experimental async client.
validation:
  scaffold: python-server-sdk/scaffolds/python-async-syntax-only
---

```python
# The async client is experimental. It is not recommended for production use,
# and its API may change or be removed in a future release.
client = AsyncLDClient(AsyncConfig("YOUR_SDK_KEY"))
await client.start()

# Or let the context manager start and close the client for you:
#   async with AsyncLDClient(AsyncConfig("YOUR_SDK_KEY")) as client:
#       ...
```
