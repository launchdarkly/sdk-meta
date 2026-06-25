# Port notes: /sdk/features/test-data-sources

Source: `ld-docs-private` `fern/topics/sdk/features/test-data-sources.mdx`.
44 code blocks extracted into `sdk-docs/features/testdata/` snippets
across 13 SDKs. All 44 are bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **Android configuration** (`android-client-sdk/.../configure-java`):
  the published block was Java with C# `using` directives, imported the
  *server* SDK packages (`com.launchdarkly.sdk.server.*` — they don't
  exist in an Android app), and passed an undefined `ldConfig` to
  `LDClient.init` after declaring the builder result as `config`.
  Rewritten with `import` + the `com.launchdarkly.sdk.android.*` /
  `...android.integrations.*` packages and `config`.
- **Android rule example** (`android-client-sdk/.../flag-behavior-java`):
  `variationFunc(context -> ... ? 1 : 0)` — the Android
  `FlagBuilder.variationFunc` only accepts a Boolean-returning
  function; the index-returning variant is `variationIndexFunc`.
  Changed to `variationIndexFunc(...)` (the lambda returns 1 or 0).
- **Java rule example** (`java-server-sdk/.../flag-behavior-v6`):
  `ifMatch(ContextAttribute.forName("admin"), ...)` — there is no
  `ContextAttribute` class in the Java SDK; the real overloads are
  `ifMatch(String attribute, LDValue...)` and
  `ifMatch(ContextKind, String, LDValue...)`. Changed to
  `ifMatch("admin", LDValue.of(true))`.
- **Go rule example** (`go-server-sdk/.../flag-behavior-v6`):
  `IfMatch(ldcontext.GetValue("admin"), ldvalue.Bool(true))` —
  `ldcontext.GetValue` is not a function (GetValue is a method on
  `Context`), and `FlagBuilder.IfMatch` takes the attribute name as a
  string. Changed to `IfMatch("admin", ldvalue.Bool(true))`.
- **Node.js configuration, both versions** (`node-server-sdk/.../configure-v8`,
  `-v7x`): `testData.update(...)` referenced an undefined variable (the
  binding is `td`), and `new LDClient('YOUR_SDK_KEY', ...)` constructs a
  class neither package exports — clients are created with the
  `init(...)` entry point. Changed to `td.update(...)` and
  `ld.init(...)` with the corresponding `require` line added.
- **Ruby configuration** (`ruby-server-sdk/.../configure`):
  `Config.new(data_source = td)` is an assignment expression, so it
  passes the TestData object itself as the positional `opts` hash and
  fails at runtime. Changed to the keyword form
  `Config.new(data_source: td)`.
- **C server SDK v2.x free** (`cpp-server-sdk/.../free-c-sdk-v2`):
  `LDTestDataFree(td)` was missing its semicolon — a syntax error in C.
- **Haskell rule examples** (`haskell-server-sdk/.../flag-behavior-v4`,
  `-v3x`): both blocks ended the `fallthroughVariation` line with `))`
  followed by a standalone `)` — one closing paren too many. Dropped
  the extra paren on the `fallthroughVariation` line.
- **Haskell configuration** (`haskell-server-sdk/.../configure`): the
  three statement lines carried a stray single-space indent in the MDX;
  flushed to column 0 (layout-significant in Haskell).

## Validation routing added in this port

- `erlang-server-sdk/scaffolds/erlang-syntax-only-block` — the existing
  `erlang-syntax-only` closes the spliced body with a bare `.`, which
  breaks on this page's fragments because every expression (including
  the last) ends with `,` — excerpts of a larger function body. The new
  variant ends the sequence with a closing `{ok, SdkKey}` expression
  and binds `SdkKey` up front (config fragments reference it as an
  ambient binding; unbound variables are compile errors in Erlang).
- v2-era C test-data stub surface: new declarations-only
  `<launchdarkly/integrations/test_data.h>` stub header in the
  `cpp-server-v2-c` validator (prototypes rather than inline no-ops so
  fragments that `#include` it from inside the wrappee function stay
  legal C), plus `LDConfigSetDataSource` / `struct LDDataSource` and
  the `LDNewArray` / `LDNewText` / `LDNewBool` / `LDArrayPush` JSON
  constructors on the stub `api.h`. The `cpp-syntax-only-v2-c`
  scaffold includes the new header at file scope and stubs an ambient
  `td`.
- Haskell do-block scaffolds (`haskell-syntax-only`, `-v3`): added the
  qualified `LD.` / `TestData.` aliases, `Data.Functor ((<&>))` and
  Data.Aeson imports the fragments use without showing, an ambient
  `td` stub, and `ExtendedDefaultRules` (so `toJSON "red"` under
  OverloadedStrings defaults to String instead of failing as
  ambiguous). The haskell-server and haskell-server-v3 validator
  images add `aeson` to the cabal build-depends to match.
- Haskell harness body-splice tweak (both haskell validators): a
  closing bracket at column 0 in a body is a continuation of the
  previous statement, but the two-space body indent used to land it on
  the do-block layout column, where GHC inserts a statement separator
  and the enclosing expression fails to parse. Such lines now indent
  four spaces. (This page's multi-line `TestData.update td =<< ( ... )`
  blocks close with `)` at column 0.)
- Rust validator image bumped `rust:1.85` --> `rust:1.94`: the
  validator project pulls the latest `launchdarkly-server-sdk`, and the
  3.x crate that carries `TestData` has an MSRV of 1.93, so the
  resolver under 1.85 picked a pre-TestData release.
- Stub-surface extensions: `td` + `startWaitTime` on
  `csharp-client-syntax-only`(+`-v3`); `td` + the
  `LaunchDarkly.Sdk.Server.Integrations` using on `csharp-syntax-only`
  (typed as the real `TestData`, not `dynamic` — C# rejects lambda
  arguments like `VariationFunc(context => ...)` in dynamically
  dispatched calls); `td` + `sdkKey` + the
  `com.launchdarkly.sdk.server.integrations` import on the Java
  `java-syntax-only`; `td` on the Android `java-syntax-only`;
  `TestData` / `FlagBuilder` imports + a `td` binding on
  `rust-syntax-only`.

## Known non-binds

None — every code block on the page is bound to a validator.
