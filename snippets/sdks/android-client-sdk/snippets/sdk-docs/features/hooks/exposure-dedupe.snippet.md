---
id: android-client-sdk/sdk-docs/features/hooks/exposure-dedupe
sdk: android-client-sdk
kind: reference
lang: java
description: Per-hook evaluation exposure deduplication for the Android SDK.
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java
LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .hooks(
      Components.hooks()
        // Observes every evaluation, because it is registered unwrapped
        .addHook(new ExampleHook("Metrics hook"))
        // Observes a flag when its result changes, and at most once per 10
        // minutes while it stays the same: the default window
        .addHook(new DedupingHook(new ExampleHook("Observability hook")))
        // The same, with a one minute window
        .addHook(new DedupingHook(new ExampleHook("Telemetry hook"), 60_000))
        // Deciding for itself which evaluations reach the hook
        .addHook(new DedupingHook(new ExampleHook("Experiment hook"),
            new EvaluationExposureDeduper(30_000)))
    )
    .build();
```
