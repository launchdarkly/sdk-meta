---
id: ios-client-sdk/sdk-docs/features/config/index-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: SDK configuration example for iOS.

---

```swift
var ldConfig = LDConfig(mobileKey: "example-mobile-key", autoEnvAttributes: .enabled)
ldConfig.connectionTimeout = 10.0
ldConfig.eventFlushInterval = 30.0
```
