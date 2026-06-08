---
id: ios-client-sdk/sdk-docs/features/evaluating/evaluating-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Flag evaluation example for iOS.

---

```objectivec
BOOL boolFlagValue = [[LDClient get] boolVariationForKey:@"example-bool-flag-key" defaultValue:NO];
LDValue* jsonFlagValue = [[LDClient get] jsonVariationForKey:@"json-flag-key-456def" defaultValue:[LDValue ofNull]];
```
