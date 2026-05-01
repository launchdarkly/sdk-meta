---
id: ios-client-sdk/sdk-docs/migration-7-to-8-objc-referencing-properties-of-an-attribute-object-8-0-syntax-context-with-object-attributes
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "8.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
---

```objectivec
LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:@"example-context-key"];
NSDictionary<NSString *, LDValue *> *address = @{
    @"street": [LDValue ofString:@"Main St"],
    @"city": [LDValue ofString:@"Springfield"]
};
[builder trySetValueWithName:@"address" value:[LDValue ofDict:address]];

LDContext *context = builder.build.success;
```
