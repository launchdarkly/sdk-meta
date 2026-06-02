---
id: flutter-client-sdk/sdk-docs/features/config/service-endpoint-configuration-v3-relay
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Service endpoint configuration example for Flutter.
---

```dart
LDConfig ldConfig = LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
  .streamUri('https://your-relay-proxy.com:8030')
  .pollUri('https://your-relay-proxy.com:8030')
  .eventsUri('https://your-relay-proxy.com:8030')
  .build();
```
