---
id: ios-client-sdk/sdk-docs/features/contextconfig/multi-context-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Multi-context example for iOS SDK v8.0+ (Objective-C).

---

```objectivec
LDContextBuilder *userBuilder = [[LDContextBuilder alloc] initWithKey:@"example-user-key"];
LDContextBuilder *deviceBuilder = [[LDContextBuilder alloc] initWithKey:@"example-device-key"];
[deviceBuilder kindWithKind:@"device"];

LDMultiContextBuilder *multiBuilder = [[LDMultiContextBuilder alloc] init];
[multiBuilder addContextWithContext:userBuilder.build.success];
[multiBuilder addContextWithContext:deviceBuilder.build.success];

LDContext *multiContext = multiBuilder.build.success;
```
