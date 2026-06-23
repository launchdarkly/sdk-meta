---
id: android-client-sdk/sdk-docs/features/identify/identify-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Identify example for the Android SDK v4.0+ (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only
---

```kotlin
val updatedContext: LDContext = LDContext.builderFromContext(context)
    .set("email", "sandy@example.com")
    .build()

client.identify(updatedContext)
```
