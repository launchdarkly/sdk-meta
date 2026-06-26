---
id: ios-client-sdk/sdk-docs/features/logging/logging-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Logger configuration example for iOS SDK v9.x (Objective-C).

---

```objectivec
@import OSLog;

LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key" autoEnvAttributes:AutoEnvAttributesEnabled];
config.logger = os_log_create("your.preferred.subsystem", "ld-sdk");

// You can disable all SDK logging by setting this property to the shared disabled logger
config.logger = OS_LOG_DISABLED;
```
