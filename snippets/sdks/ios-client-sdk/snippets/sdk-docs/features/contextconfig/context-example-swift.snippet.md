---
id: ios-client-sdk/sdk-docs/features/contextconfig/context-example-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Context example for iOS SDK v8.0+ (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
var contextBuilder = LDContextBuilder(key: "example-user-key")
contextBuilder.trySetValue("name", .string("Sandy"))
contextBuilder.trySetValue("email", .string("sandy@example.com"))

let context = try? contextBuilder.build().get()
```
