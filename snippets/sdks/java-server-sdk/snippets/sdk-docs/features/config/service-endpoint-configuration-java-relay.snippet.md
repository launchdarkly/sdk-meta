---
id: java-server-sdk/sdk-docs/features/config/service-endpoint-configuration-java-relay
sdk: java-server-sdk
kind: reference
lang: java
description: Service endpoint configuration example for Java.
---

```java
LDConfig config = new LDConfig.Builder()
  .serviceEndpoints(Components.serviceEndpoints()
    .relayProxy("https://your-relay-proxy.com:8030"))
  .build();
```
