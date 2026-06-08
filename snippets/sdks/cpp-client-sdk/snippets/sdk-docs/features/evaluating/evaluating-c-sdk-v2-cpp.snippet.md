---
id: cpp-client-sdk/sdk-docs/features/evaluating/evaluating-c-sdk-v2-cpp
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Flag evaluation example for C++ (client-side).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-cpp

---

```cpp
bool show_feature = client->boolVariation("example-flag-key", false);

if (show_feature) {
    // application code to show the feature
} else {
    // the code to run if the feature is off
}
```
