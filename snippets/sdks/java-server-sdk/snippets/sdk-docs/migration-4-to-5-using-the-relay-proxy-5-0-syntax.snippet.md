---
id: java-server-sdk/sdk-docs/migration-4-to-5-using-the-relay-proxy-5-0-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "5.0 syntax in section \"Using the Relay Proxy\""
---

```java
// 5.0 model: proxy mode with streaming
URI relayUri = URI.create("http://my-relay-host:8000");
LDConfig config = new LDConfig.Builder()
  .dataSource(Components.streamingDataSource().baseUri(relayUri))
  .events(Components.sendEvents().baseUri(relayUri)) // if you want to proxy events
  .build();

// 5.0 model: proxy mode with polling
URI relayUri = URI.create("http://my-relay-host:8000");
LDConfig config = new LDConfig.Builder()
  .dataSource(Components.pollingDataSource().baseUri(relayUri))
  .events(Components.sendEvents().baseUri(relayUri)) // if you want to proxy events
  .build();

// 5.0 model: daemon mode with a Redis database
LDConfig config = new LDConfig.Builder()
  .dataSource(Components.externalUpdatesOnly()) // replaces "useLdd"
  .dataStore(
    Components.persistentDataStore(
      Redis.dataStore().uri(URI.create("redis://my-redis-host"))
    )
  )
  .build();
```
