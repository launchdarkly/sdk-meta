---
id: ios-client-sdk/sdk-docs/migration-4-to-5-updating-the-active-user-context-5-0-syntax-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "5.0 syntax (Swift) in section \"Updating the active user context\""
---

```swift
let newUser = LDUser(key: "example-user-key")
LDClient.get()!.identify(user: newUser)
// identify can also be called with a completion
LDClient.get()!.identify(user: newUser) {
    // Flags have been received for the new user
}
```
