---
id: android-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-building-users-3-0-syntax-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "3.0 syntax (Kotlin) in section \"Understanding changes to building users\""
---

```kotlin
val user: LDUser = LDUser.Builder("userKey")
    .custom("properties", LDValue.buildArray().add("new").add("priority").build())
    .privateCustom("counts", LDValue.buildArray().add(3).add(5).build())
    .build()
```
