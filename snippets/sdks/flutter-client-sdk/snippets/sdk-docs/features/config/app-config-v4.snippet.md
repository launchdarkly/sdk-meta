---
id: flutter-client-sdk/sdk-docs/features/config/app-config-v4
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Application metadata configuration example for Flutter.
---

```dart
final config = LDConfig(
  CredentialSource.fromEnvironment(),
  AutoEnvAttributes.enabled,
  applicationInfo: ApplicationInfo(
    applicationId: 'authentication-service',
    applicationName: 'Authentication-Service',
    applicationVersion: '1.0.0',
    applicationVersionName: 'v1',
  ),
)

```
