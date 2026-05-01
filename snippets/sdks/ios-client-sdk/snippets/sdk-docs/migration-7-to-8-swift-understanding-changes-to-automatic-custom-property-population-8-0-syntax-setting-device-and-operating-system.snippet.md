---
id: ios-client-sdk/sdk-docs/migration-7-to-8-swift-understanding-changes-to-automatic-custom-property-population-8-0-syntax-setting-device-and-operating-system
sdk: ios-client-sdk
kind: reference
lang: swift
description: "8.0 syntax, setting device and operating system in section \"Understanding changes to automatic custom property population\""
---

```swift
var builder = LDContextBuilder(key: "example-context-key")
builder.kind("device")
builder.trySetValue("os", "os")
builder.trySetValue("device", "device")

let context = try builder.build().get()
```
