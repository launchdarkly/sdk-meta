---
id: ios-client-sdk/sdk-docs/features/multienv/get-environment-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Accessing a secondary environment client instance on iOS (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
let coreInstance = LDClient.get(environment: "core")
// Variation determines whether or not a flag is enabled for a specific context
let coreFlagValue = coreInstance?.boolVariation(forKey: "core-example-flag-key", defaultValue: false)
// allFlags produces a map of feature flag keys to their values
let allFlags: [String: LDValue]? = coreInstance?.allFlags
// track records actions end users take in your app
try coreInstance?.track(key: "track-example-event-key", data: data)
```
