---
id: ios-client-sdk/sdk-docs/migration-7-to-8-objc-understanding-differences-between-users-and-contexts-8-0-syntax-context-with-key
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "8.0 syntax, context with key in section \"Understanding differences between users and contexts\""
---

```objectivec
LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:@"example-context-key"];
ContextBuilderResult *result = builder.build;

if (result.success) {
  LDContext *context = result.success;
}
```
