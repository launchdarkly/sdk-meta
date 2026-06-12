# Port notes: /sdk/features/relay-proxy-config

Source: `ld-docs-private` `fern/topics/sdk/features/relay-proxy-config/`
(three pages: `index.mdx`, `proxy-mode.mdx`, `daemon-mode.mdx`).

- `index.mdx` carries no code blocks — nothing to port; it is listed
  here so the next agent knows the whole subdirectory was covered.
- `proxy-mode.mdx`: 33 code blocks extracted into
  `sdk-docs/features/relay-proxy-config/proxy-mode/` snippets across
  17 SDKs.
- `daemon-mode.mdx`: 19 code blocks extracted into
  `sdk-docs/features/relay-proxy-config/daemon-mode/` snippets across
  12 server-side SDKs.

All but one (iOS Objective-C) are bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX.

- **iOS Swift proxy mode** (`ios-client-sdk/.../proxy-mode-swift`):
  `LDConfig.streamUrl` / `baseUrl` / `eventsUrl` are non-optional
  `URL` properties, but `URL(string:)` returns `URL?` — the published
  assignments don't compile. Added `!` to each initializer (the URL
  literals are well-formed constants).
- **Flutter v4 proxy mode** (`flutter-client-sdk/.../proxy-mode-v4`):
  the second positional argument read `autoEnvAttributes.enabled` —
  no such identifier; the enum is `AutoEnvAttributes` (matches the
  validated config-port snippet for the same call).
- **C++ server C-binding proxy mode**
  (`cpp-server-sdk/.../proxy-mode-cpp-c-v3-0`): the
  `LDServerConfigBuilder_ServiceEndpoints_RelayProxyBaseURL(...)`
  statement was missing its terminating semicolon — a syntax error.
  Same fix the config port applied to its copy of this block.
- **Node.js v8 / v7 TypeScript proxy mode**
  (`node-server-sdk/.../proxy-mode-ts-v8`, `-ts-v7`): both blocks read
  `const options ld.LDOptions = { ... }` — missing the `:` of the type
  annotation, a syntax error in both TS and JS.
- **Node.js v7 TypeScript daemon mode**
  (`node-server-sdk/.../daemon-mode-ts-v7`): the published block used
  a default import (`import LDOptions from
  'launchdarkly-node-server-sdk'`), but `LDOptions` is a named
  interface export in the v7 typings — a default import binds a
  value, which cannot be used in the `const options: LDOptions` type
  annotation. Changed to the named import (same fix as the web-proxy
  port's v7 TypeScript block).
- **Haskell daemon mode** (`haskell-server-sdk/.../daemon-mode`):
  `configSetUseLdd true` — Haskell's Bool constructor is `True`;
  lowercase `true` is an unbound name.
- **Rust daemon mode** (`rust-server-sdk/.../daemon-mode`): three
  errors against the real API:
  `SomeKindOfFeatureStore.new(storeOptions)` is Ruby-style (Rust
  associated functions use `::`, and the binding that the next line
  consumes is snake_case), the `let persistent_store_factory = ...`
  statement was missing its semicolon, and `.data_store(&builder)`
  referenced an undefined `builder` — the builder declared on the
  previous line is `persistent_data_store_builder`.

## Validation routing added in this port

No new scaffolds and no new validators. Stub-surface and dependency
extensions (all additive; the Redis dependency additions mirror the
identical hunks on the in-flight big-segments port so the branches
merge cleanly):

- `dotnet-server-sdk/scaffolds/csharp-syntax-only` — added
  `LaunchDarkly.ServerSdk.Redis` to the requirements and
  `using LaunchDarkly.Sdk.Server.Integrations;` to the preamble (the
  daemon DataSystem fragment calls `Redis.DataStore()` unqualified);
  added a dynamic `SomeDatabaseName` stub for the placeholder
  database integration the DataStore fragment references.
- `java-server-sdk/scaffolds/java-syntax-only` — added
  `import com.launchdarkly.sdk.server.integrations.*;` plus stub
  `SomeDatabaseName` / `storeOptions` symbols.
- `validators/languages/jvm/harness/run.sh` — added the pinned
  `launchdarkly-java-server-sdk-redis-store` dependency to the
  synthesized pom so `Redis.dataStore()` resolves.
- `cpp-server-sdk/scaffolds/cpp-syntax-only` — added an ambient
  `sdk_key` local and a `YourDatabaseIntegration()` stub returning
  the `ISerializedDataReader` pointer `LazyLoadBuilder::Source()`
  expects.
- `cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c` — added a
  `ConstructYourFeatureStoreInterface()` placeholder-stub.
- `validators/languages/cpp-server-v2-c/api.h` — added the
  `struct LDStoreInterface` forward declaration and the
  `LDConfigSetFeatureStoreBackend` / `LDConfigSetUseLDD` setters,
  mirroring the real v2 header's daemon-mode surface.
- `haskell-server-sdk/scaffolds/haskell-config-syntax-only` — added a
  module-scope `backend = Nothing` stub (generalizes to the
  `Maybe PersistentDataStore` that `configSetStoreBackend` expects).
- `rust-server-sdk/scaffolds/rust-syntax-only` — imported
  `PersistentDataStore` / `PersistentDataStoreFactory`, added a
  `SomeKindOfFeatureStore` stub implementing
  `PersistentDataStoreFactory`, and an ambient `store_options` local.

## Known non-binds

- `ios-client-sdk/.../proxy-mode/proxy-mode-objc` — no Objective-C
  parse scaffold exists; the iOS validator is the macOS-only native
  harness (same blocker as the evaluating and evaluation-reasons
  ports' objc snippets). Wiring it up requires either an Objective-C
  target in the xcodegen scaffold or a clang -fsyntax-only stub
  harness.
