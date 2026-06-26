# Port notes: /sdk/features/identify

Source: `ld-docs-private` `fern/topics/sdk/features/identify.mdx`.
48 code blocks extracted into `sdk-docs/features/identify/`
snippets across 25 SDKs plus one shared JSON example. All but two
(iOS Objective-C and the shared JSON multi-context example) are
bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **Android identify** (`android-client-sdk/.../identify-java`,
  `identify-kotlin`): `LDContext.builderFromContext(context).email(...)`
  — `ContextBuilder` has no `email()` method (that was the legacy
  `LDUser.Builder` API). Rewritten to
  `.set("email", "sandy@example.com")`, the real attribute setter.
- **C++ client identify result, C binding**
  (`cpp-client-sdk/.../identify-result-cpp-c-v3-0`): missing `}`
  between the "Identification succeeded" branch and its `else` — a
  syntax error. Added the closing brace.
- **Flutter v4 single context** (`flutter-client-sdk/.../identify-v4`):
  `.setString('email', 'sandy@example.com'))` had a stray extra `)` —
  a syntax error. Removed it.
- **Electron identify** (`electron-client-sdk/.../identify-js`,
  `identify-ts`): the callback was passed as `identify`'s second
  argument, but the signature is `identify(context, hash, onDone)` —
  the callback landed in the secure-mode `hash` slot and would never
  fire. Also the completion callback's first parameter is the error
  (`(err, flags)`), so `(newFlags)` would have bound to the error
  value. Fixed to `client.identify(newUser, null, (err, newFlags) => ...)`.
- **Node.js (client-side) identify** (`node-client-sdk/.../identify`,
  `identify-multi`): same callback-in-the-hash-slot problem (the
  callback would never be invoked). Inserted the `null` hash argument,
  matching the page's own React sample.
- **iOS identify (Objective-C)** (`ios-client-sdk/.../identify-objc`):
  the identify calls referenced an undefined `newContext`; the builder
  result was unwrapped into a variable named `context` that
  immediately went out of scope. Declared `LDContext *newContext`
  inside the `result.success` check and moved the identify calls into
  that scope.

## Validation routing added in this port

- `android-client-sdk/scaffolds/kotlin-syntax-only` — added a
  module-scope `context: LDContext` stub (the Java scaffold already
  had one).
- `ios-client-sdk/scaffolds/swift-syntax-only` — `_wrappee()` is now
  `throws` so fragments using bare `try`
  (`try LDContextBuilder(...).build().get()`) compile without
  per-fragment error handling.
- `electron-client-sdk/scaffolds/electron-syntax-only` — added the
  `//IMPORT_LIFT_TARGET` + `//BODY_BEGIN`/`//BODY_END` marker triad
  (the js-client harness already supports it) so the TS fragment's
  top-level `import * as LDElectron ...` can be lifted to module
  scope.
- `cpp-client-sdk/scaffolds/cpp-client-syntax-only` — added an
  `IdentifyAsync` stub returning a deferred `future<bool>` (mirrors
  `StartAsync`) and an `updated_context` stub local.
- `cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-c` — added a
  `newUser` stub local; `validators/languages/cpp-client-v2-c/api.h`
  gained an inline `LDClientIdentify(client, user)` stub matching the
  real v2 C client SDK surface.
- `validators/languages/cpp-server-v2-c/api.h` — added the server
  flavor of `LDClientIdentify(client, user)`.

## Known non-binds

- `ios-client-sdk/.../identify-objc` — no Objective-C parse scaffold
  exists; the iOS validator is the macOS-only native harness (same
  blocker as the evaluating and evaluation-reasons ports' objc
  snippets). Wiring it up requires either an Objective-C target in the
  xcodegen scaffold or a clang -fsyntax-only stub harness.
- `sdk-docs/features/identify/multi-context` (`_shared`) — the page's
  generic JSON multi-context example. Two concrete blockers: there is
  no `json` validator runtime yet (a `python3 -m json.tool`-style
  parse harness would do), and `snippets validate` requires `--sdk`
  and filters on the frontmatter `sdk:` field, which `_shared`
  snippets do not carry — so no CI row can currently select it.
  Wiring it up needs both a json runner and CLI/matrix support for
  sdk-less snippets.
