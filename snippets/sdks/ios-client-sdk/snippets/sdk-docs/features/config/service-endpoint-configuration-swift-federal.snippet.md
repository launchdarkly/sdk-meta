---
id: ios-client-sdk/sdk-docs/features/config/service-endpoint-configuration-swift-federal
sdk: ios-client-sdk
kind: reference
lang: swift
description: Service endpoint configuration example for iOS.
---

```swift
var ldConfig = LDConfig(mobileKey: "example-mobile-key", autoEnvAttributes: .enabled)
ldConfig.streamUrl = URL(string: "https://clientstream.launchdarkly.us")
ldConfig.baseUrl = URL(string: "https://clientsdk.launchdarkly.us")
ldConfig.eventsUrl = URL(string: "https://events.launchdarkly.us")
```
