---
id: flutter-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-application-information-flutter-sdk-v3-0-dart
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v3.0 (Dart) in section \"Understanding changes to application information\""
---

```dart
final LDConfig ldConfig = LDConfigBuilder("example-mobile-key", AutoEnvAttributes.Enabled)
  .applicationId("authentication-service")
  .applicationName("Authentication-Service")
  .applicationVersion("1.0.0")
  .applicationVersionName("v1")
  .build();
```
