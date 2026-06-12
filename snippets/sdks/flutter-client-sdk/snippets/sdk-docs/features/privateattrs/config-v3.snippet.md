---
id: flutter-client-sdk/sdk-docs/features/privateattrs/config-v3
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Private attribute configuration for Flutter SDK v3.x.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v3

---

```dart
// All attributes marked private
LDConfig config = LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
    .allAttributesPrivate(true)
    .build();

// Two attributes marked private
LDConfig ldConfig = new LDConfigBuilder('example-mobile-key', AutoEnvAttributes.Enabled)
    .privateAttributes({'email', 'group'})
    .build();
```
