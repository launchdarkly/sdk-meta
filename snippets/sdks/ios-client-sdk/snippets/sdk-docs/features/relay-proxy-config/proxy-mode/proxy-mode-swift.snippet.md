---
id: ios-client-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Proxy mode configuration example for iOS (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
var ldConfig = LDConfig(mobileKey: "example-mobile-key", autoEnvAttributes: .enabled)
ldConfig.streamUrl = URL(string: "https://your-relay-proxy.com:8030")!
ldConfig.baseUrl = URL(string: "https://your-relay-proxy.com:8030")!
ldConfig.eventsUrl = URL(string: "https://your-relay-proxy.com:8030")!
```
