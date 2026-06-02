---
id: java-server-sdk/sdk-docs/features/config/app-config
sdk: java-server-sdk
kind: reference
lang: java
description: Application metadata configuration example for Java.
---

```java
LDConfig config = new LDConfig.Builder()
  .applicationInfo(
    Components.applicationInfo()
      .applicationId("authentication-service")
      .applicationVersion("1.0.0")
  ).build();
```
