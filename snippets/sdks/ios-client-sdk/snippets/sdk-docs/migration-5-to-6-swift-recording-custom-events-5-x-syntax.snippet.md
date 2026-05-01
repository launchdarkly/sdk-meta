---
id: ios-client-sdk/sdk-docs/migration-5-to-6-swift-recording-custom-events-5-x-syntax
sdk: ios-client-sdk
kind: reference
lang: swift
description: "5.x syntax in section \"Recording custom events\""
---

```swift
do {
    let customData: Any = ["abc": 123]
    try LDClient.get()!.track(key: "key", data: customData)
} catch let error as LDInvalidArgumentError {
    // Do something with the error
} catch {}
```
