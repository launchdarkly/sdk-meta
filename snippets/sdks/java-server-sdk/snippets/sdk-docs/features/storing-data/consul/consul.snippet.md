---
id: java-server-sdk/sdk-docs/features/storing-data/consul/consul
sdk: java-server-sdk
kind: reference
lang: java
description: Consul feature store configuration example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
import com.launchdarkly.sdk.server.*;
import com.launchdarkly.sdk.server.integrations.*;

LDConfig config = new LDConfig.Builder()
  .dataStore(
    Components.persistentDataStore(
      Consul.dataStore().url(new URL("http://my-consul:8100"))
        .prefix("my-key-prefix")
    ).cacheSeconds(30)
  )
  .build();
LDClient client = new LDClient(sdkKey, config);
```
