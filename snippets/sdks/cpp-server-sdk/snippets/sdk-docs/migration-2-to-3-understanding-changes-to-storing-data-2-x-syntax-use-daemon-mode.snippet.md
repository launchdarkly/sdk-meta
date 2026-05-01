---
id: cpp-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-storing-data-2-x-syntax-use-daemon-mode
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: "2.x syntax, use daemon mode in section \"Understanding changes to storing data\""
---

```cpp
struct LDConfig *config = LDConfigNew("YOUR_SDK_KEY");

struct LDStoreInterface *store = ConstructYourFeatureStoreInterface();

LDConfigSetFeatureStoreBackend(config, store);
LDConfigSetUseLDD(config, true);

```
