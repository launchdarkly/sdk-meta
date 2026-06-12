---
id: java-server-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode
sdk: java-server-sdk
kind: reference
lang: java
description: Proxy mode configuration example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
LDConfig config = new LDConfig.Builder()
  .serviceEndpoints(Components.serviceEndpoints()
    .relayProxy("https://your-relay-proxy.com:8030"))
  .build();
```
