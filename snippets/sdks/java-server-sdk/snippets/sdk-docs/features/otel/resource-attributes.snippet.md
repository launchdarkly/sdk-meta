---
id: java-server-sdk/sdk-docs/features/otel/resource-attributes
sdk: java-server-sdk
kind: reference
lang: java
description: Programmatic OpenTelemetry resource attribute configuration (Java).
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
Resource resource = Resource.getDefault().toBuilder()
    .put("launchdarkly.project_id", "YOUR_SDK_KEY")
    .build();

SdkTracerProvider tracerProvider = SdkTracerProvider.builder()
    .setResource(resource)
    .build();

OpenTelemetry openTelemetry = OpenTelemetrySdk.builder()
    .setTracerProvider(tracerProvider)
    .buildAndRegisterGlobal();
```
