---
id: cpp-server-sdk/sdk-docs/features/storing-data/redis/redis-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: Redis feature store configuration example for the C server SDK v2.x (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c

---

```c
#include <launchdarkly/api.h>
#include <launchdarkly/store/redis.h>

struct LDConfig *config = LDConfigNew("YOUR_SDK_KEY");
struct LDRedisConfig *redisConfig = LDRedisConfigNew();
struct LDStoreInterface *store = LDStoreInterfaceRedisNew(redisConfig);

LDConfigSetFeatureStoreBackend(config, store);
```
