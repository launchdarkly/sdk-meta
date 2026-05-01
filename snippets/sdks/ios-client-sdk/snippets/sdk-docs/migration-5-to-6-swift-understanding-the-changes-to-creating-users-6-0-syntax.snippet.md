---
id: ios-client-sdk/sdk-docs/migration-5-to-6-swift-understanding-the-changes-to-creating-users-6-0-syntax
sdk: ios-client-sdk
kind: reference
lang: swift
description: "6.0 syntax in section \"Understanding the changes to creating users\""
---

```swift
let privateAttributes: [UserAttribute] = [UserAttribute.BuiltIn.name, UserAttribute.forName("jobFunction")]
let customAttributes: [String: LDValue] = ["jobFunction": ["doctor"]]
let user = LDUser(key: "example-user-key", custom: customAttributes, privateAttributes: privateAttributes)
```
