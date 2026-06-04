---
id: ios-client-sdk/sdk-docs/features/config/service-endpoint-configuration-objc-federal
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Service endpoint configuration example for iOS.

---

```objectivec
LDConfig *ldConfig = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key" autoEnvAttributes:AutoEnvAttributesEnabled];
ldConfig.streamUrl = [NSURL URLWithString:@"https://clientstream.launchdarkly.us"];
ldConfig.baseUrl = [NSURL URLWithString:@"https://clientsdk.launchdarkly.us"];
ldConfig.eventsUrl = [NSURL URLWithString:@"https://events.launchdarkly.us"];
```
