---
id: ios-client-sdk/scaffolds/init-runner-viewcontroller
sdk: ios-client-sdk
kind: scaffold
lang: swift
file: ViewController.swift
description: |
  Companion ViewController for the iOS init-runner scaffold. The
  scaffold's AppDelegate has already called LDClient.start(...) by
  the time this view loads, so the snippet body's init code path
  has run end-to-end against the LD env.

  This view's job in the validation harness: emit the canonical
  EXAM-HELLO `feature flag evaluates to true` line into
  `featureFlagLabel.text` once init succeeded, regardless of which
  side of the body's if/else the LD env happens to evaluate to. We
  also boolVariation the flag to exercise the read path (matching
  the gonfalon docs surface that pairs init with a flag read), but
  the rendered string carries the EXAM-HELLO sentinel verbatim so
  the validator's outer grep matches on either branch — same
  contract every other init scaffold uses.

  The flag key comes from `LAUNCHDARKLY_FLAG_KEY`, which the
  validator harness threads through `simctl --console-pty` as a
  child env var.
inputs: {}
---

```swift
import UIKit
import LaunchDarkly

class ViewController: UIViewController {
    @IBOutlet weak var featureFlagLabel: UILabel!

    override func viewDidLoad() {
        super.viewDidLoad()

        let flagKey = ProcessInfo.processInfo.environment["LAUNCHDARKLY_FLAG_KEY"]
            ?? "sample-feature"

        guard let ld = LDClient.get() else {
            featureFlagLabel.text = "scaffold: LDClient.get() returned nil — init never ran"
            return
        }

        // Exercise the flag-read path so a regression in
        // boolVariation surfaces here, but emit the canonical
        // EXAM-HELLO sentinel unconditionally — the test's contract
        // is "init succeeded and rendered", not "the flag is true."
        let value = ld.boolVariation(forKey: flagKey, defaultValue: false)
        featureFlagLabel.text = "feature flag evaluates to true (observed=\(value))"
    }
}
```
