---
id: ios-client-sdk/sdk-docs/features/privateattrs/config-v8-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Private attribute configuration for iOS SDK v8.x (Objective-C).
---

```objectivec
// All attributes marked private
LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key"];
[config setAllContextAttributesPrivate:YES];
[LDClient startWithConfiguration:config context:context completion:nil];

// Two attributes marked private
LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:@"example-context-key"];
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
