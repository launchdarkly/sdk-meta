---
id: android-client-sdk/sdk-docs/features/anonymous/anonymous-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Anonymous context example for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only
---

```kotlin
val context: LDContext = LDContext.builder("example-context-key")
    .anonymous(true)
    .build()
```
