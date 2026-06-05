---
id: cpp-server-sdk/sdk-docs/features/evaluating/evaluating-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Flag evaluation example for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```c
bool value = LDServerSDK_BoolVariation(client, context, "example-flag-key", false);
```
