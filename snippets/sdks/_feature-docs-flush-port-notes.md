# Port notes: /sdk/features/flush

Source: `ld-docs-private` `fern/topics/sdk/features/flush.mdx`.
36 code blocks extracted into `sdk-docs/features/flush/` snippets
across 23 SDKs. All but one (iOS Objective-C) are bound to
validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX.

- **C++ (client-side) flush result, native**
  (`cpp-client-sdk/.../flush-result-cpp-native-v3-0`): the published
  block was the identify-feature "examine the result" example with
  `IdentifyAsync` renamed to `FlushAsync` — but `Client::FlushAsync()`
  takes no arguments and returns `void`, so there is no future to
  `wait_for` and no result to examine. Rewritten to the real
  fire-and-forget call with a comment stating that no result is
  reported.
- **C++ (client-side) flush result, C binding**
  (`cpp-client-sdk/.../flush-result-cpp-c-v3-0`): same identify-derived
  block (`LDClientSDK_Flush(client, context, maxwait,
  &flushed_successfully)`); the real `LDClientSDK_Flush(LDClientSDK,
  unsigned int)` takes exactly two arguments, must be passed
  `LD_NONBLOCKING`, and returns `void`. The published block also had a
  missing `}` before its first `else`. Rewritten to the real
  fire-and-forget call with a comment.
- **PHP flush** (`php-server-sdk/.../flush`): `ldclient->flush();` is
  missing the `$` variable sigil — as written it dereferences an
  undefined constant. Fixed to `$ldclient->flush();`.
- **Rust synchronous flush** (`rust-server-sdk/.../flush-blocking`):
  `Client::flush_blocking` is `async` (the SDK's own rustdoc example
  awaits it); without `.await` the returned future is never polled and
  `if !success` does not compile. Added `.await`.
- **Fastly flush** (`fastly-server-sdk/.../flush`): the body used bare
  `...` lines as elision placeholders, which are TypeScript syntax
  errors. Converted to `// ...` comments.

## Validation routing added in this port

- `FlushAsync` stub added to the `_AnyClient` parse-only stubs on
  `cpp-client-sdk/scaffolds/cpp-client-syntax-only` and
  `cpp-server-sdk/scaffolds/cpp-syntax-only` (void-returning, matching
  the real client surface).
- `LDClientFlush` added to the v2 stub headers
  `validators/languages/cpp-client-v2-c/api.h` and
  `validators/languages/cpp-server-v2-c/api.h`.
- The rust validator base image moved from `rust:1.85` to `rust:1.94`:
  cargo's MSRV-aware resolver was pinning `launchdarkly-server-sdk` to
  2.6.2 (the newest release whose `rust-version` fit the 1.85
  toolchain), so the 3.0 `flush_blocking` API could not compile. Every
  rust snippet was re-validated against the 3.x line after the bump.

No new scaffolds, validators, or CI matrix rows were needed; every
bound snippet routes through a scaffold and matrix row that already
exists.

## Known non-binds

- `ios-client-sdk/.../flush-objc` — no Objective-C parse scaffold
  exists; the iOS validator is the macOS-only native harness (same
  blocker as the evaluating and evaluation-reasons ports' objc
  snippets). Wiring it up requires either an Objective-C target in the
  xcodegen scaffold or a clang -fsyntax-only stub harness.
