---
id: cpp-server-sdk/sdk-docs/features/config/service-endpoint-configuration-c-sdk-v2-relay
sdk: cpp-server-sdk
kind: reference
lang: c
description: Service endpoint configuration example for C++ (server-side).
---

```c
struct LDConfig *config = LDConfigNew("YOUR_SDK_KEY");
LDConfigSetStreamURI(config, "https://your-relay-proxy.com:8030");
LDConfigSetBaseURI(config, "https://your-relay-proxy.com:8030");
LDConfigSetEventsURI(config, "https://your-relay-proxy.com:8030");
```
