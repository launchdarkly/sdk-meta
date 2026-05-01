---
id: ios-client-sdk/sdk-docs/migration-5-to-6-objc-understanding-the-changes-to-flag-value-observers-5-x-syntax
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "5.x syntax in section \"Understanding the changes to flag value observers\""
---

```objectivec
[[LDClient get] observeBool:@"example-flag-key" owner:self handler:^(LDBoolChangedFlag * _Nonnull changedFlag) {
    Bool newValue = changedFlag.newValue;
}];
```
