---
id: ios-client-sdk/sdk-docs/migration-4-to-5-evaluating-flags-4-x-syntax-swift-2
sdk: ios-client-sdk
kind: reference
lang: swift
description: "4.x syntax (Swift) in section \"Evaluating flags\""
---

```swift
let flagValueDetail: EvaluationDetail<Bool> = LDClient.shared.variationDetail(forKey: "example-flag-key", fallback: false)
```
