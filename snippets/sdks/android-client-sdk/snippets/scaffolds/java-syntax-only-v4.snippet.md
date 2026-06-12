---
id: android-client-sdk/scaffolds/java-syntax-only-v4
sdk: android-client-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/Snippet.java
description: |
  Parse-and-type-check validator for Android Java doc fragments that
  target the v4.x API surface (`new LDConfig.Builder()` with no
  arguments — v5 made the `AutoEnvAttributes` constructor argument
  mandatory, so v4-era fragments cannot compile against the v5 aar
  the `android-client` validator container pre-bakes).

  Routes through the `jvm` validator instead, following the same stub
  approach as the legacy C SDK validators (cpp-client-v2-c and
  friends): nested stub classes declare just the v4 android surface
  the fragments reference. The shared `com.launchdarkly.sdk` types
  (`EvaluationDetail`, `EvaluationReason`, `LDContext`) come from the
  real java-sdk-common classes on the jvm validator's classpath —
  only the `com.launchdarkly.sdk.android` types are stubbed, since
  the android aar itself can't be resolved by the jvm validator's
  Maven path.

  The nested `LDConfig` / `LDClient` classes shadow any same-named
  imports inside the wrappee body, and `getApplication()` /
  `context` / `secondsToBlock` mirror the ambient names the doc
  fragments assume an Activity host provides.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: jvm
  entrypoint: src/main/java/com/launchdarkly/Snippet.java
---

```java
package com.launchdarkly;

import com.launchdarkly.sdk.*;

public class Snippet {
    // Stub of the v4 android LDConfig builder chain. Only the members
    // the doc fragments call are declared; everything returns a stub
    // so the chained-builder shape type-checks.
    static class LDConfig {
        static class Builder {
            Builder() {}
            Builder mobileKey(String key) { return this; }
            Builder evaluationReasons(boolean enabled) { return this; }
            Builder offline(boolean offline) { return this; }
            LDConfig build() { return new LDConfig(); }
        }
    }

    // Stub of the v4 android LDClient surface. The application
    // parameter is typed Object so `this.getApplication()` (stubbed
    // below) satisfies it without an android.app.Application class.
    static class LDClient {
        static LDClient init(Object application, LDConfig config, LDContext context, int secondsToBlock) {
            return new LDClient();
        }
        EvaluationDetail<Boolean> boolVariationDetail(String key, boolean defaultValue) {
            return null;
        }
        EvaluationDetail<String> stringVariationDetail(String key, String defaultValue) {
            return null;
        }
        void setOffline() {}
    }

    // Ambient names the doc fragments assume an Activity host
    // provides.
    @SuppressWarnings("unused")
    private static final LDContext context = null;
    @SuppressWarnings("unused")
    private static final int secondsToBlock = 0;

    private Object getApplication() { return null; }

    public static void main(String[] args) {
        System.out.println("feature flag evaluates to true");
    }

    @SuppressWarnings({"unused", "ConstantConditions"})
    private void wrappee() {
        try {
{{ body }}
        } catch (Throwable ignored) { }
    }
}
```
