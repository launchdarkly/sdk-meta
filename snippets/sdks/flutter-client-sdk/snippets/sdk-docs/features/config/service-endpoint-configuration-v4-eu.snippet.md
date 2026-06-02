---
id: flutter-client-sdk/sdk-docs/features/config/service-endpoint-configuration-v4-eu
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
    streaming: 'https://clientstream.eu.launchdarkly.com',
    polling: 'https://clientsdk.eu.launchdarkly.com',
    events: 'https://events.eu.launchdarkly.com',
  )
);
```
