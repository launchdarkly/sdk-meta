---
id: ios-client-sdk/sdk-docs/background-fetch-ios-sdk-v8-0-and-earlier-objective-c
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "iOS SDK v8.0 and earlier (Objective-C) in section \"Background fetch\""
# TODO(validate): no Objective-C scaffold exists, and the body uses the v8 LDConfig initializer with no
# autoEnvAttributes:, which the pinned SDK 11.x does not have. Needs an Objective-C scaffold and a v8 pin.
---

```objectivec
LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key"];
config.backgroundFlagPollingInterval = 3600;
```
