---
id: ios-client-sdk/sdk-docs/features/contextconfig/context-example-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Context example for iOS SDK v8.0+ (Objective-C).

---

```objectivec
LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:@"example-user-key"];
[builder trySetValueWithName:@"name" value:[LDValue ofString:@"Sandy"]];
[builder trySetValueWithName:@"email" value:[LDValue ofString:@"sandy@example.com"]];

LDContext *context = builder.build.success;
```
