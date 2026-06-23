---
id: ios-client-sdk/sdk-docs/features/privateattrs/config-v8-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Private attribute configuration for iOS SDK v8.x (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
// All attributes marked private
var config = LDConfig(mobileKey: "example-mobile-key")
config.allContextAttributesPrivate = true

// Two attributes marked private
config = LDConfig(mobileKey: "example-mobile-key")
config.privateContextAttributes = [Reference("email"), Reference("address")]
```
