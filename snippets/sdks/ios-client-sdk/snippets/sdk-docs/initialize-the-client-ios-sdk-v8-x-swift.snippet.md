---
id: ios-client-sdk/sdk-docs/initialize-the-client-ios-sdk-v8-x-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "iOS SDK v8.x (Swift) in section \"Initialize the client\""
# TODO(snippet-bug): body uses iOS SDK v8 (and earlier) LDConfig API
# (LDConfig(mobileKey:) with implicit autoEnvAttributes). In v9+ the
# autoEnvAttributes parameter is required. swift-syntax-only compiles
# against the latest launchdarkly-ios-client-sdk, so this v8-shape
# call fails. Fix in the follow-up snippet-bugs PR: either update to
# current v9 API and drop the v8-pinned variant, or pin a v8 SDK
# in a parallel scaffold if back-compat docs must stay live.
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
