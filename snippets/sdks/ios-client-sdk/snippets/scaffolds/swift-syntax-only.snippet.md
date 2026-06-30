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
  # Syntax-only fragments are compiled, not run: the batch ios-client
  # harness builds them for the simulator SDK (no boot) rather than
  # running an `xcodebuild test`. init-runner snippets carry no
  # SNIPPET_CHECK and default to the runtime (simulator) path.
  env:
    SNIPPET_CHECK: parse
---

```swift
import UIKit
import Foundation
import LaunchDarkly
// The validator's project.yml depends on LaunchDarklyObservability;
// importing it here brings `Observability`, `LDOptions`, the
// `.enabled`/`.disabled` log/trace/metric enums into scope so doc
// fragments that show the optional plugin compile.
import LaunchDarklyObservability

// File-scope stubs so wrappee bodies that reference caller-supplied
// helpers (e.g. the experimentation onboarding `applyVariant(_:)`)
// type-check. Never invoked.
func applyVariant(_ variant: String) {}

// Stub of the legacy alias API (removed at v8) so the v7-era
// aliasing fragment type-checks against the current SDK. The
// ambient `newUser` / `previousUser` names are typed LDContext
// because the stub only needs self-consistent opaque arguments.
// Never invoked.
extension LDClient {
    func alias(context: LDContext, previousContext: LDContext) {}
}

// v8-era convenience surface: v9 made the `autoEnvAttributes:`
// constructor argument mandatory, so v8.x doc fragments that say
// `LDConfig(mobileKey:)` would not compile against the current SDK
// without this shim. Never invoked at runtime.
extension LDConfig {
    init(mobileKey: String) {
        self.init(mobileKey: mobileKey, autoEnvAttributes: .disabled)
    }
}

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

    // Stubs for fragments that mutate an ambient config or pass an
    // ambient context to `LDClient.start` — the docs assume earlier
    // init snippets created them.
    var ldConfig = LDConfig(mobileKey: "stub-mobile-key", autoEnvAttributes: .disabled)
    // Some config fragments assign to a bare `config` the docs assume
    // an earlier snippet declared.
    var config = LDConfig(mobileKey: "stub-mobile-key", autoEnvAttributes: .disabled)
    var context = try! LDContextBuilder(key: "stub-context-key").build().get()

    // Multi-environment fragments call methods on an ambient
    // `coreInstance` (a secondary-environment client an earlier
    // snippet fetched via `LDClient.get(environment:)`) and pass an
    // ambient `data` payload to `track(key:data:)`.
    var coreInstance: LDClient! = nil
    var data: LDValue = "stub-data"

    // Ambient names the legacy aliasing fragment assumes earlier
    // snippets created.
    var newUser = try! LDContextBuilder(key: "stub-new-user-key").build().get()
    var previousUser = try! LDContextBuilder(key: "stub-previous-user-key").build().get()

    // Wrappee body — references to client/context here resolve
    // through the stubs above; xcodebuild type-checks but doesn't
    // run. Marked `throws` so fragments that use bare `try`
    // (e.g. `try LDContextBuilder(...).build().get()`) compile
    // without per-fragment error handling.
    @objc func _wrappee() throws {
{{ body }}
    }
}
```
