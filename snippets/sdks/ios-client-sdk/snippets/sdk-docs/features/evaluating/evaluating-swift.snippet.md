---
id: ios-client-sdk/sdk-docs/features/evaluating/evaluating-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Flag evaluation example for iOS.
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
let boolFlagValue = LDClient.get()!.boolVariation(forKey: "example-bool-flag-key", defaultValue: false)
let jsonFlagValue = LDClient.get()!.jsonVariation(forKey: "json-flag-key-456def", defaultValue: ["enabled": false, "special": "none"])
```
