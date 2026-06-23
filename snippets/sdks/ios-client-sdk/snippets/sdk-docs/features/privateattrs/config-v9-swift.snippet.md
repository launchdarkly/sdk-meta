---
id: ios-client-sdk/sdk-docs/features/privateattrs/config-v9-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Private attribute configuration for iOS SDK v9.0 (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
// All attributes marked private
config = LDConfig(mobileKey: "example-mobile-key", autoEnvAttributes: .enabled)
config.allContextAttributesPrivate = true

// Two attributes marked private
config = LDConfig(mobileKey: "example-mobile-key", autoEnvAttributes: .enabled)
config.privateContextAttributes = [Reference("email"), Reference("address")]
```
