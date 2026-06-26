---
id: ios-client-sdk/sdk-docs/features/privateattrs/context-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Marking context attributes private with the context builder for iOS SDK v8.0+ (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
var contextBuilder = LDContextBuilder(key: "example-context-key")
contextBuilder.trySetValue("name", .string("Sandy"))
contextBuilder.trySetValue("group", .array([LDValue(stringLiteral: "microsoft")]))
contextBuilder.addPrivateAttribute(Reference("name"))
contextBuilder.addPrivateAttribute(Reference("group"))

let context = try contextBuilder.build().get()
```
