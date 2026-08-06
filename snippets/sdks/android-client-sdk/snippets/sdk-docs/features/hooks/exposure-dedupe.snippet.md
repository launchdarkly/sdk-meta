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
        // Observes an evaluation at most once per minute per unique result,
        // tracking at most 2000 results at a time
        .addHook(new ExampleHook("Observability hook")
            .evaluationExposureDeduper(60_000, 2_000))
        // Observes every evaluation, stated explicitly
        .addHook(new ExampleHook("Audit hook")
            .evaluationExposureDeduper(EvaluationExposureDeduper.disabled()))
    )
    .build();
```
