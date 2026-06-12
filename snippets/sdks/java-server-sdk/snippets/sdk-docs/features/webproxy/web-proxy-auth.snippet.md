---
id: java-server-sdk/sdk-docs/features/webproxy/web-proxy-auth
sdk: java-server-sdk
kind: reference
lang: java
description: Web proxy configuration with authentication for the Java SDK.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
LDConfig config = new LDConfig.Builder()
  .http(Components.httpConfiguration()
    .proxyHostAndPort("my-proxy-host", 8080)
    .proxyAuth(Components.httpBasicAuthentication("username", "password")))
  .build();
```
