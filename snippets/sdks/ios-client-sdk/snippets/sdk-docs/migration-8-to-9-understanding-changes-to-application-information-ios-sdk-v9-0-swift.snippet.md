---
id: ios-client-sdk/sdk-docs/migration-8-to-9-understanding-changes-to-application-information-ios-sdk-v9-0-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "iOS SDK v9.0 (Swift) in section \"Understanding changes to application information\""
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
