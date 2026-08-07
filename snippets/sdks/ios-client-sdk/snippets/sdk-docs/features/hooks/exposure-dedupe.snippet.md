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
    /// Observes a flag when its result changes, and at most once per 10
    /// minutes while it stays the same: the default window
    let evaluationExposureDeduper: EvaluationExposureDeduper? =
        EvaluationExposureDeduper()
}

class TelemetryHook: Hook {
    /// The same, with a one minute window
    let evaluationExposureDeduper: EvaluationExposureDeduper? =
        EvaluationExposureDeduper(window: 60)
}

class AuditHook: Hook {
    /// Observes every evaluation, stated explicitly
    let evaluationExposureDeduper: EvaluationExposureDeduper? = .disabled
}

var config = LDConfig(
  mobileKey: "example-mobile-key",
  autoEnvAttributes: .enabled
)
config.hooks = [MetricsHook(), ObservabilityHook(), TelemetryHook(), AuditHook()]
```
