---
id: flutter-client-sdk/sdk-docs/migration-3-to-4-subscribing-to-flag-changes-flutter-sdk-v3-x
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK v3.x in section \"Subscribing to flag changes\""
---

```dart
LDFlagUpdatedCallback listener = (String flagKey) {
  LDClient.boolVariation(flagKey, false).then((bool val) {
      print('${flagKey}: ${val}');
  });
};

await LDClient.registerFeatureFlagListener('example-flag-key', listener);
```
