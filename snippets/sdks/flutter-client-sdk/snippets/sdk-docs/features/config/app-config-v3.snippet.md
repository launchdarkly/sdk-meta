---
id: flutter-client-sdk/sdk-docs/features/config/app-config-v3
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Application metadata configuration example for Flutter.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v3

---

```dart
LDConfig config = LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
  .applicationId('authentication-service')
  .applicationName('Authentication-Service')
  .applicationVersion('1.0.0')
  .applicationVersionName('v1')
  .build();
```
