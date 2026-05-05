---
id: ios-client-sdk/sdk-docs/background-fetch-ios-sdk-v9-0-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "iOS SDK v9.0 (Swift) in section \"Background fetch\""
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
var ldConfig = LDConfig(mobileKey: "example-mobile-key", autoEnvAttributes: .enabled)
ldConfig.backgroundFlagPollingInterval = 3600
```
