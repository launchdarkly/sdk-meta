---
id: ios-client-sdk/sdk-docs/features/localstorage/localstorage-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Local storage caching example for iOS (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
var config = LDConfig(mobileKey: "example-mobile-key", autoEnvAttributes: .enabled)

// Local storage is enabled by default
// You can optionally configure the maximum number of cached contexts (default is 5)
config.maxCachedContexts = 3

let contextBuilder = LDContextBuilder(key: "example-context-key")
guard case .success(let context) = contextBuilder.build() else { return }

LDClient.start(config: config, context: context)
```
