---
id: java-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode-datastore
sdk: java-server-sdk
kind: reference
lang: java
description: Daemon mode configuration example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
LDConfig config = new LDConfig.Builder()
  .dataStore(
    Components.persistentDataStore(
      SomeDatabaseName.DataStore(storeOptions)
    )
  )
  .dataSource(Components.externalUpdatesOnly())
  .build();
```
