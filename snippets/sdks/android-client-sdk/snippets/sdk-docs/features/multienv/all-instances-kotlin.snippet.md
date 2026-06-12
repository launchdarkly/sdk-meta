---
id: android-client-sdk/sdk-docs/features/multienv/all-instances-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Calls that affect all LDClient instances for Android SDK v4.0+ (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
val coreInstance: LDClient = LDClient.getForMobileKey("core")

// Calls affect all LDClient Instances
coreInstance.identify(context)
coreInstance.flush()
coreInstance.setOffline()
coreInstance.setOnline()
coreInstance.close()
```
