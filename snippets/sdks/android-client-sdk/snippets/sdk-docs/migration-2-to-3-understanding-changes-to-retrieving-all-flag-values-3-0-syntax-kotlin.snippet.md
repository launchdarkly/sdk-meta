---
id: android-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-retrieving-all-flag-values-3-0-syntax-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "3.0 syntax (Kotlin) in section \"Understanding changes to retrieving all flag values\""
---

```kotlin
val flagValues: Map<String, ?> = client.allFlags()
for (flag in flagValues.entries) {
    when (flag.value.type) {
        LDValueType.NULL -> { } // Do something with flag missing value
        LDValueType.BOOLEAN -> { } // Do something with boolean flag
        LDValueType.NUMBER -> { } // Do something with numeric flag
        LDValueType.STRING -> { } // Do something with string flag
        LDValueType.ARRAY -> { } // Do something with array flag
        LDValueType.OBJECT -> { } // Do something with object flag
    }
}
```
