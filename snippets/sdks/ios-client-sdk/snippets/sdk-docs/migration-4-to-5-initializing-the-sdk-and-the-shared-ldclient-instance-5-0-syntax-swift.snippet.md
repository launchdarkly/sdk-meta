---
id: ios-client-sdk/sdk-docs/migration-4-to-5-initializing-the-sdk-and-the-shared-ldclient-instance-5-0-syntax-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "5.0 syntax (Swift) in section \"Initializing the SDK and the shared LDClient instance\""
---

```swift
let config = LDConfig(mobileKey: "example-mobile-key")
let user = LDUser(key: "example-user-key")
LDClient.start(config: config, user: user)
// Safe to force get
let shared = LDClient.get()!
```
