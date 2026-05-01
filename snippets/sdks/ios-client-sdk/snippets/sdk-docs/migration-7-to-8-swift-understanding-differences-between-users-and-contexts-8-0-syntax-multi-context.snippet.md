---
id: ios-client-sdk/sdk-docs/migration-7-to-8-swift-understanding-differences-between-users-and-contexts-8-0-syntax-multi-context
sdk: ios-client-sdk
kind: reference
lang: swift
description: "8.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
---

```swift
var userBuilder = LDContextBuilder(key: "example-user-key")
var deviceBuilder = LDContextBuilder(key: "example-device-key")
deviceBuilder.kind("device")

var multiBuilder = LDMultiContextBuilder()
multiBuilder.addContext(try userBuilder.build().get())
multiBuilder.addContext(try deviceBuilder.build().get())

let context = try multiBuilder.build().get()
```
