---
id: ios-client-sdk/experimentation/button-copy
sdk: ios-client-sdk
kind: reference
lang: swift
description: Configures a UIButton to display the assigned variation's label and track taps. Pass your existing button instance to wire it up in place.
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
import LaunchDarkly
import UIKit

// Configures a UIButton to display the assigned variation's title and track taps.
// Pass your existing button instance to wire it up without changing your layout.
// Call this after startLaunchDarkly() and after identify() resolves (if the user
// became known mid-session).
//
// Prerequisites:
//   - A string flag whose key matches YOUR_FLAG_KEY. Set each variation's value to
//     the button label you want users to see (e.g. "Get started", "Start for free").
//     The flag value is used as the button title directly.
//   - A tap metric whose key matches YOUR_METRIC_KEY attached to your experiment.
// Requires iOS 14+. For iOS 13 support, replace UIAction with addTarget(_:action:for:).
@available(iOS 14.0, *)
func configureExperimentButton(_ button: UIButton, onTap: (() -> Void)? = nil) {
    let client = LDClient.get()!

    // The flag value is the button title. The default is shown when the flag is off
    // or the SDK hasn't finished initializing yet.
    // Don't cache the result — LaunchDarkly deduplicates exposure events automatically.
    let label = client.stringVariation(forKey: "YOUR_FLAG_KEY", defaultValue: "Get started")
    button.setTitle(label, for: .normal)

    // Use a stable identifier so re-calling this function (e.g. after identify())
    // replaces the existing handler rather than stacking a second one.
    let actionID = UIAction.Identifier("com.example.experimentButton")
    button.removeAction(identifiedBy: actionID, for: .touchUpInside)

    let action = UIAction(identifier: actionID) { _ in
        // Track the tap so LaunchDarkly can attribute it to the right variation.
        // Use the same context that was active during the flag evaluation above —
        // mismatched contexts break conversion attribution.
        client.track(key: "YOUR_METRIC_KEY")
        onTap?()
    }
    button.addAction(action, for: .touchUpInside)
}
```
