---
id: cpp-server-sdk/sdk-docs/features/evaluating/evaluating-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: Flag evaluation example for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c

---

```c
bool value = LDBoolVariation(client, user, "example-flag-key", false, NULL);
```
