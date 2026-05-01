---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-networking-4-0-syntax
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax in section \"Understanding changes to networking\""
---

```java
// Setting connection timeout
LDConfig config = new LDConfig.Builder()
  .http(
    Components.httpConfiguration()
      .connectTimeoutMillis(3000)
  )
  .build();
```
