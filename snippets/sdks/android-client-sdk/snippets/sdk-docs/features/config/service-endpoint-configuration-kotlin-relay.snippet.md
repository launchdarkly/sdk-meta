---
id: android-client-sdk/sdk-docs/features/config/service-endpoint-configuration-kotlin-relay
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Service endpoint configuration example for Android.
---

```kotlin
val ldConfig: LDConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
  .mobileKey("example-mobile-key")
  .serviceEndpoints(
    Components.serviceEndpoints()
      .relayProxy("https://your-relay-proxy.com:8030")
  )
  .build();
```
