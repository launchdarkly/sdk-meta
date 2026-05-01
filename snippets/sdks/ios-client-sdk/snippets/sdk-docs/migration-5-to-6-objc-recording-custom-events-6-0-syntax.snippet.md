---
id: ios-client-sdk/sdk-docs/migration-5-to-6-objc-recording-custom-events-6-0-syntax
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "6.0 syntax in section \"Recording custom events\""
---

```objectivec
LDValue* data = [LDValue ofDict:@{@"abc": [LDValue ofNumber:@123]}];
[[LDClient get] trackWithKey: @"key" data:data];
```
