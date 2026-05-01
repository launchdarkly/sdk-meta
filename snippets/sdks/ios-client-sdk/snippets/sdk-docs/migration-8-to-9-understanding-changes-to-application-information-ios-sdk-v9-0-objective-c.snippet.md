---
id: ios-client-sdk/sdk-docs/migration-8-to-9-understanding-changes-to-application-information-ios-sdk-v9-0-objective-c
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "iOS SDK v9.0 (Objective-C) in section \"Understanding changes to application information\""
---

```objc
LDApplicationInfo *applicationInfo = [[LDApplicationInfo alloc] init];
[applicationInfo applicationIdentifier:@"authentication-service"];
[applicationInfo applicationName:@"Authentication-Service"];
[applicationInfo applicationVersion:@"1.0.0"];
[applicationInfo applicationVersionName:@"v1"];

LDConfig *config = [[LDConfig alloc] initWithMobileKey:mobileKey autoEnvAttributes:true];
config.applicationInfo = applicationInfo;
```
