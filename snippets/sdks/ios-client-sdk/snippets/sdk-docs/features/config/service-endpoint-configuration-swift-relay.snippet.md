---
id: ios-client-sdk/sdk-docs/features/config/service-endpoint-configuration-swift-relay
sdk: ios-client-sdk
kind: reference
lang: swift
description: Service endpoint configuration example for iOS.
---

```swift
var ldConfig = LDConfig(mobileKey: "example-mobile-key", autoEnvAttributes: .enabled)
ldConfig.streamUrl = URL(string: "https://your-relay-proxy.com:8030")
ldConfig.baseUrl = URL(string: "https://your-relay-proxy.com:8030")
ldConfig.eventsUrl = URL(string: "https://your-relay-proxy.com:8030")
```
