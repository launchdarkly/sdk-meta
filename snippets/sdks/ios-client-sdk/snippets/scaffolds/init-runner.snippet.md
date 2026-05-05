---
id: ios-client-sdk/scaffolds/init-runner
sdk: ios-client-sdk
kind: scaffold
lang: swift
file: AppDelegate.swift
description: |
  Runs an `init.txt`-style iOS Swift snippet end-to-end against a real
  LaunchDarkly mobile env. The body is a `LDClient.start(...)` call
  that needs an `Application` lifecycle to host it; we splice it into
  `application(_:didFinishLaunchingWithOptions:)` of an `AppDelegate`,
  link the `LaunchDarkly` Swift Package via the validator scaffold's
  `project.yml`, and let the existing `ios-client` validator drive the
  app on the iOS Simulator via `xcodebuild test`.

  Layout:
    - AppDelegate.swift (this scaffold) — splices the snippet body
      into didFinishLaunching. The body's `guard case .success(...)
      else { return }` returns from the function, abandoning init —
      acceptable for a simulator harness, since the test asserts
      LDClient.get() resolved to a real client. We use
      validation.placeholders to substitute the snippet's literal
      `YOUR_MOBILE_KEY` for the env-injected mobile key.
    - ViewController.swift (companion) — observes the test flag and
      writes the canonical EXAM-HELLO `feature flag evaluates to
      true` line into `featureFlagLabel.text` so the Tests target's
      XCTest sees it. The companion's flag-key reference comes from
      `LAUNCHDARKLY_FLAG_KEY` via the standard test environment, not
      from the snippet body itself; the snippet body is init-only.

  Why this is the right wrapper: gonfalon's docs frame this body as
  "drop into your AppDelegate" — exactly where it lands here. The
  body's `return` statement is now in a function, the
  `LDClient.start` completion fires asynchronously, and the
  validator's XCTest waits up to 30s for the flag to evaluate. All
  three of those align with the snippet's natural surface.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: ios-client
  entrypoint: AppDelegate.swift
  companions:
    - ios-client-sdk/scaffolds/init-runner-viewcontroller
---

```swift
import UIKit
import LaunchDarkly

@UIApplicationMain
class AppDelegate: UIResponder, UIApplicationDelegate {
    var window: UIWindow?

    func application(_ application: UIApplication,
                     didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?) -> Bool {
        setUpLDClient()
        return true
    }

    func setUpLDClient() {
{{ body }}
    }
}
```
