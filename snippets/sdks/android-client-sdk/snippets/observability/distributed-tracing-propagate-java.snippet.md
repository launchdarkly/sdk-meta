---
id: android-client-sdk/observability/distributed-tracing-propagate-java
sdk: android-client-sdk
kind: reference
lang: java
file: app/src/main/java/com/launchdarkly/hello_android/SnippetObsJava.java
description: "Android observability plugin: Distributed tracing, propagate headers (Java)"
validation:
  scaffold: android-client-sdk/scaffolds/java-observability
---

```java
import io.opentelemetry.api.trace.propagation.W3CTraceContextPropagator;
import io.opentelemetry.context.Context;
import okhttp3.Request;

W3CTraceContextPropagator propagator = W3CTraceContextPropagator.getInstance();
Request.Builder requestBuilder = new Request.Builder().url(url);

propagator.inject(
    Context.current(),
    requestBuilder,
    (carrier, key, value) -> carrier.addHeader(key, value)
);

Request request = requestBuilder.build();
```
