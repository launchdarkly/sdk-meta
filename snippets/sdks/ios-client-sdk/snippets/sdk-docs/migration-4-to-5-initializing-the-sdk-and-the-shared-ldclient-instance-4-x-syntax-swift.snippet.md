---
id: ios-client-sdk/sdk-docs/migration-4-to-5-initializing-the-sdk-and-the-shared-ldclient-instance-4-x-syntax-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "4.x syntax (Swift) in section \"Initializing the SDK and the shared LDClient instance\""
---

```swift
let config = LDConfig(mobileKey: "example-mobile-key")
let user = LDUser(key: "example-user-key")
LDClient.shared.start(config: config, user: user)
```
