---
id: android-client-sdk/sdk-docs/features/config/app-config-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Application metadata configuration example for Android.
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
val ldConfig: LDConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .applicationInfo(
        Components.applicationInfo()
            .applicationId("authentication-service")
            .applicationName("Authentication-Service")
            .applicationVersion("1.0.0")
            .applicationVersionName("v1")
    )
    .build()
```
