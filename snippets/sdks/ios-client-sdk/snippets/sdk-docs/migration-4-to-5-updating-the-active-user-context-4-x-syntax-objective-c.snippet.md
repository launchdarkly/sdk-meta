---
id: ios-client-sdk/sdk-docs/migration-4-to-5-updating-the-active-user-context-4-x-syntax-objective-c
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "4.x syntax (Objective-C) in section \"Updating the active user context\""
---

```objectivec
LDUser *newUser = [[LDUser alloc] initWithKey:@"example-user-key"];
[LDClient sharedInstance].user = newUser;
```
