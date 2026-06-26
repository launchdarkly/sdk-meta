---
id: android-client-sdk/sdk-docs/features/contextconfig/context-example-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Context example for Android SDK v4.0+ (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
val context: LDContext = LDContext.builder("example-context-key")
    .set("email", "sandy@example.com")
    .set("firstName", "Sandy")
    .set("lastName", "Smith")
    .set("group", "Global Health Services")
    .build()
```
