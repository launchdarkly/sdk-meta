---
id: cpp-client-sdk/sdk-docs/initialize-the-client-c-sdk-v2-x-c-binding-2
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "C SDK v2.x (C++ binding) in section \"Initialize the client\""
# Bucket C: cpp v2.x API surface no longer available in cpp-sdks v3 (the
# Dockerfile-pinned validator). See _sdk-docs-port-notes.md.
---

```cpp
unsigned int maxwait = 10 * 1000; /* 10 seconds */
LDClientCPP *client = LDClientCPP::Init(config, user, maxwait);
```
