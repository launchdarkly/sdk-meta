---
id: android-client-sdk/observability/distributed-tracing-thread-java
sdk: android-client-sdk
kind: reference
lang: java
file: app/src/main/java/com/launchdarkly/hello_android/SnippetObsJava.java
description: "Android observability plugin: Distributed tracing across a thread (Java)"
validation:
  scaffold: android-client-sdk/scaffolds/java-observability
---

```java
import io.opentelemetry.context.Context;
import io.opentelemetry.context.Scope;

Span parentSpan = LDObserve.Companion.startSpan("parentSpan", Attributes.empty());
try (Scope parentScope = parentSpan.makeCurrent()) {
    Context context = Context.current();

    new Thread(() -> {
        try (Scope childScope = context.makeCurrent()) {
            Span childSpan = LDObserve.Companion.startSpan("childSpan", Attributes.empty());
            // do work
            childSpan.end();
        }
    }).start();
} finally {
    parentSpan.end();
}
```
