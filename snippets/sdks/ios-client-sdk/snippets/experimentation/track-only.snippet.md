---
id: ios-client-sdk/experimentation/track-only
sdk: ios-client-sdk
kind: reference
lang: swift
description: Experimentation onboarding (track only) for ios-client-sdk — start the client from a setup function and add a trackMetric helper.
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
import LaunchDarkly

// Call this once from your app startup (e.g. application(_:didFinishLaunchingWithOptions:)).
// You can't invoke LDClient.start() at the top level in Swift — wrap it in a function.
func startLaunchDarkly() {
    // This is your mobile key.
    let config = LDConfig(mobileKey: "YOUR_MOBILE_KEY", autoEnvAttributes: .enabled)

    // A "context" is a data object representing users, devices, organizations, and other entities.
    let contextBuilder = LDContextBuilder(key: "EXAMPLE_CONTEXT_KEY")
    guard case .success(let context) = contextBuilder.build() else { return }

    LDClient.start(config: config, context: context, startWaitSeconds: 5) { timedOut in
        if timedOut {
            print("SDK didn't initialize in 5 seconds. SDK is still running and trying to get latest flags.")
        } else {
            print("SDK successfully initialized with the latest flags")
        }
    }
}

// Call trackMetric when a metric action occurs in your app —
// a tap, a form submit, a screen view, a custom event, whatever your metric measures.
func trackMetric(metricKey: String, data: LDValue = .null) {
    let client = LDClient.get()!
    client.track(key: metricKey, data: data)
}
```
