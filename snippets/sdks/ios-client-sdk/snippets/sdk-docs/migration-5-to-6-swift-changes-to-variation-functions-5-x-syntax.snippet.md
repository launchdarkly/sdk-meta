---
id: ios-client-sdk/sdk-docs/migration-5-to-6-swift-changes-to-variation-functions-5-x-syntax
sdk: ios-client-sdk
kind: reference
lang: swift
description: "5.x syntax in section \"Changes to variation functions\""
---

```swift
let boolValue: Bool = LDClient.get()!.variation(forKey: "boolFlag", defaultValue: false)
let arrayValue: [Any] = LDClient.get()!.variation(forKey: "arrayFlag", defaultValue: ["abc", "def"] as [Any])
```
