# sdk-info port notes

Issues observed while lifting `gonfalon/packages/sdk-info/src/snippets/`
into sdk-meta during Phase 1 of [SDK-2316][].
Bodies were copied verbatim; nothing here was fixed in this PR — each
entry is a flag for triage during Phase 2 (consumer refactor) or for a
follow-up content pass.

[SDK-2316]: https://launchdarkly.atlassian.net/browse/SDK-2316

## Cursor prompt template uses an unsupported placeholder syntax (fixed)

**Severity**: ~~medium~~ resolved

**SDKs affected**: cursor-prompt (root)

**What we observed**: `prompt.txt` uses `{{SDK_NAME}}`, `{{SDK_DOCS_URL}}`,
`{{SDK_EVENT_DOC_URL}}` placeholders that gonfalon's
`CursorSdkInstall.tsx` substitutes at runtime via a regex
(`/{{SDK_NAME}}/g`) that matches the no-whitespace form only. Originally
the DSL parsed any `{{ NAME }}` token (with or without inner whitespace)
into a Var node and round-tripped it as `{{ NAME }}` (with whitespace),
so the lifted body rendered byte-different from the source — gonfalon's
regex would not have matched the rendered output.

**Resolution**: The DSL now captures the original source-text form on
the `Var` node (`Raw` field) and `literalVar` emits that verbatim for
undeclared (foreign-template) names. All 52 sdk-info files now round-trip
byte-identical to the gonfalon source, including the cursor prompt's
no-whitespace placeholders. Consumers that switch their `?raw` import to
the rendered output need no further coordination.

## Inconsistent npm install package-name conventions

**Severity**: low

**SDKs affected**: react-client-sdk, vue-client-sdk, node-client-sdk
(all use `launchdarkly-<x>-sdk`); js-client-sdk, node-server-sdk,
react-native-client-sdk (use `@launchdarkly/<x>-sdk`)

**What we observed**: Newer SDK packages have moved to the
`@launchdarkly/` npm scope, but the older packages haven't, so the
install snippets diverge along package-name lines. Each command is
correct for its SDK at the time of authoring — but a Phase 2 consumer
would notice that the install-npm.txt files are nearly-but-not-quite
a single template parameterized by package name.

**js-client-sdk fix applied**: gonfalon's source had the install
snippets installing `launchdarkly-js-client-sdk` (v3) while
`init.txt` already used the v4 API (`createClient` from
`@launchdarkly/js-client-sdk`). Install commands updated in sdk-meta
to install `@launchdarkly/js-client-sdk` (npm/pnpm/yarn), and the
bower install URL bumped from `unpkg.com/launchdarkly-js-client-sdk@3`
to `unpkg.com/@launchdarkly/js-client-sdk@4` so installed code
matches the rendered example. This is a deliberate divergence from
the gonfalon source; the byte-equality round-trip test now expects
4 differing js-client-sdk install files.

**Recommended action**: When migrating consumers, consider whether
these can be expressed as a single `install-<pm>` snippet shared across
SDKs with a `package` input. Tracked under the larger
"deduplicate sdk-info with getting-started" effort in the design doc;
not urgent.

## flagEval snippets only exist for 6 of 13 SDKs

**Severity**: low

**SDKs affected**: missing for android-client-sdk, dotnet-server-sdk,
go-server-sdk, ios-client-sdk, node-client-sdk, react-native-client-sdk,
vue-client-sdk

**What we observed**: Only 6 of the 13 SDKs ship a `flagEval.txt`.
Gonfalon's FlagAddToCode UI presumably has a fallback for the missing
ones (or just hides the step). Backfilling these is a content question,
not a structural one.

**Recommended action**: Hold off on backfilling in this PR (verbatim
lift only). When Phase 2 lands the consumers, file a content-team
ticket to source flagEval examples for the 7 missing SDKs from
each SDK's reference docs.

## init.txt fragments aren't standalone runnable

**Severity**: medium

**SDKs affected**: all 13

