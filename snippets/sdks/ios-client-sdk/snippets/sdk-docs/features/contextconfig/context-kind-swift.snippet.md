---
id: ios-client-sdk/sdk-docs/features/contextconfig/context-kind-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Context with a non-user kind for iOS SDK v8.0+ (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
var contextBuilder = LDContextBuilder(key: "example-organization-key")
contextBuilder.kind("organization")

let context = try? contextBuilder.build().get()
```
