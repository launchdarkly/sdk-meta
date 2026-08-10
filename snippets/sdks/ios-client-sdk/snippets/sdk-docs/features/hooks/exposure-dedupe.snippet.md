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

class MetricsHook: Hook {}
class ObservabilityHook: Hook {}
class TelemetryHook: Hook {}
class ExperimentHook: Hook {}

var config = LDConfig(
  mobileKey: "example-mobile-key",
  autoEnvAttributes: .enabled
)
config.hooks = [
  /// Observes every evaluation, because it is registered unwrapped
  MetricsHook(),
  /// Observes a flag when its result changes, and at most once per 10
  /// minutes while it stays the same: the default window
  DedupingHook(ObservabilityHook()),
  /// The same, with a one minute window
  DedupingHook(TelemetryHook(), window: 60),
  /// Deciding for itself which evaluations reach the hook
  DedupingHook(ExperimentHook(), deduper: EvaluationExposureDeduper(window: 30))
]
```
