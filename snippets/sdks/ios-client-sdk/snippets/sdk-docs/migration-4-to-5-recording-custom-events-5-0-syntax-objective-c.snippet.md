---
id: ios-client-sdk/sdk-docs/migration-4-to-5-recording-custom-events-5-0-syntax-objective-c
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "5.0 syntax (Objective-C) in section \"Recording custom events\""
---

```objectivec
NSError* err = nil;
[[LDClient get] trackWithKey: @"key" data: @{@"abc": @123} error: &err];
if (err != nil) {
    // Do something with the error
}
```
