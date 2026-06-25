---
id: flutter-client-sdk/sdk-docs/features/flagchanges/flag-changes-v3
sdk: flutter-client-sdk
kind: reference
lang: dart
description: Flag change subscription example for Flutter SDK v3.x.
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only-v3

---

```dart
LDFlagUpdatedCallback listener = (String flagKey) {
  LDClient.boolVariation(flagKey, false).then((bool val) {
      print('${flagKey}: ${val}');
  });
};

await LDClient.registerFeatureFlagListener('example-flag-key', listener);
```
