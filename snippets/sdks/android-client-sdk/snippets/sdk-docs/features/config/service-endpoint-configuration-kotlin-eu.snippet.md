---
id: android-client-sdk/sdk-docs/features/config/service-endpoint-configuration-kotlin-eu
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Service endpoint configuration example for Android.
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
val ldConfig: LDConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
  .mobileKey("example-mobile-key")
  .serviceEndpoints(
    Components.serviceEndpoints()
      .streaming("https://clientstream.eu.launchdarkly.com")
      .polling("https://clientsdk.eu.launchdarkly.com")
      .events("https://events.eu.launchdarkly.com")
  )
  .build();
```
