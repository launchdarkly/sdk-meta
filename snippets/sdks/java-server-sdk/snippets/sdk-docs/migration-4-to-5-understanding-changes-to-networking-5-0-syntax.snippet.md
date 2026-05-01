---
id: java-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-networking-5-0-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "5.0 syntax in section \"Understanding changes to networking\""
---

```java
// 5.0 model: setting connection and socket timeouts
LDConfig config = new LDConfig.Builder()
  .http(
    Components.httpConfiguration()
      .connectTimeout(Duration.ofSeconds(3))
      .socketTimeout(Duration.ofSeconds(4))
  )
  .build();

// 5.0 model: specifying an HTTP proxy with basic authentication
LDConfig config = new LDConfig.Builder()
  .http(
    Components.httpConfiguration()
      .proxyHostAndPort("my-proxy", 8080)
      .proxyAuth(Components.httpBasicAuthentication("user", "pass"))
  )
  .build();
```
