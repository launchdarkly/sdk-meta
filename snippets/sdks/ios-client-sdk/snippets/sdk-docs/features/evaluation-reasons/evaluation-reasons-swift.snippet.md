---
id: ios-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Flag evaluation reason example for iOS (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
ldConfig.evaluationReasons = true
LDClient.start(config: ldConfig, context: context)

let detail = client.boolVariationDetail(forKey: "example-flag-key", defaultValue: false);

let value: Bool = detail.value
let variationIndex: Int? = detail.variationIndex
let reason: [String: LDValue]? = detail.reason
```
