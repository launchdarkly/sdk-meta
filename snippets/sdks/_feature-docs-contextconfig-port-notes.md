# Port notes: /sdk/features/context-config

Source: `ld-docs-private` `fern/topics/sdk/features/context-config.mdx`.
115 code blocks extracted into `sdk-docs/features/contextconfig/`
snippets across 23 SDK directories. All but three (the iOS
Objective-C blocks) are bound to validators.

The page's five AI SDK sections (.NET AI, Go AI, Node.js AI, Python
AI, Ruby AI) demonstrate the host server SDKs' context APIs, so their
blocks live under the corresponding server SDK directories with an
`-ai` slug suffix (e.g.
`dotnet-server-sdk/sdk-docs/features/contextconfig/context-example-ai`)
— the same hosting pattern the `ai-configs` group uses.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX.

- **.NET (server) + .NET AI context example**
  (`dotnet-server-sdk/.../context-example`, `context-example-ai`):
  declared the variable as `LDContext context = Context.Builder(...)`;
  the .NET server SDK has no `LDContext` type — the context type is
  `LaunchDarkly.Sdk.Context`. Declared as `Context`.
- **C++ v3 C-binding context examples** (`cpp-client-sdk/.../context-example-cpp-c-v3-0`,
  `cpp-server-sdk/.../context-example-cpp-c-v3-0`): the firstName /
  lastName lines passed raw C strings to
  `LDContextBuilder_Attributes_Set`, whose value parameter is an
  `LDValue` (opaque pointer type). Wrapped them in
  `LDValue_NewString(...)`, matching the groups line in the same
  sample.
- **C client SDK v2.x user example** (`cpp-client-sdk/.../context-example-c-sdk-v2`):
  `LDObjectSetKey("groups", groups)` was missing the object argument —
  the real signature is `LDObjectSetKey(object, key, item)`. Now
  `LDObjectSetKey(attributes, "groups", groups)`.
- **Erlang v2 context example** (`erlang-server-sdk/.../context-example-v2`):
  the expression ended with a dangling comma, leaving an incomplete
  form. Dropped the comma.
- **Erlang v2 context kind** (`erlang-server-sdk/.../context-kind-v2`):
  the two alternative bindings (`ldclient_context:new(...)` and the
  map form) had no expression separator between them. Added the comma.
- **Erlang v2 multi-context** (`erlang-server-sdk/.../multi-context-v2`):
  the expression ended with `]). %% kind = device` — the `.`
  terminates the enclosing form mid-fragment, and the trailing
  comment after it swallows any terminator a host module supplies.
  Dropped the period and the (redundant) trailing comment; the
  preceding comment line already explains the `new/2` kind.
- **Go + Go AI multi-context** (`go-server-sdk/.../multi-context`,
  `multi-context-ai`): the multi-line `ldcontext.NewMulti(...)` call
  had no comma after the last argument; Go's semicolon insertion makes
  that a syntax error when `)` is on its own line. Added the comma.
- **Go scoped-client update** (`go-server-sdk/.../scoped-client-update`):
  the builder chain put `.Set(...)` / `.Build()` at the start of
  continuation lines; Go inserts a semicolon after `)` at line end, so
  leading-dot chaining does not parse. Moved the dots to the line
  ends.
- **Haskell v4 context example** (`haskell-server-sdk/.../context-example-v4`):
  `& withAttribute "address" $ Object $ fromList [...]` mis-associates
  (`&` binds tighter than `$`, so the context pipes into
  `withAttribute "address"` as its *value* argument), and `Object` /
  `fromList` had no imports. Parenthesized the value expression and
  added `Data.Aeson (Value (Object))` / `Data.Aeson.KeyMap (fromList)`
  imports.
- **Haskell v3 user example** (`haskell-server-sdk/.../context-example-v3`):
  the setters were written as `setFirstName "Sandy"` etc.; the v3
  `LaunchDarkly.Server.User` module exports `userSetFirstName` /
  `userSetLastName` / `userSetEmail`, each taking `Maybe Text`.
  Renamed and wrapped the values in `Just`.
