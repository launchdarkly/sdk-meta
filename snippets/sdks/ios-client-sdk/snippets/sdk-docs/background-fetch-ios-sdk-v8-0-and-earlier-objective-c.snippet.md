---
id: ios-client-sdk/sdk-docs/background-fetch-ios-sdk-v8-0-and-earlier-objective-c
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "iOS SDK v8.0 and earlier (Objective-C) in section \"Background fetch\""
# TODO(validate): . See _sdk-docs-port-notes.md.
---

```objectivec
LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key"];
config.backgroundFlagPollingInterval = 3600;
```
