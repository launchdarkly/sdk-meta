---
id: ios-client-sdk/sdk-docs/background-fetch-ios-sdk-v8-0-and-earlier-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "iOS SDK v8.0 and earlier (Swift) in section \"Background fetch\""
# TODO(snippet-bug): body uses iOS SDK v8 (and earlier) LDConfig API
# (LDConfig(mobileKey:) with implicit autoEnvAttributes). In v9+ the
# autoEnvAttributes parameter is required. swift-syntax-only compiles
# against the latest launchdarkly-ios-client-sdk, so this v8-shape
# call fails. Fix in the follow-up snippet-bugs PR: either update to
# current v9 API and drop the v8-pinned variant, or pin a v8 SDK
# in a parallel scaffold if back-compat docs must stay live.
---

```swift
var ldConfig = LDConfig(mobileKey: "example-mobile-key")
ldConfig.backgroundFlagPollingInterval = 3600
```
