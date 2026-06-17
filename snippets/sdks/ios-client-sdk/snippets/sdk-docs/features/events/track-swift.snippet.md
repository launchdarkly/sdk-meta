---
id: ios-client-sdk/sdk-docs/features/events/track-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Custom event tracking example for iOS (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
let data: LDValue = ["some-custom-key": "some-custom-value", "another-custom-key": 7]
try LDClient.get()!.track(key: "example-event-key", data: data)
```
