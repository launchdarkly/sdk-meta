---
id: cpp-client-sdk/sdk-docs/features/evaluating/evaluating-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Flag evaluation example for C++ (client-side).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
bool show_feature = client.BoolVariation("example-flag-key", false);
if (show_feature) {
    // Application code to show the feature
} else {
    // The code to run if the feature is off
}
```
