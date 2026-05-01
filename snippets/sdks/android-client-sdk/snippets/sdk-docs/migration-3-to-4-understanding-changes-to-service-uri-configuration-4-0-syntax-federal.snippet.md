---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-service-uri-configuration-4-0-syntax-federal
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax, federal in section \"Understanding changes to service URI configuration\""
---

```java
LDConfig config = new LDConfig.Builder()
  .serviceEndpoints(
    Components.serviceEndpoints()
      .polling("https://clientsdk.launchdarkly.us")
      .streaming("https://clientstream.launchdarkly.us")
      .events("https://events.launchdarkly.us")
  )
  .build();
```
