---
id: ios-client-sdk/sdk-docs/features/multienv/init-v8-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Multi-environment configuration for iOS SDK v8.x (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
let context = try LDContextBuilder(key: "example-context-key").build().get()
var config = LDConfig(mobileKey: "example-mobile-key")
// The SDK throws error strings if you add duplicate keys or put the primary key or name in secondaryMobileKeys.
try! config.setSecondaryMobileKeys(["platform": "platform-example-mobile-key", "core": "core-example-mobile-key"])
LDClient.start(config: config, context: context)
```
