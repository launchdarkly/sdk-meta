# Port notes: /sdk/features/multiple-environments

Source: `ld-docs-private` `fern/topics/sdk/features/multiple-environments.mdx`.
17 code blocks extracted into `sdk-docs/features/multienv/` snippets
across 3 SDKs (android-client-sdk, ios-client-sdk,
react-native-client-sdk). All but four (the iOS Objective-C blocks)
are bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **Android secondary-instance usage** (`android-client-sdk/.../get-for-mobile-key-java`,
  `get-for-mobile-key-kotlin`): `coreInstance.track("example-metric-key", data)`
  — the Android SDK has no `track(String, LDValue)` overload; the
  data-carrying variant is `trackData(String, LDValue)` (`track`
  takes only the event name). Renamed the call (and the comment
  above it) to `trackData`.
- **Android calls affecting all instances** (`android-client-sdk/.../all-instances-java`,
  `all-instances-kotlin`): `coreInstance.identify(/*Context Object*/)`
  passed a comment as the argument — `identify` has no zero-argument
  overload, so the line cannot compile. Replaced the comment
  placeholder with `context`, the ambient LDContext the page's init
  blocks build.
- **iOS calls affecting all instances** (`ios-client-sdk/.../all-instances-swift`):
  same comment-placeholder problem twice over —
  `identify(/*Context Object*/)` (no zero-argument overload; the real
  signature is `identify(context:)`) and `setOnline(/*true or false*/)`
  (`setOnline(_:)` requires a Bool). Replaced with
  `identify(context: context)` and `setOnline(true)`. The
  Objective-C sibling block keeps its placeholders verbatim — it is
  not bound (see below), and rewriting unvalidated code was out of
  scope.

## Validation routing added in this port

- `android-client-sdk/scaffolds/java-syntax-only-v4-android` — new
  android-container-routed stub scaffold for the v4-era Java init
  fragment (`new LDConfig.Builder()` has no v5 equivalent; the
  android-client container only carries the v5 aar). Nested stub
  classes declare just the v4 surface; shared `com.launchdarkly.sdk`
  types are real. Chosen over the existing jvm-routed
  `java-syntax-only-v4` because the jvm harness requires a server
  SDK key and therefore each bound snippet needs its own CI matrix
  row — the android sdk-docs row's `snippet_skip` field holds a
  single id and is already occupied by the evaluation-reasons v4
  fragment. This scaffold runs in the existing android sdk-docs row
  with no workflow changes.
- `android-client-sdk/scaffolds/kotlin-syntax-only-v4` — Kotlin
  sibling of the above for the v4-era Kotlin init fragment.
  File-scope stub classes (legal in Kotlin, unlike Java) shadow
  nothing; `LDContext` is real. Stays in the android-client
  container in parse mode.
- Stub-surface extensions: `java.util.Map` / `java.util.HashMap`
  imports, an `LDValue data` field, and a try/catch wrapper around
  the unreachable body (for `getForMobileKey`'s checked
  `LaunchDarklyException` and `close()`'s `IOException`) on
  `android-client-sdk/scaffolds/java-syntax-only`; file-scope
  `context` / `application` / `data` stubs on
  `android-client-sdk/scaffolds/kotlin-syntax-only`; a v8-era
  `LDConfig(mobileKey:)` convenience `init` (extension), ambient
  `coreInstance` / `data` stubs, and `throws` on `_wrappee()` (for
  bare `try LDContextBuilder(...).build().get()`) on
  `ios-client-sdk/scaffolds/swift-syntax-only`.

## Known non-binds

- `ios-client-sdk/.../init-v9-objc`, `init-v8-objc`,
  `get-environment-objc`, `all-instances-objc` — no Objective-C
  parse scaffold exists; the iOS validator is the macOS-only native
  harness (same blocker as the evaluating and evaluation-reasons
  ports' objc snippets). Wiring them up requires either an
  Objective-C target in the xcodegen scaffold or a clang
  -fsyntax-only stub harness.
