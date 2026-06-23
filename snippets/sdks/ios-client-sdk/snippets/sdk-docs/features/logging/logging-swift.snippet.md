---
id: ios-client-sdk/sdk-docs/features/logging/logging-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Logger configuration example for iOS SDK v9.x (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only

---

```swift
import OSLog

var config = LDConfig(mobileKey: "example-mobile-key", autoEnvAttributes: .enabled)
config.logger = OSLog(subsystem: "your.preferred.subsystem", category: "ld-sdk")

// You can disable all SDK logging by setting this property to the shared disabled logger
config.logger = .disabled
```
