---
id: ios-client-sdk/sdk-docs/features/allflags/allflags-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: All flags example for iOS (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
let allFlags: [String: LDValue]? = LDClient.get()!.allFlags
```
