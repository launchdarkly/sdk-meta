---
id: ios-client-sdk/sdk-docs/features/offlinemode/offline-mode-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Offline mode example for iOS (Objective-C).

---

```objectivec
[[LDClient get] setOnline:NO];
[[LDClient get] setOnline:YES completion:^() {
    // Client is online
}];
```
