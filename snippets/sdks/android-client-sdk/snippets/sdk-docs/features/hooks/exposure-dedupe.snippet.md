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
        // Observes every evaluation, because it has no deduper
        .addHook(new ExampleHook("Metrics hook"))
        // Observes a flag when its result changes, and at most once per 10
        // minutes while it stays the same, for up to 2000 flags: the defaults
        .addHook(new ExampleHook("Observability hook")
            .evaluationExposureDeduper())
        // The same, with a one minute window over at most 5000 flags
        .addHook(new ExampleHook("Telemetry hook")
            .evaluationExposureDeduper(60_000, 5_000))
        // Observes every evaluation, stated explicitly
        .addHook(new ExampleHook("Audit hook")
            .evaluationExposureDeduper(EvaluationExposureDeduper.disabled()))
    )
    .build();
```
