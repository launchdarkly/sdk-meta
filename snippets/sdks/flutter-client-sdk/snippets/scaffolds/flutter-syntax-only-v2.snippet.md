---
id: flutter-client-sdk/scaffolds/flutter-syntax-only-v2
sdk: flutter-client-sdk
kind: scaffold
lang: dart
file: lib/main.dart
description: |
  Parse-and-type-check validator for Flutter client SDK doc fragments
  that target the v2.x API surface (e.g. `LDConfigBuilder(...).build()`,
  `LDClient.startWithContext(...)`).

  Routes through the `flutter-client-v2` validator container, which
  pins `launchdarkly_flutter_client_sdk` to a 2.x release. The
  current-version `flutter-syntax-only` scaffold compiles against
  the latest SDK and won't resolve names removed at 3.x; this
  scaffold exists so v2.x-specific docs validate against the actual
  v2.x SDK that the docs cover.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: flutter-client-v2
  entrypoint: lib/main.dart
---

```dart
import 'package:flutter/material.dart';
// ignore: unused_import
import 'package:launchdarkly_flutter_client_sdk/launchdarkly_flutter_client_sdk.dart';
//IMPORT_LIFT_TARGET

// ignore: unused_element
Future<void> _wrappee() async {
  // ignore: unused_local_variable
  dynamic client = Object();
  // ignore: unused_local_variable
  dynamic context = Object();
  // ignore: unused_local_variable
  String flagKey = '';
  {
//BODY_BEGIN
{{ body }}
//BODY_END
  }
}

void main() {
  runApp(const MaterialApp(
    home: Scaffold(body: Center(child: Text('feature flag evaluates to true'))),
  ));
}
```
