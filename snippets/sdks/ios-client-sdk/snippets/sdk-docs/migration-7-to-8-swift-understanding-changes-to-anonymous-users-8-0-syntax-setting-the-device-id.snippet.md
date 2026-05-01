---
id: ios-client-sdk/sdk-docs/migration-7-to-8-swift-understanding-changes-to-anonymous-users-8-0-syntax-setting-the-device-id
sdk: ios-client-sdk
kind: reference
lang: swift
description: "8.0 syntax, setting the device ID in section \"Understanding changes to anonymous users\""
---

```swift
var builder = LDContextBuilder(key: "example-context-key")
builder.kind("device")
builder.trySetValue("deviceID", LDValue(stringLiteral: UIDevice.current.identifierForVendor?.uuidString ?? ""))

let context = try builder.build().get()
```
