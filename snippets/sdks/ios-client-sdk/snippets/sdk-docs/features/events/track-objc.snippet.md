---
id: ios-client-sdk/sdk-docs/features/events/track-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Custom event tracking example for iOS (Objective-C).

---

```objectivec
LDValue *data = [LDValue ofArray:@[[LDValue ofBool:YES], [LDValue ofNumber:@5.5]]];
[[LDClient get] trackWithKey:@"example-event-key" data:data];
```
