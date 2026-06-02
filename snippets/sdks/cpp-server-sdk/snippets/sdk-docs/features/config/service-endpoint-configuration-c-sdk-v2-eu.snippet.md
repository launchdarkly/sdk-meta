---
id: cpp-server-sdk/sdk-docs/features/config/service-endpoint-configuration-c-sdk-v2-eu
sdk: cpp-server-sdk
kind: reference
lang: c
description: Service endpoint configuration example for C++ (server-side).
---

```c
struct LDConfig *config = LDConfigNew("YOUR_SDK_KEY");
LDConfigSetStreamURI(config, "https://stream.eu.launchdarkly.com");
LDConfigSetBaseURI(config, "https://sdk.eu.launchdarkly.com");
LDConfigSetEventsURI(config, "https://events.eu.launchdarkly.com");
```
