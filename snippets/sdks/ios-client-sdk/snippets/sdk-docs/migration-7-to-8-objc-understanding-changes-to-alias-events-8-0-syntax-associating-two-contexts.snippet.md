---
id: ios-client-sdk/sdk-docs/migration-7-to-8-objc-understanding-changes-to-alias-events-8-0-syntax-associating-two-contexts
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "8.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```objectivec
LDContextBuilder *userBuilder = [[LDContextBuilder alloc] initWithKey:@"example-user-key"];
LDContextBuilder *deviceBuilder = [[LDContextBuilder alloc] initWithKey:@"example-device-key"];
[deviceBuilder kindWithKind:@"device"];

LDMultiContextBuilder *multiBuilder = [[LDMultiContextBuilder alloc] init];
[multiBuilder addContextWithContext:userBuilder.build.success];
[multiBuilder addContextWithContext:deviceBuilder.build.success];

LDContext *multiContext = multiBuilder.build.success;

[[LDClient get] identifyWithContext:multiContext];
```
