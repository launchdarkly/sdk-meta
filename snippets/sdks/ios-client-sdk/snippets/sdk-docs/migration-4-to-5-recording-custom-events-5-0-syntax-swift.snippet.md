---
id: ios-client-sdk/sdk-docs/migration-4-to-5-recording-custom-events-5-0-syntax-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "5.0 syntax (Swift) in section \"Recording custom events\""
---

```swift
do {
    try LDClient.get()!.track(key: "key", data: ["abc": 123])
} catch let error as LDInvalidArgumentError {
    // Do something with the error
} catch {}
```