**What we observed**: Every `init.txt` assumes a surrounding scaffold —
e.g. java-server-sdk's body has a `Main` class and a `main` method but
no package declaration; android-client-sdk's body uses `this@BaseApplication`
with no Activity context; react-client-sdk's body uses a JSX `<App>`
function with no module-level imports for `App`. The SDK-2316 design
defers validation to "reuse the scaffold model from SDK-2308," meaning
Phase 1 deliberately doesn't add `validation:` blocks to these
snippets. They're documentation fragments, not runnable units.

**Recommended action**: When the scaffold work from SDK-2308 lands,
revisit each sdk-info `init` and `flagEval` snippet and decide whether
to (a) add a scaffold-companion that wraps the fragment in a runnable
shell, or (b) leave them documentation-only and rely on the existing
`getting-started/` snippets for runnable coverage.

## Cross-source drift between sdk-info and getting-started

**Severity**: medium

**SDKs affected**: js-client-sdk, others likely

**What we observed**: js-client-sdk's `sdk-info/init.txt` uses
`createClient('YOUR_CLIENT_SIDE_ID', context)` (the v4 API), while
sdk-meta's existing `js-client-sdk/snippets/getting-started/app-ts.snippet.md`
and gonfalon's older Connect-an-SDK content reference different API
generations. The design doc flags this as the canonical
"v3 vs v4 vs docs" cross-source drift.

**Recommended action**: Phase 1 keeps the two sources side-by-side
under separate snippet groups (`sdk-info/` and `getting-started/`). The
drift is real and pre-existing; deduplication is explicitly Phase 2
scope per SDK-2316. Flag the divergence in the Phase 2 ticket so the
consumer-refactor PR resolves both consumers onto a single canonical
snippet rather than silently picking one and breaking the other.

## .NET install-csharp snippet uses a PowerShell command

**Severity**: low

**SDKs affected**: dotnet-server-sdk

**What we observed**: `install-csharp.txt` contains
`Install-Package LaunchDarkly.ServerSdk` — a NuGet PowerShell cmdlet,
not a `dotnet` CLI invocation. The `lang:` field on the lifted snippet
is therefore set to `powershell` (not `csharp`). The filename
`install-csharp.txt` is misleading.

**Recommended action**: Verbatim-lift kept the original filename so
gonfalon's existing `?raw` import keeps working. When Phase 2 migrates
the consumers, consider renaming to `install-nuget` and/or adding a
`install-dotnet-cli` companion (`dotnet add package LaunchDarkly.ServerSdk`)
for projects on the modern CLI workflow.

## install-package-swift.txt is a fragment, not a standalone Package.swift

**Severity**: low

**SDKs affected**: ios-client-sdk

**What we observed**: The body opens with `//...` and ends with `//...`,
i.e. it's intended to be pasted inside an existing `Package.swift`
manifest. Validation as-is would fail; gonfalon shows it as a
copy-paste hint.

**Recommended action**: When scaffold validation lands, mark this
snippet as documentation-only (no `validation:` block) and rely on a
separate runnable `Package.swift` companion under `getting-started/`
for end-to-end coverage.

## python-server-sdk flagEval.txt has an empty if/else body (fixed)

**Severity**: ~~medium~~ resolved

**SDKs affected**: python-server-sdk

**What we observed**: `flagEval.txt` reads:

```python
if flag_value:
    # TODO: Put your feature here
else:
    # TODO: Put your fallback behavior here
```

A bare comment is not a statement; Python rejects this with
`SyntaxError: expected an indented block after 'if' statement`. The
snippet is presented as the "Add the SDK to your code" example and is
copy-pasted by users into their own program — pasting verbatim
produces an immediate syntax error before the user's TODO has a chance
to replace the comments.

**Resolution**: Each branch was rewritten as `pass  # TODO: …` so the
fragment is syntactically complete and parses cleanly with the
`python-server-sdk/scaffolds/python-syntax-only` validator. The visible
TODO hint is preserved as a trailing comment; users still know where
to insert their feature code.

## android-client-sdk init.txt missing `AutoEnvAttributes` import (fixed)

**Severity**: ~~medium~~ resolved

**SDKs affected**: android-client-sdk

