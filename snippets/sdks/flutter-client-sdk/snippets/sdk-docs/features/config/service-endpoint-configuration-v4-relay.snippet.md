---
id: flutter-client-sdk/sdk-docs/features/config/service-endpoint-configuration-v4-relay
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Service endpoint configuration example for Flutter.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only

---

```dart
final config = LDConfig(
  CredentialSource.fromEnvironment(),
  AutoEnvAttributes.enabled,
  serviceEndpoints: ServiceEndpoints.relayProxy('https://your-relay-proxy.com:8030')
);
```
