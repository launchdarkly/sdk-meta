---
id: ios-client-sdk/sdk-docs/migration-5-to-6-swift-changes-to-variationdetail-functions-5-x-syntax
sdk: ios-client-sdk
kind: reference
lang: swift
description: "5.x syntax in section \"Changes to variationDetail functions\""
---

```swift
let boolValue: LDEvaluationDetail<Bool> = LDClient.get()!.variationDetail(forKey: "boolFlag", defaultValue: false)
let arrayValue: LDEvaluationDetail<[Any]> = LDClient.get()!.variationDetail(forKey: "arrayFlag", defaultValue: ["abc", "def"] as [Any])
let arrayReason: [String: Any]? = arrayValue.reason
```
