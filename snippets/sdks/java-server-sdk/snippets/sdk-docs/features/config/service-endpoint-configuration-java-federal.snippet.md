---
id: java-server-sdk/sdk-docs/features/config/service-endpoint-configuration-java-federal
sdk: java-server-sdk
kind: reference
lang: java
description: Service endpoint configuration example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
LDConfig config = new LDConfig.Builder()
  .serviceEndpoints(Components.serviceEndpoints()
    .streaming("https://stream.launchdarkly.us")
    .polling("https://sdk.launchdarkly.us")
    .events("https://events.launchdarkly.us"))
  .build();
```
