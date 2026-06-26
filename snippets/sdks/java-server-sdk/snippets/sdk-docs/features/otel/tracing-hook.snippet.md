---
id: java-server-sdk/sdk-docs/features/otel/tracing-hook
sdk: java-server-sdk
kind: reference
lang: java
description: OpenTelemetry tracing hook configuration for the Java SDK.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
import com.launchdarkly.sdk.*;
import com.launchdarkly.sdk.server.*;
import com.launchdarkly.integrations.*;

TracingHook tracingHook = new TracingHook.Builder().build();

LDConfig config = new LDConfig.Builder()
    .hooks(
        Components.hooks().setHooks(Collections.singletonList(tracingHook)))
    .build();

LDClient client = new LDClient("YOUR_SDK_KEY", config);

```
