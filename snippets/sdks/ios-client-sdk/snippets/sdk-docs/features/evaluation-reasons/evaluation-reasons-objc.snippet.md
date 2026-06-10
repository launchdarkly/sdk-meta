---
id: ios-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Flag evaluation reason example for iOS (Objective-C).

---

```objectivec
config.evaluationReasons = true;
[LDClient startWithConfiguration:config context:result.success completion: nil];

LDBoolEvaluationDetail *detail = [[LDClient get] boolVariationDetailForKey:@"example-flag-key" defaultValue:NO];

BOOL value = detail.value;
NSInteger variationIndex = detail.variationIndex;
NSDictionary<NSString*, LDValue*> *reason = detail.reason;
```
