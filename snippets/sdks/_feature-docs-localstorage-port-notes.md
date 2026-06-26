# Port notes: /sdk/features/local-storage

Source: `ld-docs-private` `fern/topics/sdk/features/local-storage.mdx`.
13 code blocks extracted into `sdk-docs/features/localstorage/`
snippets across 8 SDKs. All 13 are bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **JavaScript SDK v3.x TypeScript** (`js-client-sdk/.../localstorage-v3-ts`):
  the block imported `LDClient` as a named import and then called
  `LDClient.initialize(...)` — in the v3 typings `LDClient` is an
  interface (a type, not a value), so the call cannot work. Split
  into a namespace import (`import * as LDClient`) for the
  `initialize` call plus a named import for `LDOptions`, matching the
  namespace-import pattern the config and bootstrapping pages use for
  the same SDK.

## Validation routing added in this port

- `android-client-sdk/scaffolds/kotlin-syntax-only` — added a
  module-scope `application: Application` stub. The Kotlin init
  fragment calls `LDClient.init(application, ...)` with the ambient
  `application` an Activity host provides; the scaffold splices
  bodies into an `Application` subclass, which has no such property
  (`getApplication()` exists only on `android.app.Activity`).

No new scaffolds or validators. The Electron TypeScript fragment's
top-level `import * as LaunchDarkly from 'launchdarkly-electron-client-sdk'`
passes through `electron-syntax-only`'s dead-function path as is
(same shape the bootstrapping port validated), so the scaffold is
deliberately unchanged.

## Known non-binds

None. The page has no Objective-C or bare-configuration fragments;
every block has a fitting validator. The iOS Swift snippet is bound
to `swift-syntax-only`, whose `ios-client` harness is the macOS-only
native runner — it validates on the CI macOS row rather than locally.
