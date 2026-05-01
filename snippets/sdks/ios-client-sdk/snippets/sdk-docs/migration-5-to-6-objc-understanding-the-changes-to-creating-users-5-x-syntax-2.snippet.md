---
id: ios-client-sdk/sdk-docs/migration-5-to-6-objc-understanding-the-changes-to-creating-users-5-x-syntax-2
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "5.x syntax in section \"Understanding the changes to creating users\""
---

```objectivec
LDUser *user1 = [[LDUser alloc] initWithKey:@"example-user-key"];
user1.name = @"Jane Smith";
LDUser *user2 = [[LDUser alloc] initWithKey:@"example-user-key"];
user2.name = @"John Smith";
// Results in true
[user1 isEqual:user2]
```
