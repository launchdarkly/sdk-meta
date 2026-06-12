---
id: java-server-sdk/sdk-docs/features/storing-data/dynamodb/dynamodb
sdk: java-server-sdk
kind: reference
lang: java
description: DynamoDB feature store configuration example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
import com.launchdarkly.sdk.server.*;
import com.launchdarkly.sdk.server.integrations.*;

LDConfig config = new LDConfig.Builder()
  .dataStore(
    Components.persistentDataStore(
      DynamoDb.dataStore("my-table")
    ).cacheSeconds(30)
  )
  .build();
LDClient client = new LDClient(sdkKey, config);
```
