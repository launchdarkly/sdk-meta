---
id: ios-client-sdk/sdk-docs/features/config/service-endpoint-configuration-swift-eu
sdk: ios-client-sdk
kind: reference
lang: swift
description: Service endpoint configuration example for iOS.

---

```swift
var ldConfig = LDConfig(mobileKey: "example-mobile-key", autoEnvAttributes: .enabled)
ldConfig.streamUrl = URL(string: "https://clientstream.eu.launchdarkly.com")
ldConfig.baseUrl = URL(string: "https://clientsdk.eu.launchdarkly.com")
ldConfig.eventsUrl = URL(string: "https://events.eu.launchdarkly.com")
```
