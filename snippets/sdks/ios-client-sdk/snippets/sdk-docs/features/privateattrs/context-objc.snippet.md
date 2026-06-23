---
id: ios-client-sdk/sdk-docs/features/privateattrs/context-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Marking context attributes private with the context builder for iOS SDK v8.0+ (Objective-C).
---

```objectivec
LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:@"example-context-key"];
[builder trySetValueWithName:@"name" value:[LDValue ofString:@"Sandy"]];

NSArray *groups = [NSArray arrayWithObjects:[LDValue ofString:@"microsoft"], nil];
[builder trySetValueWithName:@"group" value:[LDValue ofArray:groups]];

[builder addPrivateAttributeWithReference:[[Reference alloc] initWithValue:@"email"]];
[builder addPrivateAttributeWithReference:[[Reference alloc] initWithValue:@"group"]];

LDContext *context = builder.build.success;
```
