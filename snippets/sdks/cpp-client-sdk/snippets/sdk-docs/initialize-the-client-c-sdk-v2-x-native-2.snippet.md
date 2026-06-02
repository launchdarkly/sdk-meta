---
id: cpp-client-sdk/sdk-docs/initialize-the-client-c-sdk-v2-x-native-2
sdk: cpp-client-sdk
kind: reference
lang: c
description: "C SDK v2.x (native) in section \"Initialize the client\""
# Bucket C: cpp v2.x API surface no longer available in cpp-sdks v3 (the
# Dockerfile-pinned validator). See _sdk-docs-port-notes.md.
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```c
unsigned int maxwait = 10 * 1000; /* 10 seconds */
struct LDClient *client = LDClientInit(config, user, maxwait);
```
