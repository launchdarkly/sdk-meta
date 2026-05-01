---
id: ios-client-sdk/sdk-docs/evaluate-a-flag-swift-2
sdk: ios-client-sdk
kind: reference
lang: swift
description: "Swift in section \"Evaluate a flag\""
---

```swift
let showFeature = client.boolVariation(forKey: "example-flag-key", defaultValue: false)

if showFeature {
  // Application code to show the feature
else {
  // The code to run if the feature is off
}
```
