---
id: ios-client-sdk/sdk-docs/migration-7-to-8-objc-understanding-changes-to-private-attributes-8-0-syntax-two-attributes-marked-private
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "8.0 syntax, two attributes marked private in section \"Understanding changes to private attributes\""
---

```objectivec
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
