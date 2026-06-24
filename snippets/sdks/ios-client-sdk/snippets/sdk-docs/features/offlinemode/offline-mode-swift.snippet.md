---
id: ios-client-sdk/sdk-docs/features/offlinemode/offline-mode-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Offline mode example for iOS (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
LDClient.get()!.setOnline(false)
LDClient.get()!.setOnline(true) {
    // Client is online
}
let connectionStatus = LDClient.get()!.isOnline
```
