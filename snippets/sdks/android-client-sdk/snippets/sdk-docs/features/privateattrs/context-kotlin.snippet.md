---
id: android-client-sdk/sdk-docs/features/privateattrs/context-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Marking context attributes private with the context builder for Android SDK v4.0+ (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
val context: LDContext = LDContext.builder("example-context-key")
    .set("email", "sandy@example.com")
    .set("name", "Sandy")
    .set("group", "Global Health Services")
    .privateAttributes("name", "group")
    .build()
```
