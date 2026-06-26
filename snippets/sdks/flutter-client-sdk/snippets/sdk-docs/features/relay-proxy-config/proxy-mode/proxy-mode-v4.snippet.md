---
id: flutter-client-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Proxy mode configuration example for Flutter.
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
