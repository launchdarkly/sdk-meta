---
id: java-server-sdk/sdk-docs/features/storing-data/index/persistent-store-data-system
sdk: java-server-sdk
kind: reference
lang: java
description: Persistent store configuration via the dataSystem builder for Java SDK 7.11+.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
import com.launchdarkly.sdk.server.*;

LDConfig config = new LDConfig.Builder()
    .dataSystem(
        Components.dataSystem().persistentStore(
            Components.persistentDataStore(Redis.dataStore())
        )
    )
    .build();

LDClient client = new LDClient("YOUR_SDK_KEY", config);
```
