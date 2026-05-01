---
id: ios-client-sdk/sdk-docs/migration-5-to-6-swift-understanding-the-changes-to-configuration-6-0-syntax
sdk: ios-client-sdk
kind: reference
lang: swift
description: "6.0 syntax in section \"Understanding the changes to configuration\""
---

```swift
var config = LDConfig(mobileKey: "example-mobile-key")
config.privateUserAttributes = [UserAttribute.BuiltIn.name, UserAttribute.forName("premium")]
```
