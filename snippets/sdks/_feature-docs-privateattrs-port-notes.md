# Port notes: /sdk/features/private-attrs

Source: `ld-docs-private` `fern/topics/sdk/features/private-attrs.mdx`.
85 code blocks extracted into `sdk-docs/features/privateattrs/`
snippets across 23 SDK directories (the five AI-SDK sections land
under their matching server SDK directory with a `-ai` slug suffix,
mirroring how the `ai-configs` group lives under `node-server-sdk`).
All but three (the iOS Objective-C blocks) are bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0 for the indented C++ server / Apex accordions).

- **.NET (client) v4.0 config** (`dotnet-client-sdk/.../config-v4-0`):
  neither v4 nor any later release has `AllAttributesPrivate` /
  `PrivateAttributes` on `ConfigurationBuilder` — they live on
  `Components.SendEvents()` (the v4 `EventProcessorBuilder`).
  `Configuration.Builder` also has no single-argument overload in v4
  (the `AutoEnvAttributes` parameter is mandatory), and
  `LdClient.Init(Configuration, Context)` has no two-argument overload
  (a `TimeSpan` is required). Rewritten against the real v4.0.0
  surface; the duplicate `LdClient client` declaration became a
  reassignment.
- **.NET (client) v3.0 config** (`.../config-v3-0`): same
  `Events(Components.SendEvents()...)` rewrite — v3.1.0 also has no
  builder-level private-attribute methods — plus the same `Init`
  `TimeSpan` and duplicate-declaration fixes. The v3 single-argument
  `Configuration.Builder("example-mobile-key")` is real and kept.
- **Android v5 config (Java/Kotlin)** (`android-client-sdk/.../config-v5-java`,
  `-v5-kotlin`): `LDConfig ldConfig` / `val ldConfig` was declared
  twice in one scope — a compile error in both languages. Second
  declaration became a reassignment (Kotlin: first declaration is now
  `var`).
- **Android v4 config (Kotlin)** (`.../config-v4-kotlin`): same
  duplicate `val config` fix.
- **Android v4.0+ context (Java)** (`.../context-java`): the builder
  chain had no `.build()` and no statement terminator; added
  `.build();`.
- **C++ client v3.0 config (native)** (`cpp-client-sdk/.../config-cpp-native-v3-0`):
  `auto config_builder` was declared twice in one scope (redefinition).
  Second declaration became a reassignment.
- **C++ client/server v3.0 config (C binding)**
  (`cpp-client-sdk/.../config-cpp-c-v3-0`,
  `cpp-server-sdk/.../config-all-cpp-c-v3-0`): `config_builder`,
  `config`, and `config_status` were each declared twice
  (redefinitions). The second group now reassigns the existing
  variables.
- **Flutter v4 config** (`flutter-client-sdk/.../config-v4`): the
  `final config = LDConfig(...)` statement had no terminating
  semicolon; added it.
- **Flutter v3 config** (`.../config-v3`): the v3 builder method is
  `privateAttributes(Set<String>)`; there is no
  `privateAttributeNames(List)` in any 3.x release. Rewrote the call
  as `.privateAttributes({'email', 'group'})`.
- **Flutter v4 context** (`.../context-v4`): stray comma in the middle
  of the builder chain (`.kind('user', 'example-user-key'),`) — a
  syntax error. Removed.
- **Flutter v2 context** (`.../context-v2`): the
  `builder.kind(...)...privateAttributes([...])` statement had no
  terminating semicolon; added it.
- **iOS v8.x config (Swift)** (`ios-client-sdk/.../config-v8-swift`):
  `var config` was declared twice in one scope (invalid
  redeclaration). Second declaration became a reassignment.
- **JavaScript TypeScript blocks** (`js-client-sdk/.../config-all-attributes-ts`,
  `config-private-attributes-ts`, `context-ts`):
  `import { LDClient } as ld from 'launchdarkly-js-client-sdk';` is
  not valid TypeScript/ES module syntax. Replaced with
  `import * as ld from 'launchdarkly-js-client-sdk';`, matching how
  the bodies use `ld.*`.
