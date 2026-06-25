---
id: ios-client-sdk/sdk-docs/features/privateattrs/config-v9-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Private attribute configuration for iOS SDK v9.0 (Objective-C).
---

```objectivec
// All attributes marked private
LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key" autoEnvAttributes:AutoEnvAttributesEnabled];
[config setAllContextAttributesPrivate:YES];
[LDClient startWithConfiguration:config context:context completion:nil];

// Two attributes marked private
LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key" autoEnvAttributes:AutoEnvAttributesEnabled];
[builder nameWithName:@"name"];
[builder trySetValueWithName:@"email" value:[LDValue ofString:@"example@email.com"]];
NSDictionary<NSString *, LDValue *> *address = @{
    @"street": [LDValue ofString:@"Main St"],
    @"city": [LDValue ofString:@"Springfield"]
};
LDContext *context = builder.build.success;

LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key"];
[config setPrivateContextAttributes:@[@"email", @"address"]];
```
