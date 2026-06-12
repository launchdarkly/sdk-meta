---
id: java-server-sdk/sdk-docs/features/storing-data/redis/redis
sdk: java-server-sdk
kind: reference
lang: java
description: Redis feature store configuration example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
import com.launchdarkly.sdk.server.*;
import com.launchdarkly.sdk.server.integrations.*;

LDConfig config = new LDConfig.Builder()
  .dataStore(
    Components.persistentDataStore(
      Redis.dataStore().uri(URI.create("redis://my-redis:6379"))
        .prefix("my-key-prefix")
    ).cacheSeconds(30)
  )
  .build();
LDClient client = new LDClient(sdkKey, config);
```
