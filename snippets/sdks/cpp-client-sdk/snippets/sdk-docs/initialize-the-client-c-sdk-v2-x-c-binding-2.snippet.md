---
id: cpp-client-sdk/sdk-docs/initialize-the-client-c-sdk-v2-x-c-binding-2
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "C SDK v2.x (C++ binding) in section \"Initialize the client\""
---

```cpp
unsigned int maxwait = 10 * 1000; /* 10 seconds */
LDClientCPP *client = LDClientCPP::Init(config, user, maxwait);
```
