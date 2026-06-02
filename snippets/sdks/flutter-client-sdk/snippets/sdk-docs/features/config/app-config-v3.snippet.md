---
id: flutter-client-sdk/sdk-docs/features/config/app-config-v3
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Application metadata configuration example for Flutter.
---

```dart
LDConfig config = LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
  .applicationId('authentication-service')
  .applicationName('Authentication-Service')
  .applicationVersion('1.0.0')
  .applicationVersionName('v1')
  .build();
```
