---
id: ios-client-sdk/sdk-docs/migration-8-to-9-understanding-what-was-removed-ios-sdk-v8-context-with-key-objective-c
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "iOS SDK v8+, context with key (Objective-C) in section \"Understanding what was removed\""
---

```objectivec
LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:@"example-context-key"];
ContextBuilderResult *result = builder.build;

if (result.success) {
  LDContext *context = result.success;
}
```
