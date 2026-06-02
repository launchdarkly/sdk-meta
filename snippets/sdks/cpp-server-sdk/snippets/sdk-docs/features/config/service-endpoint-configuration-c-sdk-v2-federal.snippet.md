---
id: cpp-server-sdk/sdk-docs/features/config/service-endpoint-configuration-c-sdk-v2-federal
sdk: cpp-server-sdk
kind: reference
lang: c
description: Service endpoint configuration example for C++ (server-side).
---

```c
struct LDConfig *config = LDConfigNew("YOUR_SDK_KEY");
LDConfigSetStreamURI(config, "https://stream.launchdarkly.us");
LDConfigSetBaseURI(config, "https://sdk.launchdarkly.us");
LDConfigSetEventsURI(config, "https://events.launchdarkly.us");
```
