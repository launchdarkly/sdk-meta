---
id: ios-client-sdk/sdk-docs/migration-7-to-8-swift-working-with-built-in-and-custom-attributes-8-0-syntax-context-with-attributes
sdk: ios-client-sdk
kind: reference
lang: swift
description: "8.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```swift
var builder = LDContextBuilder(key: "example-context-key")
builder.kind("user")
builder.name("Sandy Smith")
builder.trySetValue("email", "sandy@example.com")

let context = try builder.build().get()
```
