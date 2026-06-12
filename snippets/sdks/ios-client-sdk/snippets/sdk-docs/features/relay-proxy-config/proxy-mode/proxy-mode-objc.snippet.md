---
id: ios-client-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Proxy mode configuration example for iOS (Objective-C).
---

```objectivec
LDConfig *ldConfig = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key" autoEnvAttributes:AutoEnvAttributesEnabled];
ldConfig.streamUrl = [NSURL URLWithString:@"https://your-relay-proxy.com:8030"];
ldConfig.baseUrl = [NSURL URLWithString:@"https://your-relay-proxy.com:8030"];
ldConfig.eventsUrl = [NSURL URLWithString:@"https://your-relay-proxy.com:8030"];
```
