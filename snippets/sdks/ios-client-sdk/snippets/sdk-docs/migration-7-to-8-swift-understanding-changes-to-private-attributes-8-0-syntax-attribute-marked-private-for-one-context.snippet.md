---
id: ios-client-sdk/sdk-docs/migration-7-to-8-swift-understanding-changes-to-private-attributes-8-0-syntax-attribute-marked-private-for-one-context
sdk: ios-client-sdk
kind: reference
lang: swift
description: "8.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```swift
var builder = LDContextBuilder(key: "example-context-key")
builder.addPrivateAttribute(Reference("email"))
builder.addPrivateAttribute(Reference("address"))

let context = try builder.build().get()
```
