---
id: cpp-server-sdk/sdk-docs/c-c---evaluate-a-context-c-sdk-v3-0-native
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: "C++ SDK v3.0 (native) in section \"Evaluate a context\""
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```cpp
auto context = ContextBuilder().Kind("user", "example-user-key").Name("Sandy").Build();

bool show_feature = client.BoolVariation(context, "example-flag-key", false);

if (show_feature) {
    // application code to show the feature
} else {
    // the code to run if the feature is off
}
```
