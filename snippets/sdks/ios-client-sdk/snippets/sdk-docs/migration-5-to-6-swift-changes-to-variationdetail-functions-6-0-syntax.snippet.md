---
id: ios-client-sdk/sdk-docs/migration-5-to-6-swift-changes-to-variationdetail-functions-6-0-syntax
sdk: ios-client-sdk
kind: reference
lang: swift
description: "6.0 syntax in section \"Changes to variationDetail functions\""
---

```swift
let boolValue: LDEvaluationDetail<Bool> = LDClient.get()!.boolVariationDetail(forKey: "boolFlag", defaultValue: false)
let arrayValue: LDEvaluationDetail<LDValue> = LDClient.get()!.jsonVariationDetail(forKey: "arrayFlag", defaultValue: ["abc", "def"])
let arrayReason: [String: LDValue]? = arrayValue.reason
```