- **React Native context example** (`react-native-client-sdk/.../context-example`):
  the import line was missing `from`
  (`import { type LDContext } '@launchdarkly/...'`). Added it.
- **Node.js (server) TypeScript context kind** (`node-server-sdk/.../context-kind-ts`):
  the fence was tagged `js` but the body is TypeScript
  (`const context: ld.LDContext = ...`); the sibling block in the same
  accordion is the JavaScript flavor. Retagged as `ts` (the MDX fence
  tag is updated in the docs-side marker PR).
- **JS (browser) context-kind + multi-context**
  (`js-client-sdk/.../context-kind`, `.../multi-context`): the MDX
  blocks appended `LDClient.initialize(...)` plus a
  `waitForInitialization` try/catch after the context object. These
  are context-shape reference snippets; every other SDK's context-kind
  / multi-context block shows the context object only. Trimmed the
  client-init and error-handling lines so the two js-client blocks
  match their siblings (the docs-side render reflects the shorter
  body).

## Validation routing added in this port

- `haskell-server-sdk/scaffolds/haskell-config-syntax-only-v3` — v3
  flavor of the complete-module scaffold (`{{ body }}` + appended
  `main`), routed through the `haskell-server-v3` validator, for the
  v3.x user example (pragma + `LaunchDarkly.Server.User` imports +
  top-level bindings).
- `haskell-server-sdk/scaffolds/haskell-syntax-only-expression` —
  binds a bare `Context` expression fragment (`makeContext ...`,
  `makeMultiContext [...]`) to a module-scope `_fragment :: Context`
  so GHC can type-check it; neither the do-block scaffold (non-IO
  statement) nor module scope (bare expression) accepts these shapes.
- `aeson` added to the haskell-server validator's pre-baked cabal
  `build-depends` (already a transitive dep of the SDK, so no extra
  compile cost) so bodies importing `Data.Aeson` / `Data.Aeson.KeyMap`
  resolve.
- IMPORT_LIFT markers (`//IMPORT_LIFT_TARGET` + `//BODY_BEGIN` /
  `//BODY_END`) added to `electron-client-sdk/scaffolds/electron-syntax-only`
  so the Electron TypeScript example's top-level `import * as
  LDElectron ...` lifts to module scope (the js-client harness's
  existing awk pre-step keys on the markers; same extension
  js-syntax-only received earlier).
- `ios-client-sdk/scaffolds/swift-syntax-only`: `_wrappee()` is now
  `throws` so the multi-context sample's bare `try` calls compile
  without a do/catch wrapper. Additive — non-throwing bodies are
  unaffected.
- Stub-surface extensions on the v2 C stub headers:
  - `validators/languages/cpp-client-v2-c/api.h`: `LDUserSetFirstName`,
    `LDUserSetLastName`, `LDUserSetCustomAttributesJSON`, plus the
    LDJSON builders (`LDNewObject`, `LDNewArray`, `LDNewText`,
    `LDArrayPush`, `LDObjectSetKey`).
  - `validators/languages/cpp-server-v2-c/api.h`: `LDUserSetCustom`
    plus the same LDJSON builders.
- `<launchdarkly/bindings/c/array_builder.h>` included at file scope
  in `cpp-syntax-only` / `cpp-client-syntax-only` so the C-binding
  example's in-body `#include` of that header is satisfied by the
  include guard (a first include inside the wrappee's function body
  would open an `extern "C"` block at function scope, which is
  invalid C++).

No CI workflow changes: every touched SDK already has a matrix row
whose filters pick up the new `sdk-docs/features/contextconfig/*`
ids, and the new scaffolds route through existing validators.

## Known non-binds

- `ios-client-sdk/.../context-example-objc`, `context-kind-objc`,
  `multi-context-objc` — no Objective-C parse scaffold exists; the
  iOS validator is the macOS-only native harness (same blocker as the
  evaluating and evaluation-reasons ports' objc snippets). Wiring them
  up requires either an Objective-C target in the xcodegen scaffold or
  a clang -fsyntax-only stub harness.
