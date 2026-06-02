---
id: ios-client-sdk/sdk-docs/features/config/app-config-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Application metadata configuration example for iOS.
---

```swift
var applicationInfo = ApplicationInfo()
applicationInfo.applicationIdentifier("authentication-service")
applicationInfo.applicationName("Authentication-Service")
applicationInfo.applicationVersion("1.0.0")
applicationInfo.applicationVersionName("v1")

var config = LDConfig(mobileKey: mobileKey, autoEnvAttributes: .enabled)
config.applicationInfo = applicationInfo
```
