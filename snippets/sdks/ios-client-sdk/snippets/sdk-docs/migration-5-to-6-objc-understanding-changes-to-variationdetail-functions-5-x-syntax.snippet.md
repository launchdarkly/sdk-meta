---
id: ios-client-sdk/sdk-docs/migration-5-to-6-objc-understanding-changes-to-variationdetail-functions-5-x-syntax
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "5.x syntax in section \"Understanding changes to variationDetail functions\""
---

```objectivec
LDArrayEvaluationDetail *result = [[LDClient get] arrayVariationDetailForKey:@"example-flag-key" defaultValue:@[]];
NSArray* resultValue = result.value;
NSDictionary<NSString*, id> *reason = result.reason;
```
