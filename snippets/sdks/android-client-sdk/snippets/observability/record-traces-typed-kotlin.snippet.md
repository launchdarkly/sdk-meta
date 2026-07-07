---
id: android-client-sdk/observability/record-traces-typed-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Record custom traces with OTel-typed attributes (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
// Use OTel-typed attributes when exact OpenTelemetry typing is required
val span = LDObserve.startSpan(
    name = "database_query",
    attributes = Attributes.of(
        AttributeKey.stringKey("table"), "users",
        AttributeKey.stringKey("operation"), "select"
    )
)
performDatabaseQuery()
span.end()
```
