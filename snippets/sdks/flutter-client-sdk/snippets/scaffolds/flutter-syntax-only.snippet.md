---
id: flutter-client-sdk/scaffolds/flutter-syntax-only
sdk: flutter-client-sdk
kind: scaffold
lang: dart
file: lib/main.dart
description: |
  Parse-only validator for Flutter client SDK doc fragments.

  Renders the EXAM-HELLO success line as an actual `Text` widget so
  the playwright DOM check (via flutter-client/harness/check.js) sees
  it. The wrappee body lives inside `_wrappee()`, an async function
  that's never invoked at runtime — its references to `client`,
  `LDClient.start(...)`, etc. compile against the
  launchdarkly_flutter_client_sdk package the prebuilt project
  already pulls in.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: flutter-client
  entrypoint: lib/main.dart
---

```dart
import 'package:flutter/material.dart';
// ignore: unused_import
import 'package:launchdarkly_flutter_client_sdk/launchdarkly_flutter_client_sdk.dart';

// ignore: unused_element
Future<void> _wrappee() async {
  // Stub locals so the wrappee body's symbols resolve at compile time.
  // Never reached at runtime — main() short-circuits to the success
  // widget below.
  // ignore: unused_local_variable
  late final dynamic client;
  // ignore: unused_local_variable
  late final dynamic context;
  // ignore: unused_local_variable
  late final String flagKey;
{{ body }}
}

void main() {
  runApp(const MaterialApp(
    home: Scaffold(body: Center(child: Text('feature flag evaluates to true'))),
  ));
}
```
