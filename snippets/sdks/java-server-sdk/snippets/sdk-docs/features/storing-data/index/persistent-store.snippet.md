---
id: java-server-sdk/sdk-docs/features/storing-data/index/persistent-store
sdk: java-server-sdk
kind: reference
lang: java
description: Persistent feature store configuration example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
import com.launchdarkly.sdk.server.*;
import com.launchdarkly.sdk.server.integrations.*;

LDConfig config = new LDConfig.Builder()
  .dataStore(
    Components.persistentDataStore(
      SomeDatabaseName.dataStore(storeOptions)
    )
  )
  .build();
LDClient client = new LDClient(sdkKey, config);
```
