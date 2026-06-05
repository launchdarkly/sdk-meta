---
id: flutter-client-sdk/sdk-docs/features/config/index-v3
sdk: flutter-client-sdk
kind: reference
lang: dart
description: SDK configuration example for Flutter.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v3

---

```dart
LDConfig ldConfig = LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
    .evaluationReasons(true)
    .build();
```
