---
id: ios-client-sdk/sdk-docs/migration-5-to-6-objc-understanding-the-changes-to-creating-users-6-0-syntax
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "6.0 syntax in section \"Understanding the changes to creating users\""
---

```objectivec
LDUser *user = [[LDUser alloc] initWithKey:@"example-user-key"];
user.custom = @{@"group": [LDValue ofString:@"beta"]};
```
