---
id: flutter-client-sdk/sdk-docs/features/config/service-endpoint-configuration-v3-eu
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Service endpoint configuration example for Flutter.
---

```dart
LDConfig ldConfig = LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
  .streamUri('https://clientstream.eu.launchdarkly.com')
  .pollUri('https://clientsdk.eu.launchdarkly.com')
  .eventsUri('https://events.eu.launchdarkly.com')
  .build();
```
