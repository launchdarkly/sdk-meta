---
id: android-client-sdk/sdk-docs/features/logging/logging-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Timber debug logging example for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
if (BuildConfig.DEBUG) {
    Timber.plant(Timber.DebugTree())
}
```
