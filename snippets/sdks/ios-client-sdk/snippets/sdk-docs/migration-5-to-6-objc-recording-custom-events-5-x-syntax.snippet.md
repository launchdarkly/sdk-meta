---
id: ios-client-sdk/sdk-docs/migration-5-to-6-objc-recording-custom-events-5-x-syntax
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "5.x syntax in section \"Recording custom events\""
---

```objectivec
NSError* err = nil;
NSDictionary* data = @{@"abc": @123};
[[LDClient get] trackWithKey:@"key" data:data error:&err];
if (err != nil) {
    // Do something with the error
}
```
