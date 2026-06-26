---
id: ios-client-sdk/sdk-docs/features/anonymous/anonymous-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Anonymous context example for iOS (Objective-C).
---

```objectivec
LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:@"example-context-key"];
[builder trySetValueWithName:@"anonymous" value:[LDValue ofBool:true]];

LDContext *context = builder.build.success;
```
