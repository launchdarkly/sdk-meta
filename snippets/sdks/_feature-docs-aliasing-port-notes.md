# Port notes: /sdk/features/aliasing-users

Source: `ld-docs-private` `fern/topics/sdk/features/aliasing-users.mdx`.
25 code blocks extracted into `sdk-docs/features/aliasing/` snippets
across 23 SDKs. All but one (iOS Objective-C) are bound to validators.

The page documents the legacy `alias` API, which most SDKs removed in
the major version that introduced contexts. Every snippet body is the
legacy call, kept verbatim; the version each sample applies to is
recorded in the snippet's `description` (and mirrors the page's
per-SDK accordion title). Where the current SDK no longer exposes the
API, validation goes through version-pinned scaffolds (haskell v3,
dotnet-client v3, the C v2 stub headers) or through self-contained /
additive stubs — the snippets were not rewritten to a replacement API
(none exists; the page is explicitly about the removed feature).
Electron and Apex are the exceptions where the API is still current.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX.

- **Apex alias** (`apex-server-sdk/.../aliasing/alias`): the statement
  had no terminating semicolon, which Apex requires — the
  `apex-anonymous` parser rejects it. The method itself is real:
  `LDClient.alias(LDUser currentUser, LDUser previousUser)` in the
  apex-server-sdk repo. Added the `;`.

## Validation routing added in this port

- `flutter-client-sdk/scaffolds/flutter-syntax-only-v1` — the v1.x
  `LDClient.alias` is a static method removed at 2.0, and the 1.x
  package predates Dart null safety so no validator container can pin
  it. The scaffold routes through the current `flutter-client`
  container with a self-contained static-`alias` stub class (the SDK
  package is deliberately not imported).
- `java-server-sdk/scaffolds/java-syntax-only-v5` — jvm-routed stub
  scaffold for the v5-era `client.alias(user, previousUser)` fragment
  (removed at 6.0; the jvm validator's pom pins the v7 jar). Same
  nested-stub approach as the android `java-syntax-only-v4` scaffold.
- `android-client-sdk/scaffolds/java-syntax-only-v3` — stub scaffold
  for the v3-era Android Java alias fragment (removed at 4.0). Routed
  through the `android-client` container's parse path rather than the
  `jvm` validator: the jvm harness requires a server-side SDK key that
  the android CI row does not provision (that constraint is why the
  existing v4 jvm binding needed its own matrix row), while the
  android parse path needs no key. The stubs are self-contained, so
  the real v5 aar surface is never touched.
- `erlang-server-sdk/scaffolds/erlang-syntax-only-users` — variant of
  `erlang-syntax-only` that pre-binds the ambient `User` /
  `PreviousUser` variables the v1.x fragment assumes (erlc rejects
  unbound variables, so the plain scaffold cannot host the body).
- Stub-surface extensions:
  - `kotlin-syntax-only`: an `LDClient.alias` extension-function stub
    plus ambient `newUser` / `previousUser` vals — Kotlin extension
    functions let the legacy method resolve against the real v5 aar.
  - `swift-syntax-only`: an `LDClient.alias(context:previousContext:)`
    extension stub plus `newUser` / `previousUser` properties beside
    the existing `context` stub.
  - `rust-syntax-only`: a `BetaAliasExt` extension trait (the beta
    `alias` was dropped at 1.0 with the rest of the User API the
    scaffold already stubs) plus `user` / `previous_user` bindings in
    `_wrappee`.
  - `haskell-syntax-only-v3`: `newUser` / `previousUser` do-block
    stubs (the pinned 3.1.1 SDK still exports the real `alias`).
  - `csharp-syntax-only` and `csharp-client-syntax-only-v3`:
    `newUser` / `previousUser` ambient stubs (`client` is `dynamic`
    in both, so the legacy `Alias` call resolves at parse time).
  - `LDClientAlias` on both legacy C stub headers
    (`cpp-client-v2-c/api.h`, `cpp-server-v2-c/api.h` — the latter is
    shared by the `cpp-server-v2-cpp` validator) plus `newUser` /
    `previousUser` file-scope stubs on the two `*-syntax-only-v2-c`
    scaffolds.

## Known non-binds

- `ios-client-sdk/.../aliasing/alias-objc` — no Objective-C parse
  scaffold exists; the iOS validator is the macOS-only native harness
  (same blocker as the evaluating and evaluation-reasons ports' objc
  snippets). Wiring it up requires either an Objective-C target in the
  xcodegen scaffold or a clang -fsyntax-only stub harness.
