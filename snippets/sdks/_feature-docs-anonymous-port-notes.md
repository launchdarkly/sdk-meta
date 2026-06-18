# Port notes: /sdk/features/anonymous

Source: `ld-docs-private` `fern/topics/sdk/features/anonymous.mdx`.
64 code blocks extracted into `sdk-docs/features/anonymous/`
snippets across 27 SDK directories. All but two (the iOS
Objective-C pair) are bound to validators.

## AI SDK sections

The page's five AI SDK sections (.NET AI, Go AI, Node.js AI,
Python AI, Ruby AI) show the base server SDK's context API, and
only the Node.js AI SDK exists in sdk-meta's product metadata, so
these blocks are hosted under the corresponding base SDK
directories with `-ai` slugs (e.g.
`go-server-sdk/sdk-docs/features/anonymous/anonymous-ai`) rather
than new `sdks/<id>/` directories. They validate through the base
SDK's existing scaffolds and CI rows.

## Content corrections

Fixes applied where the published sample could not work as
written. Everything else is verbatim from the MDX.

- **C client SDK v2.x examples** (`cpp-client-sdk/.../anonymous-c-sdk-v2`,
  `-c-sdk-v2-cpp`): both blocks declared `user` twice in the same
  scope (`struct LDUser *user = LDUserNew("example-user-key")`
  followed by `struct LDUser *user = LDUserNew(NULL)`) — a
  redefinition error in C and C++. The second declaration is
  renamed `anonymousUser`.
- **iOS Swift** (`ios-client-sdk/.../anonymous-swift`):
  `contextBuilder.build().get()` calls the throwing
  `Result.get()` without `try`, which does not compile; the
  sibling auto-generated-key block already shows
  `try ... .build().get()`. Added `try`.
- **Erlang v1.x** (`erlang-server-sdk/.../anonymous-v1x`): map
  literal had a trailing comma (`anonymous => Anonymous,}`), a
  syntax error in Erlang. Removed (same fix as the evaluating
  port's v1.x block).
- **Haskell v3.x** (`haskell-server-sdk/.../anonymous-v3`): the
  binding was written ``user` = userSetAnonymous True user`` with
  a backtick — invalid Haskell (backticks delimit infix
  operators). The intended prime identifier `user'` is restored.
- **C++ (server-side) native v3.0**
  (`cpp-server-sdk/.../anonymous-cpp-native-v3-0`): the MDX fence
  was tagged `c` but the body is the C++ native API
  (`auto context = ContextBuilder()...`); the snippet is tagged
  `lang: cpp` and the docs-side fence tag is corrected to `cpp`
  in the marker pass (the client-side twin of this block is
  already tagged `cpp` on the page).

## Validation routing added in this port

- `haskell-server-sdk/scaffolds/haskell-syntax-only-expr` — binds
  a bare pure expression body (`makeContext ... & withAnonymous
  True`) to a module-level name. The do-block scaffold rejects
  non-IO expression statements and the toplevel variants need
  declaration-shaped bodies, so neither fit.
- `ios-client-sdk/scaffolds/swift-syntax-only` — `_wrappee()` is
  now `throws` so bodies containing `try` expressions type-check
  without their own do/catch.
- Stub-surface extensions: `LDUserSetAnonymous` added to the
  cpp-client v2 stub headers (`validators/languages/cpp-client-v2-c/api.h`,
  `cpp-client-v2-cpp/api.hpp`); the server v2 stubs already had it.

## Known non-binds

- `ios-client-sdk/.../anonymous-objc` and `.../anonymous-autogen-objc`
  — no Objective-C parse scaffold exists; the iOS validator is the
  macOS-only native harness (same blocker as the evaluating and
  evaluation-reasons ports' objc snippets). Wiring them up requires
  either an Objective-C target in the xcodegen scaffold or a
  clang -fsyntax-only stub harness.
