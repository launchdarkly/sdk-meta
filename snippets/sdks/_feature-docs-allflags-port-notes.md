# Port notes: /sdk/features/all-flags

Source: `ld-docs-private` `fern/topics/sdk/features/all-flags.mdx`.
40 code blocks extracted into `sdk-docs/features/allflags/` snippets
across 25 SDKs. All but one (iOS Objective-C) are bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (including incidental
formatting like doubled spaces and trailing comment lines).

- **C++ (server-side) v3.0 native** (`cpp-server-sdk/.../allflags-cpp-native-v3-0`):
  `if (all_flags.Valid() {` was missing the closing parenthesis on the
  condition — a syntax error. Now `if (all_flags.Valid()) {`.
- **C++ (client-side) v3.0 C binding** (`cpp-client-sdk/.../allflags-cpp-c-v3-0`):
  the iterator constructor was written as `LDValue_CreateObjectIter`,
  which has never existed in the C binding; the real function (and the
  one the SDK's own header documentation uses) is
  `LDValue_ObjectIter_New`.
- **iOS Swift** (`ios-client-sdk/.../allflags-swift`): `LDClient.allFlags`
  is declared `[LDFlagKey: LDValue]?` (nil when the client is not
  started — exactly what the page's prose says), so assigning it to a
  non-optional `[String: LDValue]` cannot compile. Declared the result
  as `[String: LDValue]?`.
- **Flutter v3.x** (`flutter-client-sdk/.../allflags-v3`): the v3 SDK's
  `LDClient` is an all-static API (the instance client arrived in v4),
  so `await client.allFlags()` has no `client` to call through.
  Rewritten as `await LDClient.allFlags()`, matching the v3 examples on
  the rest of the page family.

## Validation routing added in this port

No new scaffolds or validators; all 39 bound snippets route through
scaffolds that already exist. Stub-surface extensions (additive):

- `cpp-client-sdk/scaffolds/cpp-client-syntax-only` — `_AnyClient`'s
  `AllFlags` stub now returns
  `std::unordered_map<std::string, launchdarkly::Value>` (mirroring the
  real client) so the range-for + structured-bindings fragment
  compiles.
- `cpp-server-sdk/scaffolds/cpp-syntax-only` — `_AnyClient`'s
  `AllFlagsState` stub now returns a default-constructed (invalid)
  `launchdarkly::server_side::AllFlagsState` so `.Valid()` /
  `.Values()` chains compile.
- `validators/languages/cpp-client-v2-c/api.h` — `LDAllFlags(client)`
  plus the shared LDJSON collection helpers (`LDGetIter`, `LDIterNext`,
  `LDIterKey`, `LDJSONSerialize`, `LDFree`, `LDJSONFree`) and
  `<stdio.h>` for the fragment's `printf`.
- `validators/languages/cpp-client-v2-cpp/api.hpp` — `getAllFlags()` on
  `LDClientCPP`, the same LDJSON collection helpers, and `<iostream>`
  for the fragment's `std::cout`.
- `validators/languages/cpp-server-v2-c/api.h` —
  `LDAllFlags(client, user)` (the server v2 two-argument form).
- `rust-server-sdk/scaffolds/rust-syntax-only` — `FlagDetailConfig`
  added to the prelude import; `ldclient` stubbed alongside `client`
  (the all-flags fragment names the ambient client `ldclient`).

## Known non-binds

- `ios-client-sdk/.../allflags-objc` — no Objective-C parse scaffold
  exists; the iOS validator is the macOS-only native harness (same
  blocker as the evaluating and evaluation-reasons ports' objc
  snippets). Wiring it up requires either an Objective-C target in the
  xcodegen scaffold or a clang -fsyntax-only stub harness.
