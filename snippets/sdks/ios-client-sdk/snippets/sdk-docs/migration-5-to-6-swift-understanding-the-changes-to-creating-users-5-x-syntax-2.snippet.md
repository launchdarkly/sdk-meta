---
id: ios-client-sdk/sdk-docs/migration-5-to-6-swift-understanding-the-changes-to-creating-users-5-x-syntax-2
sdk: ios-client-sdk
kind: reference
lang: swift
description: "5.x syntax in section \"Understanding the changes to creating users\""
---

```swift
let user1 = LDUser(key: "example-user-key", name: "Sandy Smith")
let user2 = LDUser(key: "example-user-key", name: "Jesse Smith")
// Results in true
user1 == user2
```
