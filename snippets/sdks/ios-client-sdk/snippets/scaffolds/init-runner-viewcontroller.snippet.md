---
id: ios-client-sdk/scaffolds/init-runner-viewcontroller
sdk: ios-client-sdk
kind: scaffold
lang: swift
file: ViewController.swift
description: |
  Companion ViewController for the iOS init-runner scaffold. The
  scaffold's AppDelegate has already called LDClient.start(...) by
  the time this view loads. We observe the test flag and write the
  canonical EXAM-HELLO `feature flag evaluates to true` line into
  `featureFlagLabel.text` once the flag evaluates true; the
  validator's XCTest case (validators/languages/ios-client/scaffold/
  Tests/SnippetTest.swift) drives the same code path and asserts on
  the label text.

  The flag key is injected via `LAUNCHDARKLY_FLAG_KEY` and threaded
  through `simctl --console-pty` as a child env var.
inputs: {}
---

```swift
import UIKit
import LaunchDarkly

class ViewController: UIViewController {
    @IBOutlet weak var featureFlagLabel: UILabel!

    override func viewDidLoad() {
        super.viewDidLoad()

        // The flag key is injected by the validator harness via
        // SIMCTL_CHILD_LAUNCHDARKLY_FLAG_KEY → LAUNCHDARKLY_FLAG_KEY
        // inside the simulator process.
        let flagKey = ProcessInfo.processInfo.environment["LAUNCHDARKLY_FLAG_KEY"]
            ?? "sample-feature"

        guard let ld = LDClient.get() else {
            featureFlagLabel.text = "scaffold: LDClient.get() returned nil — init never ran"
            return
        }

        let render: (Bool) -> Void = { [weak self] value in
            self?.featureFlagLabel.text = value
                ? "feature flag evaluates to true"
                : "scaffold: flag evaluated to false"
        }

        // Seed with the cached value (in case streaming has already
        // delivered the flag), then observe for changes.
        render(ld.boolVariation(forKey: flagKey, defaultValue: false))
        ld.observe(key: flagKey, owner: self) { changedFlag in
            if case .bool(let b) = changedFlag.newValue {
                render(b)
            }
        }
    }
}
```
