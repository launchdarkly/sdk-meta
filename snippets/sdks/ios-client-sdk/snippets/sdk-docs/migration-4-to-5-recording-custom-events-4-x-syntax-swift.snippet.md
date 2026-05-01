---
id: ios-client-sdk/sdk-docs/migration-4-to-5-recording-custom-events-4-x-syntax-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "4.x syntax (Swift) in section \"Recording custom events\""
---

```swift
do {
    try LDClient.shared.trackEvent(key: "key", data: ["abc": 123])
} catch JSONSerialization.JSONError.invalidJsonObject {
    // Do something with the error
} catch {}
```
