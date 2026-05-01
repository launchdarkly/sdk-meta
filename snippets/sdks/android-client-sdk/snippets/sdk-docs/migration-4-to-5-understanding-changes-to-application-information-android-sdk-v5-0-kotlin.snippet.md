---
id: android-client-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-application-information-android-sdk-v5-0-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Android SDK v5.0 (Kotlin) in section \"Understanding changes to application information\""
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
