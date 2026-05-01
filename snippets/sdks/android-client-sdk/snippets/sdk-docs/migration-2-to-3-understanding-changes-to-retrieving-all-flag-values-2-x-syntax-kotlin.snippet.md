---
id: android-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-retrieving-all-flag-values-2-x-syntax-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "2.x syntax (Kotlin) in section \"Understanding changes to retrieving all flag values\""
---

```kotlin
val flagValues: Map<String, ?> = client.allFlags()
for (flag in flagValues.entries) {
    when (flag.value) {
        is Boolean -> { } // Do something with boolean flag
        is Float -> { } // Do something with numeric flag
        is String -> { } // Do something with string (or serialized Json) flag
    }
}
```
