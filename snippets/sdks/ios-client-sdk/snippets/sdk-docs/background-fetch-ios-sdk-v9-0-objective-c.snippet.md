---
id: ios-client-sdk/sdk-docs/background-fetch-ios-sdk-v9-0-objective-c
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "iOS SDK v9.0 (Objective-C) in section \"Background fetch\""
# TODO(validate): swift-syntax-only compiles Swift only. To validate, add an Objective-C scaffold
# (or a harness branch on lang:) and a project.yml with an Objective-C target.
---

```objectivec
LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key", autoEnvAttributes:AutoEnvAttributesEnabled];
config.backgroundFlagPollingInterval = 3600;
```
