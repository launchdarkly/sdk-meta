---
id: android-client-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-java
sdk: android-client-sdk
kind: reference
lang: java
description: Proxy mode configuration example for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java
LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
  .mobileKey("example-mobile-key")
  .serviceEndpoints(
    Components.serviceEndpoints()
      .relayProxy("https://your-relay-proxy.com:8030")
  )
  .build();
```
