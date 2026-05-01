---
id: ios-client-sdk/sdk-docs/migration-7-to-8-swift-referencing-properties-of-an-attribute-object-8-0-syntax-context-with-object-attributes
sdk: ios-client-sdk
kind: reference
lang: swift
description: "8.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
---

```swift
var builder = LDContextBuilder(key: "example-context-key")
let address = LDValue(dictionaryLiteral: ("street", LDValue(stringLiteral: "Main St")), ("city", LDValue(stringLiteral: "Springfield")))
builder.trySetValue("address", address)

let context = try builder.build().get()
```
