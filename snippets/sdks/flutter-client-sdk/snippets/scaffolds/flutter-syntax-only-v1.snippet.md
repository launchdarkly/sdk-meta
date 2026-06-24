---
id: flutter-client-sdk/scaffolds/flutter-syntax-only-v1
sdk: flutter-client-sdk
kind: scaffold
lang: dart
file: lib/main.dart
description: |
  Parse-and-type-check validator for Flutter client SDK doc fragments
  that target the v1.x API surface (`LDClient.alias(user, previousUser)`
  — removed at 2.0).

  No validator container pins a 1.x SDK (the 1.x package predates
  Dart null safety and cannot resolve against the toolchains the
  flutter validator images carry), so this scaffold routes through
  the current `flutter-client` container with a self-contained stub:
  a local `LDClient` class declares just the static member the
  v1-era fragments call. The SDK package is deliberately not
  imported so the stub is authoritative.
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
//IMPORT_LIFT_TARGET

// Stub of the v1.x static LDClient surface. Only the members the
// doc fragments call are declared; arguments are opaque.
// ignore: avoid_classes_with_only_static_members
class LDClient {
  static Future<void> alias(dynamic user, dynamic previousUser) async {}
}

// ignore: unused_element
Future<void> _wrappee() async {
  // Stub locals for the ambient user names the v1-era fragments
  // assume earlier snippets created.
  // ignore: unused_local_variable
  dynamic newUser = Object();
  // ignore: unused_local_variable
  dynamic previousUser = Object();
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
