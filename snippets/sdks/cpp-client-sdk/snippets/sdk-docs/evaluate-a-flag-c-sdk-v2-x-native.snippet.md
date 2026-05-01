---
id: cpp-client-sdk/sdk-docs/evaluate-a-flag-c-sdk-v2-x-native
sdk: cpp-client-sdk
kind: reference
lang: c
description: "C SDK v2.x (native) in section \"Evaluate a flag\""
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```c
bool show_feature = LDBoolVariation(client, "example-flag-key", false);
if (show_feature) {
    // Application code to show the feature
} else {
    // The code to run if the feature is off
}
```
