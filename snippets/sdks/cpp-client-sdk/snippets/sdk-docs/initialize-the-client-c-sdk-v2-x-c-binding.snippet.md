---
id: cpp-client-sdk/sdk-docs/initialize-the-client-c-sdk-v2-x-c-binding
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "C SDK v2.x (C++ binding) in section \"Initialize the client\""
# Bucket C: cpp v2.x API surface no longer available in cpp-sdks v3 (the
# Dockerfile-pinned validator). See _sdk-docs-port-notes.md.
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```cpp
LDConfig *config = LDConfigNew("example-mobile-key");
LDUser *user = LDUserNew("example-user-key");
```
