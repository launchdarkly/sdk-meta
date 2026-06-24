---
id: android-client-sdk/sdk-docs/features/offlinemode/offline-mode-v5-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Offline mode example for Android SDK v5.x (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
val config: LDConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .offline(true)
    .build()

val client: LDClient = LDClient.init(application, config, context, 0);

// Or to switch an already-instantiated client to offline mode:
client.setOffline()
```
