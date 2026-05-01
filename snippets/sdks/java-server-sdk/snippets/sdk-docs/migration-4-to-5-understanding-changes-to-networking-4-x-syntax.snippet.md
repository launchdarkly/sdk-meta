---
id: java-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-networking-4-x-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "4.x syntax in section \"Understanding changes to networking\""
---

```java
// 4.x model: setting connection and socket timeouts
LDConfig config = new LDConfig.Builder()
  .connectTimeout(3)
  .socketTimeout(4)
  .build();

// 4.x model: specifying an HTTP proxy with basic authentication
LDConfig config = new LDConfig.Builder()
  .proxyHost("my-proxy")
  .proxyPort(8080)
  .proxyUsername("user")
  .proxyPassword("pass")
  .build();
```
