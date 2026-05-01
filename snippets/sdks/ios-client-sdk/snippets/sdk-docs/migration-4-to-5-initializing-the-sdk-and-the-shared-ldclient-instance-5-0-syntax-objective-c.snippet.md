---
id: ios-client-sdk/sdk-docs/migration-4-to-5-initializing-the-sdk-and-the-shared-ldclient-instance-5-0-syntax-objective-c
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "5.0 syntax (Objective-C) in section \"Initializing the SDK and the shared LDClient instance\""
---

```objectivec
LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key"];
LDUser *user = [[LDUser alloc] initWithKey:@"example-user-key"];
[LDClient start:config user:user];
LDClient *shared = [LDClient get];
```
