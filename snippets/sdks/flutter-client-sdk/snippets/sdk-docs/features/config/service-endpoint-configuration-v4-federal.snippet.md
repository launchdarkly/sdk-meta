---
id: flutter-client-sdk/sdk-docs/features/config/service-endpoint-configuration-v4-federal
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Service endpoint configuration example for Flutter.
---

```dart
final config = LDConfig(
  CredentialSource.fromEnvironment(),
  autoEnvAttributes.enabled,
  serviceEndpoints: ServiceEndpoints.custom(
    streaming: 'https://clientstream.launchdarkly.us',
    polling: 'https://clientsdk.launchdarkly.us',
    events: 'https://events.launchdarkly.us',
  )
);
```
