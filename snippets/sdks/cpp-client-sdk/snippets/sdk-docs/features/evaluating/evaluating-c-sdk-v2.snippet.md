---
id: cpp-client-sdk/sdk-docs/features/evaluating/evaluating-c-sdk-v2
sdk: cpp-client-sdk
kind: reference
lang: c
description: Flag evaluation example for C++ (client-side).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-c

---

```c
bool show_feature = LDBoolVariation(client, "example-flag-key", false);
if (show_feature) {
    // application code to show the feature
} else {
    // the code to run if the feature is off
}
```
