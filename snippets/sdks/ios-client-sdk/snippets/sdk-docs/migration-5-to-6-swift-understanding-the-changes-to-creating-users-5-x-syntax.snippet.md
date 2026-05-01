---
id: ios-client-sdk/sdk-docs/migration-5-to-6-swift-understanding-the-changes-to-creating-users-5-x-syntax
sdk: ios-client-sdk
kind: reference
lang: swift
description: "5.x syntax in section \"Understanding the changes to creating users\""
---

```swift
let privateAttributes: [String] = ["name", "jobFunction"]
let customAttributes: [String: Any] = ["jobFunction": ["doctor"]]
let user = LDUser(key: "example-user-key", custom: customAttributes, privateAttributes: privateAttributes)
```
