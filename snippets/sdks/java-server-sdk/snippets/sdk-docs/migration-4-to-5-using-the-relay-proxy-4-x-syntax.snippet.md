---
id: java-server-sdk/sdk-docs/migration-4-to-5-using-the-relay-proxy-4-x-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "4.x syntax in section \"Using the Relay Proxy\""
---

```java
// 4.x model: proxy mode with streaming
URI relayUri = URI.create("http://my-relay-host:8000");
LDConfig config = new LDConfig.Builder()
  .baseUri(relayUri)
  .streamUri(relayUri)
  .eventsUri(relayUri) // if you want to proxy events
  .build();

// 4.x model: proxy mode with polling
URI relayUri = URI.create("http://my-relay-host:8000");
LDConfig config = new LDConfig.Builder()
  .stream(false)
  .baseUri(relayUri)
  .eventsUri(relayUri) // if you want to proxy events
  .build();

// 4.x model: daemon mode with a Redis database
LDConfig config = new LDConfig.Builder()
  .featureStore(
    Components.redisFeatureStore(URI.create("redis://my-redis-host"))
  )
  .useLdd(true)
  .build();
```
