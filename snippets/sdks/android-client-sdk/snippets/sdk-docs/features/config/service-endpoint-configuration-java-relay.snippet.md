---
id: android-client-sdk/sdk-docs/features/config/service-endpoint-configuration-java-relay
sdk: android-client-sdk
kind: reference
lang: java
description: Service endpoint configuration example for Android.
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
