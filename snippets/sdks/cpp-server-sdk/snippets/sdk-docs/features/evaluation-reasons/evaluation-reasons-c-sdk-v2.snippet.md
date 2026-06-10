---
id: cpp-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: Flag evaluation reason example for the C server SDK v2.x (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c

---

```c
struct LDDetails details;
LDBoolean value;

value = LDBoolVariation(client, user, "example-flag-key", LDBooleanFalse, &details);

/* inspect details here */
if (details.reason == LD_RULE_MATCH) {
    /* ... */
}

LDDetailsClear(&details);
```
