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
//IMPORT_LIFT_TARGET

// ignore: unused_element
Future<void> _wrappee() async {
  // Stub locals so wrappee bodies referencing `client`, `context`,
  // `flagKey` resolve at compile time. Initialize each with a
  // value-of-the-right-type expression so Dart's analyzer doesn't
  // flag a definitely-unassigned read. The body lives in a nested
  // block so it can re-declare `context`, `flagKey`, etc. without
  // colliding with the stubs (Dart allows shadowing across blocks).
  //
  // Body imports (e.g. `import 'package:launchdarkly_flutter_client_sdk/...';`)
  // get lifted to module scope by the harness's awk pre-step using the
  // BODY_BEGIN / BODY_END / IMPORT_LIFT_TARGET markers — Dart forbids
  // imports inside a function body, so a fragment that shows an
  // install-time import would otherwise fail to compile.
  // ignore: unused_local_variable
  dynamic client = Object();
  // ignore: unused_local_variable
  dynamic context = Object();
  // ignore: unused_local_variable
  String flagKey = '';
  // Flag-change fragments cancel a stream subscription the docs
  // assume was created by an earlier subscribe fragment.
  // ignore: unused_local_variable
  dynamic sub = Object();
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
