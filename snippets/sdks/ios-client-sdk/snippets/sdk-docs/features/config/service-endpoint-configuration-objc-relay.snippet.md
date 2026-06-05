---
id: ios-client-sdk/sdk-docs/features/config/service-endpoint-configuration-objc-relay
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Service endpoint configuration example for iOS.

---

```objectivec
LDConfig *ldConfig = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key" autoEnvAttributes:AutoEnvAttributesEnabled];
ldConfig.streamUrl = [NSURL URLWithString:@"https://your-relay-proxy.com:8030"];
ldConfig.baseUrl = [NSURL URLWithString:@"https://your-relay-proxy.com:8030"];
ldConfig.eventsUrl = [NSURL URLWithString:@"https://your-relay-proxy.com:8030"];
```