**What we observed**: `init.txt` references `AutoEnvAttributes.Enabled`
as a bare token but its imports are wildcards
(`com.launchdarkly.sdk.*` and `com.launchdarkly.sdk.android.*`) which
do not bring nested types into scope. `AutoEnvAttributes` is a nested
enum inside `LDConfig.Builder` — its fully qualified name is
`com.launchdarkly.sdk.android.LDConfig.Builder.AutoEnvAttributes` —
so kotlinc rejects the verbatim snippet with
`Unresolved reference 'AutoEnvAttributes'`. The companion hello-android
reference repo includes the explicit nested-type import
(`import com.launchdarkly.sdk.android.LDConfig.Builder.AutoEnvAttributes`),
confirming the missing import is a verbatim drift between gonfalon's
sdk-info source and the SDK's actual surface.

**Resolution**: Added the explicit nested-type import alongside the
existing wildcards. Deliberate divergence from gonfalon's
`packages/sdk-info/src/snippets/android-client-sdk/init.txt`; the
extended-validation harness now compiles and runs the body under
Robolectric end-to-end against the real LD streaming API.

## go-server-sdk init.txt has an unused import (fixed)

**Severity**: ~~medium~~ resolved

**SDKs affected**: go-server-sdk

**What we observed**: `init.txt` imports
`github.com/launchdarkly/go-sdk-common/v3/ldcontext` but never references
any symbol from that package. Go rejects unused imports as a compile
error, so following the snippet verbatim into a `main.go` produces:

```
imported and not used: "github.com/launchdarkly/go-sdk-common/v3/ldcontext"
```

**Resolution**: Unused import removed. Deliberate divergence from
gonfalon's `packages/sdk-info/src/snippets/go-server-sdk/init.txt`;
the contradiction (the snippet doesn't reference `ldcontext` but
imports it) exists in the gonfalon source today and is a correctness
bug worth correcting before extended validation runs against the
rendered output.

## js-client-sdk install-bower URL is not a valid bower target

**Severity**: low (deferred — bower is deprecated)

**SDKs affected**: js-client-sdk

**What we observed**: `install-bower.txt` reads
`bower install https://unpkg.com/@launchdarkly/js-client-sdk@4`. Bower's
URL resolver expects a tarball, zipfile, or git remote — unpkg.com
returns a `text/javascript` body for that URL, which produces:

```
ENORESTARGET URL sources can't resolve targets
```

The original gonfalon source had the same issue against the v3 unpkg
URL; bumping to v4 (this branch) didn't change the underlying
incompatibility. Bower itself has been deprecated since 2017.

**Recommended action**: Skip validation for this snippet; leave the body
unchanged so gonfalon's `?raw` import keeps shipping the canonical
fragment. When the wider consumer-refactor lands, drop the bower
install entry from the install-card surface — bower is no longer a
realistic install path.

## Validation coverage and deferred snippets

**Severity**: informational

**Context**: Extended validation was added for sdk-info snippets across
two PRs. Current coverage:

**Install (26 of 29 snippets validated end-to-end)**:

- 23 Linux-runnable installs (npm/pnpm/yarn × Node-family SDKs, pip ×
  python, go × go-server) → `validation.runtime: shell-install` runs
  the body in a clean dir and asserts the package landed.
- 3 iOS installs (Cartfile, Package.swift, Podfile) →
  `validation.runtime: ios-install` runs on macos-latest. Each kind
  (selected via `validation.env: { INSTALL_KIND: ... }`) wraps the
  fragment in a minimal stub project (synthesized Package.swift /
  Podfile + xcodeproj-generated stub project / Cartfile) and runs the
  real package manager (`swift package resolve` / `pod install` /
  `carthage update --no-build`).

**Init (13 of 13 snippets validated end-to-end)**:

- python-server, node-server, go-server, java-server, dotnet-server →
  per-SDK init-runner scaffolds with `validation.placeholders` for
  `YOUR_SDK_KEY`. EXAM-HELLO success line emitted on a real LD init.
- node-client, dotnet-client, js-client, react-client, vue-client →
  per-SDK init-runner scaffolds with `YOUR_CLIENT_SIDE_ID` /
  `YOUR_MOBILE_KEY` placeholders.
  - js-client/react-client/vue-client run in headless Chromium via
    Playwright (Vite/tsdown bundle + preview server).
  - node-client runs under Node's dynamic-import path, asserting
    `client.waitForInitialization` resolves.
  - dotnet-client wraps top-level statements in a `Microsoft.NET.Sdk`
    console project with `LaunchDarkly.ClientSdk` pulled from NuGet.
