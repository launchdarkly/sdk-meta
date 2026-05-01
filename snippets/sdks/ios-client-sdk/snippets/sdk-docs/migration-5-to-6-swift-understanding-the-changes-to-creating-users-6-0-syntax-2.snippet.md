---
id: ios-client-sdk/sdk-docs/migration-5-to-6-swift-understanding-the-changes-to-creating-users-6-0-syntax-2
sdk: ios-client-sdk
kind: reference
lang: swift
description: "6.0 syntax in section \"Understanding the changes to creating users\""
---

```swift
let user1 = LDUser(key: "example-user-key", name: "Sandy Smith")
let user2 = LDUser(key: "example-user-key", name: "Jesse Smith")
// Results in false
user1 == user2
```
