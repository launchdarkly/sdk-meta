---
id: ios-client-sdk/sdk-docs/migration-5-to-6-swift-recording-custom-events-6-0-syntax
sdk: ios-client-sdk
kind: reference
lang: swift
description: "6.0 syntax in section \"Recording custom events\""
---

```swift
let customData: LDValue = ["abc": 123]
LDClient.get()!.track(key: "key", data: customData)
```
