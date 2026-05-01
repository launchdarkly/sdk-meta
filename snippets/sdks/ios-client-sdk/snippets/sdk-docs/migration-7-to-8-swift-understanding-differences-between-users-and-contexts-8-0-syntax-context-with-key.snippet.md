---
id: ios-client-sdk/sdk-docs/migration-7-to-8-swift-understanding-differences-between-users-and-contexts-8-0-syntax-context-with-key
sdk: ios-client-sdk
kind: reference
lang: swift
description: "8.0 syntax, context with key in section \"Understanding differences between users and contexts\""
---

```swift
var contextBuilder = LDContextBuilder(key: "example-context-key")
let context = try contextBuilder.build().get()
```
