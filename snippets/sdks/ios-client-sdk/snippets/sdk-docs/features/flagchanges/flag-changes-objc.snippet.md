---
id: ios-client-sdk/sdk-docs/features/flagchanges/flag-changes-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Flag change observation example for iOS (Objective-C).

---

```objectivec
__weak typeof(self) weakSelf = self;
[[LDClient get] observeKeys:self.flagKeys owner:self handler:^(NSDictionary<NSString *,LDChangedFlag *> * _Nonnull changedFlags) {
    __strong typeof(weakSelf) strongSelf = weakSelf;
    for (NSString* flagKey in changedFlags.allKeys) {
        LDChangedFlag *changedFlag = changedFlags[flagKey];
        NSLog(@"Flag %@ changed from %@ to %@", flagKey, changedFlag.oldValue, changedFlag.newValue);
    }
}];
[[LDClient get] observeFlagsUnchangedWithOwner:self handler:^{
    __strong typeof(weakSelf) strongSelf = weakSelf;
    // No flags changed
}];
```
