---
id: ios-client-sdk/sdk-docs/migration-5-to-6-swift-changes-to-variation-functions-6-0-syntax
sdk: ios-client-sdk
kind: reference
lang: swift
description: "6.0 syntax in section \"Changes to variation functions\""
---

```swift
let boolValue = LDClient.get()!.boolVariation(forKey: "boolFlag", defaultValue: false)
let arrayValue = LDClient.get()!.jsonVariation(forKey: "arrayFlag", defaultValue: ["abc", "def"])
```
