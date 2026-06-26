---
id: cpp-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: Daemon mode configuration example for the v2.x C server SDK.
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c
---

```c
struct LDConfig *config = LDConfigNew("YOUR_SDK_KEY");

struct LDStoreInterface *store = ConstructYourFeatureStoreInterface();

LDConfigSetFeatureStoreBackend(config, store);
LDConfigSetUseLDD(config, true);
```
