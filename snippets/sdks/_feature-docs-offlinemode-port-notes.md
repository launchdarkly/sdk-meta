# Port notes: /sdk/features/offline-mode

Source: `ld-docs-private` `fern/topics/sdk/features/offline-mode.mdx`.
31 code blocks extracted into `sdk-docs/features/offlinemode/`
snippets across 18 SDKs. All but one (iOS Objective-C) are bound to
validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **Android v5.x / v4.x, Java and Kotlin** (`android-client-sdk/.../offline-mode-v5-java`,
  `-v5-kotlin`, `-v4-java`, `-v4-kotlin`): two fixes in each of the
  four blocks. The builder method is `offline(boolean)` in both the
  v4 and v5 SDKs (`setOffline` is the v3-era name; the published
  blocks used `.setOffline(true)`), and the three-argument
  `LDClient.init(application, config, context)` overload returns
  `Future<LDClient>` in both v4 and v5 — assigning it to `LDClient`
  is a type error. Switched to the blocking four-argument overload
  with a `0` timeout, matching the page's own "Initialize the
  client" samples. The instance method `client.setOffline()` is real
  in both versions and kept.
- **Flutter v4** (`flutter-client-sdk/.../offline-mode-v4`): in v4
  `LDClient.start()` takes no positional arguments (config and
  context move to the `LDClient(config, context)` constructor), so
  `await client.start(config, context);` cannot compile. Replaced
  with `final client = LDClient(config, context); await
  client.start();`, which also gives the block's later
  `client.offline = true` a declared client.
- **.NET (client-side)** (`dotnet-client-sdk/.../offline-mode`): the
  .NET client SDK has no single-argument `LdClient.Init(config)`
  overload; the synchronous config-based overload is
  `Init(Configuration, Context, TimeSpan)`. Rewritten as
  `LdClient.Init(config, context, TimeSpan.FromSeconds(10))`,
  matching the evaluation-reasons sample on the same page family.
- **Java** (`java-server-sdk/.../offline-mode`): the
  `client.boolVariation("example-flag-key", context, false)`
  statement had no terminating semicolon — a syntax error.
- **Rust** (`rust-server-sdk/.../offline-mode`):
  `ConfigBuilder::build()` returns `Result<Config, BuildError>`, so
  passing its value straight to `Client::build(config)` is a type
  error. Appended `.unwrap()`, matching the SDK configuration sample
  in the config feature docs.

## Validation routing added in this port

- `android-client-sdk/scaffolds/kotlin-syntax-only-v4` — scaffold for
  v4-era Android Kotlin fragments (`LDConfig.Builder()` is 0-arg in
  v4; v5 requires `AutoEnvAttributes`). Stays in the `android-client`
  container (`SNIPPET_CHECK=parse`) with nested stub `LDConfig` /
  `LDClient` classes plus `application` / `context` ambient stubs. No
  new CI row needed — it rides the existing
  `android-client-sdk (sdk-docs)` mobile row.
- `offline-mode-v4-java` binds to `java-syntax-only-v4-android` (the
  in-container Java sibling, mobile key, `SNIPPET_CHECK=parse`) rather
  than the jvm-routed `java-syntax-only-v4`, so it needs no server-key
  row and no `snippet_skip`. `evaluation-reasons-v4-java` moved to the
  same scaffold, which let the jvm-routed `java-syntax-only-v4`
  scaffold and both dedicated server-key CI rows be removed; the
  mobile `android-client-sdk (sdk-docs)` row now validates both v4
  Java fragments directly.
- Stub-surface extensions (all additive):
  - `java-syntax-only-v4-android`: `offline(boolean)` /
    `evaluationReasons(boolean)` on the stub Builder, and a stub
    `LDClient` (`init`, `setOffline`, `boolVariationDetail`) plus
    `getApplication()` / `context` ambient stubs.
  - `kotlin-syntax-only`: file-scope `application` and `context`
    stubs (the docs assume an Activity host where the `application`
    property exists and an earlier fragment that created an
    LDContext; the scaffold splices bodies into an Application
    subclass with neither in scope).
  - `cpp-client-syntax-only` / `cpp-syntax-only` (server): file-scope
    `_AnyConfigBuilder` stub + `config_builder` local, satisfying the
    native `config_builder.Offline(true)` member call and the
    C-binding `LD{Client,Server}ConfigBuilder_Offline(config_builder,
    true)` call via implicit conversion to the opaque handle.
  - `cpp-server-v2-c` stub header: `LDConfigSetOffline`.
  - `flutter-syntax-only-v2`: `user` stub for v1.x-era bodies.
  - `haskell-syntax-only-toplevel`: module-scope `config` stub.
- `flutter-client-sdk/.../offline-mode-v1x` binds to
  `flutter-syntax-only-v2`: every v1 name it uses (single-argument
  `LDConfigBuilder`, `LDClient.start(config, user)`,
  `LDClient.setOnline`) still exists in the pinned 2.x package
  (`start` is deprecated there, which compiles cleanly), so no
  v1-pinned validator is needed for this fragment.

## Known non-binds

- `ios-client-sdk/.../offline-mode-objc` — no Objective-C parse
  scaffold exists; the iOS validator is the macOS-only native harness
  (same blocker as the evaluating and evaluation-reasons ports' objc
  snippets). Wiring it up requires either an Objective-C target in
  the xcodegen scaffold or a clang -fsyntax-only stub harness.
