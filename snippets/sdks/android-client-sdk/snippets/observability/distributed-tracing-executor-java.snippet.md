---
id: android-client-sdk/observability/distributed-tracing-executor-java
sdk: android-client-sdk
kind: reference
lang: java
file: app/src/main/java/com/launchdarkly/hello_android/SnippetObsJava.java
description: "Android observability plugin: Distributed tracing with an executor (Java)"
validation:
  scaffold: android-client-sdk/scaffolds/java-observability
---

```java
import io.opentelemetry.context.Context;
import io.opentelemetry.context.Scope;

Span parentSpan = LDObserve.Companion.startSpan("parentSpan", new HashMap<>());
try (Scope parentScope = parentSpan.makeCurrent()) {
    // Now parentSpan is active in Context.current()
    Context ctx = Context.current();

    // Later, in another thread, restore the context
    executor.execute(() -> {
        try (Scope scope = ctx.makeCurrent()) {
            Span nestedSpan = LDObserve.Companion.startSpan("nestedSpan", new HashMap<>());
            // do work — nestedSpan is a child of parentSpan
            nestedSpan.end();
        }
    });
} finally {
    parentSpan.end();
}
```
