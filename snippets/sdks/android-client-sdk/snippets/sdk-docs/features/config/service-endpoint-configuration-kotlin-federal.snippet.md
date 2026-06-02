---
id: android-client-sdk/sdk-docs/features/config/service-endpoint-configuration-kotlin-federal
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
      .streaming("https://clientstream.launchdarkly.us")
      .polling("https://clientsdk.launchdarkly.us")
      .events("https://events.launchdarkly.us")
  )
  .build();
```