- **Node.js (client) config (JS/TS)** (`node-client-sdk/.../config-js`,
  `config-ts`): `const client` (and in TS `const options`) was
  declared twice in one scope — a SyntaxError. Renamed the second
  pair `optionsSomePrivate` / `clientSomePrivate` (following the
  page's own `configAllPrivate` / `configSomePrivate` convention).
  The TS block also had the `ld.LDcontext` typo (real type is
  `ld.LDContext`) and called `LaunchDarkly.initialize(...)` on an
  undeclared `LaunchDarkly` binding (the import is `* as ld`).
- **Node.js (client) context+config (JS)** (`.../context-js`): the
  context object literal was missing the comma between the `email`
  member and `_meta` — a SyntaxError. Added it.
- **React Native v10 config** (`react-native-client-sdk/.../config-v10`):
  duplicate `const options` / `const client` declarations — a
  SyntaxError. Renamed the second pair `optionsSomePrivate` /
  `clientSomePrivate`.
- **Roku context** (`roku-client-sdk/.../context`): the "for an
  existing context" lines called `context.addPrivateAttribute(...)`,
  which does not exist on the SDK's context object (and never has —
  `addPrivateAttribute` is a config method; per-context private
  attributes are set via `_meta` at creation time). Removed those
  lines, keeping the `LaunchDarklyCreateContext` example.
- **.NET (server) config** (`dotnet-server-sdk/.../config`): duplicate
  `var config` / `var client` declarations (compile error). The
  second group now reassigns the existing variables. The client class
  is `LdClient`, not `LDClient` — `new LDClient(config)` does not
  compile against any server SDK release.
- **Erlang config** (`erlang-server-sdk/.../config`): the two
  `ldclient:start_instance(...)` calls were terminated with `.`
  (shell-style), which cannot compile as module code. They are now an
  expression sequence (first terminated with `,`, final terminator
  supplied by the validation scaffold).
- **Erlang context** (`.../context`): the expression ended with a
  trailing comma — a syntax error. Removed.
- **Go v6.0 context** (`go-server-sdk/.../context-v6`): the body used
  `ldvalue.ObjectBuild()` without importing `ldvalue`, and passed the
  `*ObjectBuilder` itself to `SetValue` (which takes an
  `ldvalue.Value`) — added the import and the missing `.Build()` on
  the object-builder chain.
- **Haskell v4.0 config** (`haskell-server-sdk/.../config-v4`): the
  block contained a stray Ruby line
  (`config = LaunchDarkly::Config.new(...)`), `import` lines in the
  middle of the fragment (Haskell imports must precede declarations),
  bare top-level expressions (not legal at module scope), and the
  "two attributes" example also set `configSetAllAttributesPrivate
  True`, contradicting its own comment. Removed the Ruby line and the
  redundant all-private call, hoisted the imports, and bound the two
  expressions as `configAllPrivate` / `configSomePrivate`.
- **Haskell v4.0 context** (`.../context-v4`): bare top-level
  expression; bound it as `context'` (mirroring the v3 block's
  `user'`).
- **Lua v2 context** (`lua-server-sdk/.../context-v2`): the private
  attribute was written `"email "` (trailing space inside the string),
  which would mark a nonexistent attribute private. Fixed to
  `"email"`.
- **Lua v1.x user** (`.../user-v1x`): missing comma between the
  `custom` table and `privateAttributeNames` — a syntax error. Added
  it.
- **Node.js (server) config (JS/TS)** (`node-server-sdk/.../config-js`,
  `config-ts`): duplicate `const options` declarations (and in TS a
  duplicate `import { LDOptions }` line and duplicate `const client`)
  — SyntaxErrors. Removed the duplicate import and renamed the second
  pair `optionsSomePrivate` / `clientSomePrivate`.
