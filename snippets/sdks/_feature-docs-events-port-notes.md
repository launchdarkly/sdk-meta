# Port notes: /sdk/features/events

Source: `ld-docs-private` `fern/topics/sdk/features/events.mdx`.
42 code blocks extracted into `sdk-docs/features/events/` snippets
across 26 SDKs. All but one (iOS Objective-C) are bound to
validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX.

- **Flutter v4 and v3.x** (`flutter-client-sdk/.../track-v4`, `track-v3`):
  the published samples called `LDValue.objectBuilder()`, which does
  not exist in either SDK line — the object-builder factory is
  `LDValue.buildObject()` (confirmed against the flutter-client-sdk
  source and by compiling against both the pinned 3.x package and
  4.x). Renamed the call in both blocks.
- **C++ (server) native v3.0** (`cpp-server-sdk/.../track-cpp-native-v3-0`):
  the second `client.Track(context, "example-event-key", 42)` call was
  missing its terminating semicolon — a syntax error in C++. Added it.
- **Fastly flush** (`fastly-server-sdk/.../flush`): the published block
  used bare `...` lines as elision placeholders inside
  `handleRequest`, which is not valid TypeScript (the parser reports
  "Declaration or statement expected"). Converted both to `// ...`
  comments, matching how sibling pages elide code.

## Validation routing added in this port

- `Track` variadic-template member on the `_AnyClient` stubs in both
  `cpp-client-sdk/scaffolds/cpp-client-syntax-only` and
  `cpp-server-sdk/scaffolds/cpp-syntax-only` — the stubs previously
  exposed only `TrackEvent`/`TrackData`-era names; the v3 native API
  spells the method `Track`.
- `LDClientTrack` on both v2 stub headers:
  `validators/languages/cpp-server-v2-c/api.h` (4-arg server shape:
  client, key, user, JSON data) and
  `validators/languages/cpp-client-v2-c/api.h` (2-arg client shape:
  client, key).

No new scaffolds or validators were needed; every binding reuses a
scaffold from the earlier feature-docs ports (including
`haskell-syntax-only-v3`, `flutter-syntax-only-v3`,
`cpp-syntax-only-v2-c`, `cpp-client-syntax-only-v2-c`, and the
edge-ts-backed cloudflare/fastly/vercel syntax-only scaffolds).

## Known non-binds

- `ios-client-sdk/.../track-objc` — no Objective-C parse scaffold
  exists; the iOS validator is the macOS-only native harness (same
  blocker as the evaluating port's objc snippet). Wiring it up
  requires either an Objective-C target in the xcodegen scaffold or a
  clang -fsyntax-only stub harness.
