---
id: android-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-building-users-2-x-syntax-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "2.x syntax (Kotlin) in section \"Understanding changes to building users\""
---

```kotlin
val user: LDUser = LDUser.Builder("userKey")
    .customString("properties", listOf("new", "priority"))
    .privateCustomNumber("counts", listOf(3, 5))
    .build()
```
