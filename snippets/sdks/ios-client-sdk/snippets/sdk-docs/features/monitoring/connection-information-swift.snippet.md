---
id: ios-client-sdk/sdk-docs/features/monitoring/connection-information-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Connection information and connection mode observer for iOS (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
// Get current connection information
let connectionInformation = LDClient.get()!.getConnectionInformation()
// Setting a connection mode update observer
LDClient.get()!.observeCurrentConnectionMode(owner: self) { [weak self] connectionMode in
    // do something after ConnectionMode was updated.
}
```