- react-native-client → existing jest+react-native-preset harness;
  init scaffold supplies a `YourComponent` that calls
  `useLDClient().waitForInitialization()` and renders the
  EXAM-HELLO sentinel on success. The harness was hardened to wait
  for jest's exit code (the original `await_success_line` matched
  jest's failure output, which echoes the regex pattern).
- ios-client → init-runner scaffold splices the body into
  `application(_:didFinishLaunchingWithOptions:)` of an
  `AppDelegate`, with a companion `ViewController` that observes
  the test flag and writes the EXAM-HELLO line into
  `featureFlagLabel`. Runs on macos-latest via the existing
  `ios-client` validator harness (xcodegen + xcodebuild test +
  iOS Simulator). The harness lifts any `import` lines out of the
  spliced body to file scope (Swift requires imports at top level)
  and uses `validation.placeholders` to substitute
  `YOUR_MOBILE_KEY` for the env-injected mobile key. CI matrix has
  separate rows for the install fragments (`key-type: none`) and
  the init snippet (`key-type: mobile`).
- android-client → init-runner scaffold splices the body into
  `MainApplication.onCreate()` (subclass of `Application`); the
  scaffold's class name is `MainApplication` and the harness
  rewrites the body's `this@BaseApplication` literal to
  `this@MainApplication` so the gonfalon-source verbatim still
  resolves. The companion `MainActivity.kt` reads the flag and
  writes the EXAM-HELLO line into `R.id.textview`. Runs under the
  existing Robolectric-based docker harness — no emulator
  required, the Android framework executes in the JVM. The
  harness also lifts in-function `import` lines (the gonfalon
  body uses wildcard imports that legally only sit at file
  scope).

**Flag-eval (6 of 6 snippets validated)**:

- python, node-server, java, js-client, dotnet-client →
  per-language syntax-only scaffolds.
- react-client → flag-eval-runner scaffold stages the body
  verbatim at `src/snippet-body.tsx`; the validator harness
  (in `SNIPPET_MODE=flag-eval` mode) reads the body, lifts
  module-scope `import` lines out of the wrappee block, computes
  `lodash.camelCase(LAUNCHDARKLY_FLAG_KEY)` to match the React
  SDK's auto-camelCase, substitutes the snippet's placeholder
  destructure target (`featureKey`) for the camelCased
  identifier, and synthesizes a fresh `App.tsx` + `main.tsx`
  that renders the wrapped body inside `<LDProvider>`. The
  existing Playwright DOM check observes the EXAM-HELLO success
  line on a true evaluation. The same docker image is reused
  for init and flag-eval modes, dispatched on
  `validation.env.SNIPPET_MODE`.

**Deferred (no current validation)**:

- Android install snippets (groovy + kotlin) — manifest fragments
  meant to be pasted into `build.gradle`; no standalone runnable
  surface.
- Java/Maven install (xml + gradle) — same shape as Android:
  dependency manifest fragments, not standalone units.
- .NET install (PowerShell `Install-Package`) — runs only under the
  legacy NuGet PowerShell host, not modern `dotnet add package`. See
  the dedicated note above.
- js-client-sdk install-bower — bower's resolver can't fetch the
  unpkg URL; documented separately above.

**Recommended action**: When the wider consumer refactor lands,
evaluate whether the manifest-fragment install snippets should ship
as full-file scaffolds with a copy-paste hint, or remain
documentation fragments outside the runnable validation surface.

>## Rendered files always end with a trailing newline

**Severity**: informational

**SDKs affected**: all

**What we observed**: The `.snippet.md` fence syntax strips the trailing
newline before the closing fence, so the rawfiles renderer adds one
back to match the POSIX-friendly convention the source `.txt` files
followed (every source file ends with a single `\n`). This is documented
in `internal/adapters/rawfiles/rawfiles.go` (`renderBody`).

**Recommended action**: None. Documented here so future readers don't
chase a phantom drift if a `.snippet.md` body ends with multiple blank
lines and the rendered output normalizes to one.
