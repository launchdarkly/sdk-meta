---
id: ios-client-sdk/sdk-docs/features/contextconfig/multi-context-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Multi-context example for iOS SDK v8.0+ (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

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
