---
id: ios-client-sdk/sdk-docs/migration-4-to-5-evaluating-flags-5-0-syntax-swift-2
sdk: ios-client-sdk
kind: reference
lang: swift
description: "5.0 syntax (Swift) in section \"Evaluating flags\""
---

```swift
let flagValueDetail: LDEvaluationDetail<Bool> = LDClient.get()!.variationDetail(forKey: "example-flag-key", defaultValue: false)
```
