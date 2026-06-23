---
id: ios-client-sdk/sdk-docs/features/anonymous/anonymous-autogen-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Auto-generated-key anonymous context example for iOS (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
// Have the SDK use a device persistent key.
// This sets `isAnonymous` by default.
let context = try LDContextBuilder().build().get()
```
