---
id: java-server-sdk/sdk-docs/features/config/service-endpoint-configuration-java-eu
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
    .streaming("https://stream.eu.launchdarkly.com")
    .polling("https://sdk.eu.launchdarkly.com")
    .events("https://events.eu.launchdarkly.com"))
  .build();
```
