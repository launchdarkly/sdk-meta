---
id: ios-client-sdk/sdk-docs/features/config/app-config-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Application metadata configuration example for iOS.
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