- **Rust config** (`rust-server-sdk/.../config`):
  `vec!["email".into(), "address".into()].into_iter().collect()`
  cannot infer a type — `private_attributes` is generic over
  `R: Into<Reference>`, and several dependency crates also implement
  `From<&str>`, so rustc reports E0283 (type annotations needed) on
  every release, 1.x included. Rewrote the elements as
  `Reference::new("email")` / `Reference::new("address")`, which pins
  the set's element type.
- **Node.js (server) v8 / v7 TS contexts and both Node AI contexts**
  (`node-server-sdk/.../context-v8-ts`, `context-v7-ts`,
  `context-ai-ts`, `context-ai-js`): `privateAttributes` was set as a
  top-level context attribute. The SDK only honors it inside the
  reserved `_meta` property (`LDContextMeta.privateAttributes`); a
  top-level `privateAttributes` is treated as an ordinary attribute
  and nothing is redacted. Moved it into `_meta`, matching the
  page's own v7 JavaScript block.

## Validation routing added in this port

- `dotnet-client-sdk/scaffolds/csharp-client-syntax-only-v4` —
  version-pinned scaffold (`LaunchDarkly.ClientSdk==4.0.0`) for
  v4.0-titled fragments, mirroring the existing v3 pin.
- `android-client-sdk/scaffolds/java-syntax-only-v4-android` and
  `android-client-sdk/scaffolds/kotlin-syntax-only-v4` — v4-era stub
  scaffolds that stay inside the `android-client` container
  (`SNIPPET_CHECK=parse`). Nested stub classes declare the v4 config
  surface (`new LDConfig.Builder()`, `Components.sendEvents()`,
  `allAttributesPrivate`, `privateAttributes`); the file does not
  import `com.launchdarkly.sdk.android`, so the v5 aar never
  collides. Unlike the jvm-routed `java-syntax-only-v4`, these need
  no extra CI matrix row or `snippet_skip` (the jvm route requires a
  server key; the existing android sdk-docs row can only skip one
  snippet id).
- `rust-server-sdk/scaffolds/rust-syntax-only-v1` — version-pinned
  scaffold for the page's "Rust SDK v1" blocks. The v1-era
  `EventProcessorBuilder` is non-generic; 2.x gave it a connector
  type parameter that cannot be inferred from
  `EventProcessorBuilder::new()` in expression position. The rust
  harness's `cargo add` now honors an optional `LD_RUST_SDK_VERSION`
  env (set via `validation.env`) so pinned scaffolds resolve the
  major they document.
- `containers` added to the haskell-server and haskell-server-v3
  validator images' cabal build-depends — the page's Haskell
  fragments import `Data.Set`, which lives in `containers`.
- Stub-surface extensions: `config` property and a v8-era
  `LDConfig(mobileKey:)` convenience `init` (extension) on
  `swift-syntax-only`; `EventProcessorBuilder` import on
  `rust-syntax-only`; `LDConfigSetAllAttributesPrivate` on the
  client v2 C stub header; `LDConfigSetAllAttributesPrivate`,
  `LDConfigAddPrivateAttribute`, `LDUserAddPrivateAttribute` on the
  server v2 C stub header.
- `haskell-syntax-only-toplevel` (+`-v3`) restructured to splice the
  body BEFORE the module-scope stubs so fragments carrying their own
  leading `import` lines stay legal (every import must precede the
  module's first declaration; declaration order itself is
  irrelevant). The non-v3 variant also gained qualified
  `Data.Set` / `LaunchDarkly.Server.Reference` imports,
  `Data.Function ((&))`, and a `sdkKey :: Text` stub.
- Flutter v1.x user fragment binds to `flutter-syntax-only-v2`: the
  2.x release still ships the full `LDUserBuilder` surface
  (`privateEmail`, `privateCustom`), and no v1-pinned validator
  exists.

## Known non-binds

- `ios-client-sdk/.../config-v9-objc`, `config-v8-objc`,
  `context-objc` — no Objective-C parse scaffold exists; the iOS
  validator is the macOS-only native harness (same blocker as the
  evaluating and evaluation-reasons ports' objc snippets). Wiring
  them up requires either an Objective-C target in the xcodegen
  scaffold or a clang -fsyntax-only stub harness.
