---
id: ios-client-sdk/sdk-docs/initialize-the-client-ios-sdk-v8-x-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "iOS SDK v8.x (Swift) in section \"Initialize the client\""
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
  let config = LDConfig(mobileKey: "example-mobile-key")
  
  // You'll need this context later, but you can ignore it for now.
  let context = LDContextBuilder(key: "example-context-key")

  LDClient.start(config: config, context: context, startWaitSeconds: 5) { timedOut in
      if timedOut {
          // Client may not have the most recent flags for the configured context
      } else {
          // Client has received flags for the configured context
      }
  }
```
