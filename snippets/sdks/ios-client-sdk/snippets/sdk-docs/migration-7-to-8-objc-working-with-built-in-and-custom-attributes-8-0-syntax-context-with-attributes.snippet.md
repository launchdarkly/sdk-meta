---
id: ios-client-sdk/sdk-docs/migration-7-to-8-objc-working-with-built-in-and-custom-attributes-8-0-syntax-context-with-attributes
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "8.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```objectivec
LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:@"example-context-key"];
[builder kindWithKind:@"user"];
[builder nameWithName:@"Sandy Smith"];
[builder trySetValueWithName:@"email" value:[LDValue ofString:@"sandy@example.com"]];

LDContext *context = builder.build.success;
```
