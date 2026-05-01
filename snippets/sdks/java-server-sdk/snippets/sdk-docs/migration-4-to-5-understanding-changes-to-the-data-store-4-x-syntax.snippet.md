---
id: java-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-the-data-store-4-x-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "4.x syntax in section \"Understanding changes to the data store\""
---

```java
// 4.x model: use Redis, set custom Redis URI and key prefix, set cache TTL to 45 seconds
LDConfig config = new LDConfig.Builder()
  .featureStore(
    Components.redisFeatureStore(URI.create("redis://my-redis-host"))
      .prefix("my-prefix")
      .cacheTime(45, TimeUnit.SECOND)
  )
  .build();
```
