---
id: ios-client-sdk/scaffolds/swift-syntax-only
sdk: ios-client-sdk
kind: scaffold
lang: swift
file: AppDelegate.swift
description: |
  Parse-only validator for iOS client SDK doc fragments.

  Outputs `Sources/AppDelegate.swift` (the path the ios-client
  validator harness expects) plus a companion `Sources/ViewController.swift`
  so xcodebuild has a complete project shape. The wrappee body lives
  inside `_wrappee()`, an instance method on AppDelegate that's never
  invoked at runtime — application(didFinishLaunchingWithOptions:)
  returns true immediately and the XCTest case below asserts the
  EXAM-HELLO success line is printed by the test itself, not by the
  body.

  The harness's existing import-lift Python pre-step will also
  hoist any `import Foundation` / `import LaunchDarkly` lines that
  came along in the body up to file scope.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: ios-client
  entrypoint: AppDelegate.swift
  companions:
    - ios-client-sdk/scaffolds/swift-syntax-only-viewcontroller
---

```swift
import UIKit
import Foundation
import LaunchDarkly

// File-scope stubs so wrappee bodies that reference caller-supplied
// helpers (e.g. the experimentation onboarding `applyVariant(_:)`)
// type-check. Never invoked.
func applyVariant(_ variant: String) {}

@UIApplicationMain
class AppDelegate: UIResponder, UIApplicationDelegate {
    var window: UIWindow?

    func application(_ application: UIApplication,
                     didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?) -> Bool {
        // The XCTest case asserts the EXAM-HELLO success line is
        // printed; the wrappee body is never invoked at runtime.
        print("feature flag evaluates to true")
        return true
    }

    // Stub so the wrappee body's `client.boolVariation(...)` etc.
    // resolve at compile time. Never invoked.
    var client: LDClient! = nil

    // Wrappee body — references to client/context here resolve
    // through the stubs above; xcodebuild type-checks but doesn't
    // run.
    @objc func _wrappee() {
{{ body }}
    }
}
```
