---
id: flutter-client-sdk/sdk-docs/features/config/app-config-v2
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Application metadata configuration example for Flutter.
---

```dart
LDConfig config = LDConfigBuilder('example-mobile-key')
  .applicationId('authentication-service')
  .applicationVersion('1.0.0')
  .build();
```
