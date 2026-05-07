---
id: cpp-client-sdk/sdk-docs/evaluate-a-flag-c-sdk-v2-x-c-binding
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "C SDK v2.x (C++ binding) in section \"Evaluate a flag\""
# Bucket C: cpp v2.x API surface no longer available in cpp-sdks v3 (the
# Dockerfile-pinned validator). See _sdk-docs-port-notes.md.
---

```cpp
bool show_feature = client->boolVariation("example-flag-key", false);

if (show_feature) {
    // Application code to show the feature
} else {
    // The code to run if the feature is off
}
```
