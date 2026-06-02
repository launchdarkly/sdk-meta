---
id: java-server-sdk/sdk-docs/features/config/index
sdk: java-server-sdk
kind: reference
lang: java
description: SDK configuration example for Java.
---

```java
LDConfig config = new LDConfig.Builder()
  .http(
    Components.httpConfiguration()
      .connectTimeout(Duration.ofSeconds(3))
      .socketTimeout(Duration.ofSeconds(3))
  )
  .events(
    Components.sendEvents()
      .flushInterval(Duration.ofSeconds(10))
  )
  .build();
LDClient client = new LDClient("YOUR_SDK_KEY", config);
```
