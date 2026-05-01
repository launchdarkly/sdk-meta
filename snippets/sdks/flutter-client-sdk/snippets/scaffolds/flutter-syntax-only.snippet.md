---
id: flutter-client-sdk/scaffolds/flutter-syntax-only
sdk: flutter-client-sdk
kind: scaffold
lang: dart
file: lib/main.dart
description: |
  Parse-only validator for Flutter client SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: flutter-client
  entrypoint: lib/main.dart
---

```dart
void _wrappee() {
{{ body }}
}

void main() {
  print('feature flag evaluates to true');
}
```
