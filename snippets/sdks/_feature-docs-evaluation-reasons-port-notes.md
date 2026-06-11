# Port notes: /sdk/features/evaluation-reasons

Source: `ld-docs-private` `fern/topics/sdk/features/evaluation-reasons.mdx`.
55 code blocks extracted into `sdk-docs/features/evaluation-reasons/`
snippets across 24 SDKs. All but one (iOS Objective-C) are bound to
validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **.NET (server) reason inspection** (`dotnet-server-sdk/.../print-reason`):
  the published sample was Java translated to C# — `EvaluationReason`
  has no nested `RuleMatch` / `PrerequisiteFailed` / `Error` classes to
  cast to, enum members are PascalCase (`EvaluationReasonKind.Off`),
  and one line still read `System.out.println(...)`. Rewritten against
  the real flat `EvaluationReason` API (`reason.RuleIndex`, etc.).
- **Java reason inspection** (`java-server-sdk/.../print-reason`): the
  nested-class casts (`(EvaluationReason.RuleMatch) reason`) are from
  the v4-era API; since v5 `EvaluationReason` is a single flat class.
  Rewritten to the flat getters (`reason.getRuleIndex()`, ...).
- **Android reason inspection** (`android-client-sdk/.../print-reason-java`):
  same nested-class-era code; same flat-getter rewrite. Timber calls
  kept (the validator's gradle project now includes Timber).
- **Python reason inspection** (`python-server-sdk/.../print-reason`):
  Python 2 `print` statements — a syntax error on any supported
  Python. Converted to `print(...)` calls.
- **Go reason inspection** (`go-server-sdk/.../print-reason`): the
  switch referenced an undefined variable `r`; the parameter is
  `reason`.
- **JavaScript SDK v4.x example** (`js-client-sdk/.../evaluation-reasons-v4`):
  the published block was initialization-status boilerplate that never
  called a `*variationDetail` method. Replaced with a v4
  `boolVariationDetail` example paralleling the v3.x block.
- **Flutter v4 example** (`flutter-client-sdk/.../evaluation-reasons-v4`):
  `detail.variationIndex` and `detail.reason` are nullable in v4;
  declared as `int?` / `LDEvaluationReason?` so the sample compiles.
- **Flutter reason inspection** (`flutter-client-sdk/.../print-reason`):
  `LDKind` members are lowerCamelCase in v4 (`LDKind.off`), not
  SCREAMING_CASE, and there is no `LDKind.UNKNOWN`; rewritten
  accordingly (exhaustive switch, no default).
- **Lua examples** (`lua-server-sdk/.../evaluation-reasons-v2`, `-v1x`):
  the colon call passed `client` again as the first argument
  (`client:boolVariationDetail(client, context, ...)` — colon syntax
  already supplies self), and `details.reason == "FLAG_NOT_FOUND"`
  compared a table to a string. Fixed to
  `client:boolVariationDetail(context, "example-flag-key", false)` and
  `details.reason.kind == "ERROR" and details.reason.errorKind == "FLAG_NOT_FOUND"`,
  matching the shape `LuaPushDetails` actually returns.
- **C client SDK v2.x examples** (`cpp-client-sdk/.../evaluation-reasons-c-sdk-v2`,
  `-c-sdk-v2-cpp`): `details->reason` on a stack struct (must be
  `details.reason`) and `LDFreeDetailContents(&details)` (the real v2
  header takes the struct by value).

## Validation routing added in this port

- `android-client-sdk/scaffolds/java-syntax-only-v4` — jvm-routed stub
  scaffold for v4-era Android fragments (`new LDConfig.Builder()` has
  no v5 equivalent; the android-client container only carries the v5
  aar). Follows the C-v2 stub-header approach: only the
  `com.launchdarkly.sdk.android` surface is stubbed, shared
  `com.launchdarkly.sdk` types are real.
- `android-client-sdk/scaffolds/java-syntax-only-members`,
  `java-server-sdk/scaffolds/java-syntax-only-members`,
  `dotnet-server-sdk/scaffolds/csharp-syntax-only-members` —
  class-member-scope variants for fragments that are method
  declarations (no local methods in Java/C#).
- `haskell-server-sdk/scaffolds/haskell-syntax-only-toplevel` (+`-v3`) —
  module-scope splice with module-scope `client`/`context` (v3:
  `user`) stubs, for fragments that are top-level bindings.
- `cpp-server-sdk/scaffolds/cpp-syntax-only-v2-cpp` + the
  `cpp-server-v2-cpp` validator — g++ pass over the same stub
  `<launchdarkly/api.h>` as `cpp-server-v2-c`, for the v2 server
  fragments written as C++.
- Stub-surface extensions: detail-variation members on the v2 C/C++
  stub headers and the v3 `_AnyClient` stubs; `LDDetails` /
  `LDEvalReason` on the server v2 stub header (variation signatures
  corrected from `struct LDJSON **` to the real `struct LDDetails *`);
  `context`/`secondsToBlock` fields on the android `java-syntax-only`
  scaffold; `myContext` on `csharp-syntax-only`; `ldConfig`/`context`
  on `swift-syntax-only`; `Reason` import on `rust-syntax-only`;
  Timber dependency in the android-client validator image.

## Known non-binds

- `ios-client-sdk/.../evaluation-reasons-objc` — no Objective-C parse
  scaffold exists; the iOS validator is the macOS-only native harness
  (same blocker as the evaluating port's objc snippet). Wiring it up
  requires either an Objective-C target in the xcodegen scaffold or a
  clang -fsyntax-only stub harness.
