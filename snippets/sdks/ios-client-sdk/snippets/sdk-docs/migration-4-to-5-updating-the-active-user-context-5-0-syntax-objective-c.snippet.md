---
id: ios-client-sdk/sdk-docs/migration-4-to-5-updating-the-active-user-context-5-0-syntax-objective-c
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "5.0 syntax (Objective-C) in section \"Updating the active user context\""
---

```objectivec
LDUser *newUser = [[LDUser alloc] initWithKey:@"example-user-key"];
[[LDClient get] identify:newUser];
```
