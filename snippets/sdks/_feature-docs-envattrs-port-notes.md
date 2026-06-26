# Port notes: /sdk/features/environment-attributes

Source: `ld-docs-private` `fern/topics/sdk/features/environment-attributes.mdx`.
8 code blocks extracted into `sdk-docs/features/envattrs/` snippets
across 5 SDKs. All but one (iOS Objective-C) are bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX.

- **Flutter v4 example** (`flutter-client-sdk/.../auto-env-attributes-v4`):
  the published statement had no trailing semicolon, which is a Dart
  syntax error. Semicolon added; the call itself
  (`LDConfig(CredentialSource.fromEnvironment(), AutoEnvAttributes.enabled)`)
  matches the real v4 API.

## Validation routing added in this port

None. Every bound snippet reuses an existing scaffold:

- `dotnet-client-sdk/scaffolds/csharp-client-syntax-only` (.NET C#)
- `android-client-sdk/scaffolds/java-syntax-only` and
  `kotlin-syntax-only` (Android; both compile the real v5 aar, and
  both already import the nested `LDConfig.Builder.AutoEnvAttributes`
  enum the bodies reference unqualified)
- `flutter-client-sdk/scaffolds/flutter-syntax-only` (v4) and
  `flutter-syntax-only-v3` (v3.x; `AutoEnvAttributes.Enabled` is the
  real v3 spelling — v3 predates the lowerCamelCase rename)
- `ios-client-sdk/scaffolds/swift-syntax-only` (Swift; macOS-only
  native harness, so only CI can run it)
- `react-native-client-sdk/scaffolds/react-native-syntax-only`

## Known non-binds

- `ios-client-sdk/.../auto-env-attributes-objc` — no Objective-C parse
  scaffold exists; the iOS validator is the macOS-only native harness
  (same blocker as the evaluating and evaluation-reasons ports' objc
  snippets). Wiring it up requires either an Objective-C target in the
  xcodegen scaffold or a clang -fsyntax-only stub harness.
