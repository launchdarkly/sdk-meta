---
id: cpp-client-sdk/sdk-docs/evaluate-a-flag-c-sdk-v3-0-c-binding
sdk: cpp-client-sdk
kind: reference
lang: c
description: "C++ SDK v3.0 (C binding) in section \"Evaluate a flag\""
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```c
bool show_feature = LDClientSDK_BoolVariation(client, "example-flag-key", false);
if (show_feature) {
    // Application code to show the feature
} else {
    // The code to run if the feature is off
}
```
