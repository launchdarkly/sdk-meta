---
id: android-client-sdk/observability/distributed-tracing-restore-context-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin: Distributed tracing, restore context (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability
---

```kotlin
import io.opentelemetry.context.Context

// Capture the current context
val ctx = Context.current()

// Later, in another scope, restore the context
ctx.makeCurrent().use {
    val span = LDObserve.startSpan("nestedSpan", Attributes.empty())
    // do work
    span.end()
}
```
