---
id: android-client-sdk/observability/record-traces-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Record custom traces (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
// Start a span with custom properties
val span = LDObserve.startSpan(
    name = "database_query",
    properties = mapOf(
        "table" to "users",
        "operation" to "select"
    )
)

// Perform your operation
performDatabaseQuery()

// Always end the span
span.end()
```
