---
id: ios-client-sdk/experimentation/full
sdk: ios-client-sdk
kind: reference
lang: swift
description: Full experimentation onboarding for ios-client-sdk — start from a setup function, identify on login/eligibility, evaluate, and track conversions.
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
import LaunchDarkly

// Call this once from your app startup (e.g. application(_:didFinishLaunchingWithOptions:)).
// You can't invoke LDClient.start() at the top level in Swift — wrap it in a function.
// Initialize once — extra clients can cause inconsistent experiment results.
func startLaunchDarkly() {
    let config = LDConfig(mobileKey: "YOUR_MOBILE_KEY", autoEnvAttributes: .enabled)

    // A "context" is a data object representing users, devices, organizations, and other entities.
    // If you already know the user's key at startup, initialize with it directly.
    // Use a consistent key so the same user gets the same experience.
    let contextBuilder = LDContextBuilder(key: "EXAMPLE_CONTEXT_KEY")
    guard case .success(let context) = contextBuilder.build() else { return }

    LDClient.start(config: config, context: context, startWaitSeconds: 5) { _ in }
}

// Call this when the user logs in or becomes eligible mid-session.
// Wait for identify to finish before evaluating experiment flags.
func onUserBecomesEligible(finalUserKey: String) {
    let client = LDClient.get()!

    // Update to the final context used for experiment eligibility.
    // Use the logged-in user's ID so experiment assignment stays consistent.
    let updated = try! LDContextBuilder(key: finalUserKey)
        // any attributes that affect targeting or eligibility
        .build().get()

    client.identify(context: updated) {
        // Evaluate the experiment flag where the user encounters the experience,
        // after identify completes.
        let variant = client.stringVariation(forKey: "YOUR_FLAG_KEY", defaultValue: "control")
        applyVariant(variant)
    }
}

// Call this when the user completes a metric action.
// Use the same user key you used when evaluating the flag — mismatched keys break conversion tracking.
// The data argument is optional and accepts any shape your metric needs.
func trackMetric(metricKey: String, data: LDValue = .null) {
    let client = LDClient.get()!
    client.track(key: metricKey, data: data)
}

// The SDK batches and flushes events automatically, including when the app is
// backgrounded. Don't add manual flush() calls — they're unnecessary and make
// real problems harder to spot.
// Don't skip or cache flag evaluations to reduce exposure counts — LaunchDarkly deduplicates them automatically.
```
