---
id: android-client-sdk/scaffolds/java-observability
sdk: android-client-sdk
kind: scaffold
lang: java
file: app/src/main/java/com/launchdarkly/hello_android/SnippetObsJava.java
description: |
  Type-check validator for Android observability-plugin Java doc
  fragments. Same `android-client` container / parse mode as the kotlin
  observability scaffold, but for Java: the body is spliced into a
  method of an `AppCompatActivity` host (so `this.getApplication()`
  resolves), the harness lifts any `import …;` the fragment carries to
  file scope, and the scaffold supplies the SDK / observability / OTel
  imports the fragments assume plus ambient stubs (`mobileKey`, a
  request `url`, an `executor`). OTel types are imported explicitly, not
  by wildcard, so the observability aar's
  `com.launchdarkly.observability.replay.Attributes` can't shadow
  `io.opentelemetry.api.common.Attributes`.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, type-checked against the plugin.
validation:
  runtime: android-client
  entrypoint: app/src/main/java/com/launchdarkly/hello_android/SnippetObsJava.java
  env:
    SNIPPET_CHECK: parse
---

```java
package com.launchdarkly.hello_android;

import androidx.appcompat.app.AppCompatActivity;
import java.util.Collections;
import java.util.List;
import java.util.concurrent.Executor;
import com.launchdarkly.sdk.LDContext;
import com.launchdarkly.sdk.android.LDConfig;
import com.launchdarkly.sdk.android.LDClient;
import com.launchdarkly.sdk.android.Components;
import com.launchdarkly.sdk.android.integrations.Plugin;
import com.launchdarkly.sdk.android.LDConfig.Builder.AutoEnvAttributes;
import com.launchdarkly.observability.plugin.Observability;
import com.launchdarkly.observability.api.ObservabilityOptions;
import com.launchdarkly.observability.sdk.LDObserve;
import io.opentelemetry.api.common.Attributes;
import io.opentelemetry.api.common.AttributeKey;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.propagation.W3CTraceContextPropagator;
import io.opentelemetry.context.Context;
import io.opentelemetry.context.Scope;
import okhttp3.Request;

@SuppressWarnings("unused")
public class SnippetObsJava extends AppCompatActivity {
    private String mobileKey = "";
    private String url = "";
    private Executor executor = null;

    void wrappee() {
{{ body }}
    }
}
```
