---
id: ios-client-sdk/sdk-docs/background-fetch-ios-sdk-v9-0-objective-c
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "iOS SDK v9.0 (Objective-C) in section \"Background fetch\""
# TODO(validate): . See _sdk-docs-port-notes.md.
---

```objectivec
LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key", autoEnvAttributes:AutoEnvAttributesEnabled];
config.backgroundFlagPollingInterval = 3600;
```
