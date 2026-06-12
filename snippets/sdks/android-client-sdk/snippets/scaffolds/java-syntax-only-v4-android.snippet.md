---
id: android-client-sdk/scaffolds/java-syntax-only-v4-android
sdk: android-client-sdk
kind: scaffold
lang: java
file: app/src/main/java/com/launchdarkly/hello_android/SnippetV4.java
description: |
  Parse-and-type-check validator for Android Java doc fragments that
  target the v4.x configuration surface (`new LDConfig.Builder()` with
  no arguments — v5 made the `AutoEnvAttributes` constructor argument
  mandatory, so v4-era fragments cannot compile against the v5 aar the
  `android-client` validator container pre-bakes).

  Unlike `java-syntax-only-v4` (which routes through the `jvm`
  validator and therefore needs its own CI matrix row with a server
  key), this scaffold stays inside the `android-client` container in
  `SNIPPET_CHECK=parse` mode: nested stub classes declare just the v4
  builder surface the fragments reference, shadowing nothing — the
  file deliberately does NOT import `com.launchdarkly.sdk.android`, so
  the v5 aar's same-named types never collide with the stubs. The
  existing android sdk-docs CI row picks these snippets up with no
  workflow changes.

  Shared `com.launchdarkly.sdk` types (`LDContext`) are real — they
  ship inside the v5 aar's java-sdk-common dependency and are
  version-stable across v4/v5. A `getApplication()` stub on the host
  class mirrors the Activity accessor the init fragments call via
  `this.getApplication()`.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: android-client
  entrypoint: app/src/main/java/com/launchdarkly/hello_android/SnippetV4.java
  env:
    SNIPPET_CHECK: parse
---

```java
package com.launchdarkly.hello_android;

import com.launchdarkly.sdk.LDContext;
import java.util.Map;
import java.util.HashMap;
import java.util.concurrent.Future;

// Stub of the v4-era android config surface. Only the members the doc
// fragments call are declared; everything returns a stub so the
// chained-builder shape type-checks. Nested inside the host class so
// the unqualified names resolve from the wrappee body without
// polluting the package namespace (other files in the pre-baked
// project reference the real v5 `LDConfig`).
class SnippetV4 {
    static class LDConfig {
        static class Builder {
            Builder() {}
            Builder mobileKey(String key) { return this; }
            Builder secondaryMobileKeys(Map<String, String> keys) { return this; }
            Builder events(EventProcessorBuilder eventsConfig) { return this; }
            LDConfig build() { return new LDConfig(); }
        }
    }

    static class EventProcessorBuilder {
        EventProcessorBuilder allAttributesPrivate(boolean allAttributesPrivate) { return this; }
        EventProcessorBuilder privateAttributes(String... attributeNames) { return this; }
    }

    static class Components {
        static EventProcessorBuilder sendEvents() { return new EventProcessorBuilder(); }
    }

    // Stub of the v4 android LDClient init surface. The application
    // parameter is typed Object so `this.getApplication()` (stubbed
    // below) satisfies it without an android Activity host. The
    // non-blocking overload (no startWaitSeconds) returns a Future,
    // matching the real v4 signature.
    static class LDClient {
        static LDClient init(Object application, LDConfig config, LDContext context, int startWaitSeconds) {
            return new LDClient();
        }
        static Future<LDClient> init(Object application, LDConfig config, LDContext context) {
            return null;
        }
    }

    // Mirrors the Activity accessor the init fragments call via
    // `this.getApplication()`.
    private Object getApplication() { return null; }

    @SuppressWarnings({"unused", "ConstantConditions"})
    private void wrappee() {
        if (false) {
{{ body }}
        }
    }
}
```
