---
id: ios-client-sdk/sdk-docs/features/hooks/exposure-dedupe
sdk: ios-client-sdk
kind: reference
lang: swift
description: Per-hook evaluation exposure deduplication for the iOS SDK (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
import LaunchDarkly

class MetricsHook: Hook {
    /// Observes every evaluation, because it does not declare a deduper
}

class ObservabilityHook: Hook {
    /// Observes an evaluation at most once per minute per unique result,
    /// tracking at most 2000 results at a time
    let evaluationExposureDeduper: EvaluationExposureDeduper? =
        EvaluationExposureDeduper(window: 60, maxSize: 2_000)
}

class AuditHook: Hook {
    /// Observes every evaluation, stated explicitly
    let evaluationExposureDeduper: EvaluationExposureDeduper? = .disabled
}

var config = LDConfig(
  mobileKey: "example-mobile-key",
  autoEnvAttributes: .enabled
)
config.hooks = [MetricsHook(), ObservabilityHook(), AuditHook()]
```
