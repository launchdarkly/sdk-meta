---
id: cpp-client-sdk/sdk-docs/evaluate-a-flag-c-sdk-v3-0-native
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "C++ SDK v3.0 (native) in section \"Evaluate a flag\""
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
