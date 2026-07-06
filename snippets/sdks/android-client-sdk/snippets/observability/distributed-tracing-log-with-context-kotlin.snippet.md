---
id: android-client-sdk/observability/distributed-tracing-log-with-context-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.kt
description: "Android observability plugin in section \"Distributed tracing\" (log with captured context, Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-observability-coroutine
---

```kotlin
import io.opentelemetry.context.Context

val parentSpan = LDObserve.startSpan("parentSpan", Attributes.empty())

// Capture the current context, which includes the active span
val context = Context.current()

launch(Dispatchers.IO) {
    val span = LDObserve.startSpan(
        name = "log-context-demo",
        attributes = Attributes.of(
            AttributeKey.stringKey("demo"), "log-with-context"
        )
    )
    // Capture span context while still on the originating thread.
    val capturedContext = span.makeCurrent().use { span.spanContext }
    span.end()
    // Simulate a detached thread where OTel context is lost automatically.
    // Span.current() here returns INVALID, so we pass the captured context explicitly.
    Thread {
        Span.wrap(capturedContext).makeCurrent().use {
            val childSpan = LDObserve.startSpan("child of log-context-demo", Attributes.empty())
            childSpan.end()
        }
        LDObserve.recordLog(
            message = text,
            severity = Severity.WARN,
            attributes = Attributes.of(
                AttributeKey.stringKey("source"), "detached-thread-demo"
            ),
            spanContext = capturedContext
        )
    }.start()
}

parentSpan.end()
```
