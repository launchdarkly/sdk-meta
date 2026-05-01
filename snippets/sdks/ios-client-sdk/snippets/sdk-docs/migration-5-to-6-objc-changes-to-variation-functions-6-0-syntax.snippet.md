---
id: ios-client-sdk/sdk-docs/migration-5-to-6-objc-changes-to-variation-functions-6-0-syntax
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "6.0 syntax in section \"Changes to variation functions\""
---

```objectivec
LDValue* result = [[LDClient get] jsonVariationForKey:@"example-flag-key" defaultValue:[LDValue ofArray:@[]]];
```
