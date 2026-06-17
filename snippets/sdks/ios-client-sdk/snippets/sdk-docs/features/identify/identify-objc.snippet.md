---
id: ios-client-sdk/sdk-docs/features/identify/identify-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Identify example for the iOS SDK v8.0+ (Objective-C).
---

```objectivec
LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:@"example-context-key"];
ContextBuilderResult *result = builder.build;

if (result.success) {
    LDContext *newContext = result.success;

    [[LDClient get] identifyWithContext:newContext];

    // You can also call identify with a completion
    [[LDClient get] identifyWithContext:newContext completion:^() {
        // Flags have been retrieved for the new context
    }];
}
```
