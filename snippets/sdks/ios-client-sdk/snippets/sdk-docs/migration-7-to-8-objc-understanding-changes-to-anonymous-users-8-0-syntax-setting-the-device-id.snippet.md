---
id: ios-client-sdk/sdk-docs/migration-7-to-8-objc-understanding-changes-to-anonymous-users-8-0-syntax-setting-the-device-id
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "8.0 syntax, setting the device ID in section \"Understanding changes to anonymous users\""
---

```objectivec
NSString *key = [[[UIDevice currentDevice] identifierForVendor] UUIDString];
LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:key];
[builder kindWithKind:@"device"];
LDContext *context = builder.build.success;
```
