---
id: ios-client-sdk/sdk-docs/features/anonymous/anonymous-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Anonymous context example for iOS (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
var contextBuilder = LDContextBuilder(key: "example-context-key")
contextBuilder.anonymous(true)

let context = try contextBuilder.build().get()
```
