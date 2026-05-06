---
id: flutter-client-sdk/sdk-docs/evaluate-a-flag-flutter-sdk
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "Flutter SDK in section \"Evaluate a flag\""
validation:
  scaffold: flutter-client-sdk/scaffolds/flutter-syntax-only
---

```dart
bool showFeature = await client.boolVariation(flagKey, false);
if (showFeature) {
    // Application code to show the feature
}
else {
    // The code to run if the feature is off
}
```
