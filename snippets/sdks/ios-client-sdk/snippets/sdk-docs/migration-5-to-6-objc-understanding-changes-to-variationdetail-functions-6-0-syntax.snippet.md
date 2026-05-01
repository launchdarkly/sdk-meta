---
id: ios-client-sdk/sdk-docs/migration-5-to-6-objc-understanding-changes-to-variationdetail-functions-6-0-syntax
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "6.0 syntax in section \"Understanding changes to variationDetail functions\""
---

```objectivec
LDJSONEvaluationDetail *result = [[LDClient get] jsonVariationDetailForKey:@"example-flag-key" defaultValue:[LDValue ofArray:@[]]];
LDValue* resultValue = result.value;
NSDictionary<NSString*, LDValue*> *reason = result.reason;
```
