---
id: android-client-sdk/sdk-docs/features/config/index-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: SDK configuration example for Android.
---

```kotlin
val ldConfig: LDConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .http(
      Components.httpConfiguration()
      .connectTimeoutMillis(5000)
    )
    .events(
      Components.sendEvents()
      .flushIntervalMillis(5000)
    )
    .build();
```
