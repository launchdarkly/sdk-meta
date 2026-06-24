---
id: ios-client-sdk/sdk-docs/features/contextconfig/context-kind-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Context with a non-user kind for iOS SDK v8.0+ (Objective-C).

---

```objectivec
LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:@"example-organization-key"];
[builder kindWithKind:@"organization"];

LDContext *context = builder.build.success;
```
