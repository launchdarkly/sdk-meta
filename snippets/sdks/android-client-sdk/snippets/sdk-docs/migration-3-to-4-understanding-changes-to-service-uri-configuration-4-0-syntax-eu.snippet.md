---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-service-uri-configuration-4-0-syntax-eu
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax, EU in section \"Understanding changes to service URI configuration\""
---

```java
LDConfig config = new LDConfig.Builder()
  .serviceEndpoints(
    Components.serviceEndpoints()
      .polling("https://clientsdk.eu.launchdarkly.com")
      .streaming("https://clientstream.eu.launchdarkly.com")
      .events("https://events.eu.launchdarkly.com")
  )
  .build();
```
