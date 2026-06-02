---
id: ios-client-sdk/sdk-docs/features/config/service-endpoint-configuration-objc-eu
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Service endpoint configuration example for iOS.
---

```objectivec
LDConfig *ldConfig = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key" autoEnvAttributes:AutoEnvAttributesEnabled];
ldConfig.streamUrl = [NSURL URLWithString:@"https://clientstream.eu.launchdarkly.com"];
ldConfig.baseUrl = [NSURL URLWithString:@"https://clientsdk.eu.launchdarkly.com"];
ldConfig.eventsUrl = [NSURL URLWithString:@"https://events.eu.launchdarkly.com"];
```
