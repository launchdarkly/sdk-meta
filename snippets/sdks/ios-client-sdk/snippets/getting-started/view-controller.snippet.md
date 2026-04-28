---
id: ios-client-sdk/getting-started/view-controller
sdk: ios-client-sdk
kind: hello-world
lang: swift
file: ViewController.swift
description: ViewController that observes the flag and renders the value.
inputs:
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source.
ld-application:
  slot: view-controller
# Validator pending — same as app-delegate.
#
# Known bug in this snippet (carried verbatim from gonfalon for now):
# the updateUi label uses `(flagKey)` and `(result)` instead of Swift's
# string interpolation `\(flagKey)` and `\(result)`. As written it
# would print the literal text `(flagKey)`. Fix-on-red when the iOS
# validator lands.
---

Open `ViewController.swift` and add the following code:

```swift
import UIKit
import LaunchDarkly

class ViewController: UIViewController {

    @IBOutlet weak var featureFlagLabel: UILabel!

    // Set featureFlagKey to the feature flag key you want to evaluate.
    fileprivate let featureFlagKey = "{{ featureKey }}"

    override func viewDidLoad() {
        super.viewDidLoad()

        if let ld = LDClient.get() {
            ld.observe(key: featureFlagKey, owner: self) { [weak self] changedFlag in
                guard let me = self else { return }
                guard case .bool(let booleanValue) = changedFlag.newValue else { return }

                me.updateUi(flagKey: changedFlag.key, result: booleanValue)
            }
            let result = ld.boolVariation(forKey: featureFlagKey, defaultValue: false)
            updateUi(flagKey: featureFlagKey, result: result)
        }
    }

    func updateUi(flagKey: String, result: Bool) {
        self.featureFlagLabel.text = "The (flagKey) feature flag evaluates to (result)"

        let toggleOn = UIColor(red: 0, green: 0.52, blue: 0.29, alpha: 1)
        let toggleOff = UIColor(red: 0.22, green: 0.22, blue: 0.25, alpha: 1)
        self.view.backgroundColor = result ? toggleOn : toggleOff
    }
}
```
