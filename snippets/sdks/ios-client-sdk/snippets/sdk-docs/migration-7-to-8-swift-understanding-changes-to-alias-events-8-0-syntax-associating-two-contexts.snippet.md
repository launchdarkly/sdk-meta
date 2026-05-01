---
id: ios-client-sdk/sdk-docs/migration-7-to-8-swift-understanding-changes-to-alias-events-8-0-syntax-associating-two-contexts
sdk: ios-client-sdk
kind: reference
lang: swift
description: "8.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
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
