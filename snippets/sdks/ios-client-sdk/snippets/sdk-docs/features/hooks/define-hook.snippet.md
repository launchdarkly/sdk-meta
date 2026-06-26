---
id: ios-client-sdk/sdk-docs/features/hooks/define-hook
sdk: ios-client-sdk
kind: reference
lang: swift
description: Hook implementation and configuration for the iOS SDK (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
import LaunchDarkly

class ExampleHook: Hook {
    func metadata() -> Metadata {
        return Metadata(name: "example-hook")
    }

    /// Implement at least one of `beforeEvaluation`, `afterEvaluation`

    /// beforeEvaluation is called during the execution of a variation method
    /// before the flag value has been determined

    /// afterEvaluation is called during the execution of a variation method
    /// after the flag value has been determined
}

let exampleHook = ExampleHook()

var config = LDConfig(
  mobileKey: "example-mobile-key",
  autoEnvAttributes: .enabled
)
config.hooks = [exampleHook]

let context = try! LDContextBuilder(key: "example-context-key").build().get()

LDClient.start(config: config, context: context, startWaitSeconds: 5) { timedOut in
    if timedOut {
        /// Client may not have the most recent flags for the configured context
    } else {
        /// Client has received flags for the configured context
    }
}
```
