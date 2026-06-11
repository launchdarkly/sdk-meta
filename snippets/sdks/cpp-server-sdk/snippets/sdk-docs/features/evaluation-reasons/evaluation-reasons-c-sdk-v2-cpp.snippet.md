---
id: cpp-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-c-sdk-v2-cpp
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Flag evaluation reason example for the C server SDK v2.x (C++ binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-cpp

---

```cpp
LDDetails details;

bool value = LDBoolVariation(client, user, "example-flag-key", false, &details);

/* inspect details here */
if (details.reason == LD_RULE_MATCH) {
    /* ... */
}

LDDetailsClear(&details);
```
