---
id: flutter-client-sdk/sdk-docs/features/config/service-endpoint-configuration-v3-federal
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Service endpoint configuration example for Flutter.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v3

---

```dart
LDConfig ldConfig = LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
  .streamUri('https://clientstream.launchdarkly.us')
  .pollUri('https://clientsdk.launchdarkly.us')
  .eventsUri('https://events.launchdarkly.us')
  .build();
```
