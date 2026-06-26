---
id: ios-client-sdk/sdk-docs/features/multienv/init-v9-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Multi-environment configuration for iOS SDK v9.x (Objective-C).

---

```objectivec
LDContext *context = [[LDContextBuilder alloc] initWithKey:@"example-context-key"].build.success;
LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key" autoEnvAttributes:AutoEnvAttributesEnabled];
NSDictionary *otherMobileKeys = @{@"platform": @"platform-example-mobile-key", @"core": @"core-example-mobile-key"};
NSError *err = nil;
[config setSecondaryMobileKeys:otherMobileKeys error:&err];
[LDClient startWithConfiguration:config context:context completion:nil];
```
