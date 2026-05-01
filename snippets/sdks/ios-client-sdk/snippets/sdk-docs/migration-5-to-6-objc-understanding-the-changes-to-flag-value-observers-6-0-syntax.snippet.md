---
id: ios-client-sdk/sdk-docs/migration-5-to-6-objc-understanding-the-changes-to-flag-value-observers-6-0-syntax
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "6.0 syntax in section \"Understanding the changes to flag value observers\""
---

```objectivec
[[LDClient get] observe:@"example-flag-key" owner:self handler:^(LDChangedFlag * _Nonnull changedFlag) {
    LDValue* newValue = changedFlag.newValue;
}];
```
