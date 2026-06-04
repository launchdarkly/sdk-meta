---
id: ios-client-sdk/sdk-docs/features/config/index-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: SDK configuration example for iOS.

---

```objectivec
LDConfig *ldConfig = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key" autoEnvAttributes:AutoEnvAttributesEnabled];
ldConfig.connectionTimeout = 10.0;
ldConfig.eventFlushInterval = 30.0;
```
