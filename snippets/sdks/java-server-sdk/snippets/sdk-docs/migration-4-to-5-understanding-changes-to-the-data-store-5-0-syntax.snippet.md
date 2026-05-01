---
id: java-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-the-data-store-5-0-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "5.0 syntax in section \"Understanding changes to the data store\""
---

```java
// 5.0 model: use Redis, set custom Redis URI and key prefix, set cache TTL to 45 seconds
LDConfig config = new LDConfig.Builder()
  .dataStore(
    Components.persistentDataStore(
      Redis.dataStore()
        .uri(URI.create("redis://my-redis-host"))
        .prefix("my-prefix")
    ).cacheSeconds(45)
  )
  .build();
```
