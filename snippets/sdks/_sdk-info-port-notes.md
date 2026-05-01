# sdk-info port notes

Issues observed while lifting `gonfalon/packages/sdk-info/src/snippets/`
into sdk-meta during Phase 1 of [SDK-2316][].
Bodies were copied verbatim; nothing here was fixed in this PR — each
entry is a flag for triage during Phase 2 (consumer refactor) or for a
follow-up content pass.

[SDK-2316]: https://launchdarkly.atlassian.net/browse/SDK-2316

## Cursor prompt template uses an unsupported placeholder syntax

**Severity**: medium

**SDKs affected**: cursor-prompt (root)

**What we observed**: `prompt.txt` uses `{{SDK_NAME}}`, `{{SDK_DOCS_URL}}`,
`{{SDK_EVENT_DOC_URL}}` placeholders that gonfalon's
`CursorSdkInstall.tsx` substitutes at runtime via a string-replace.
sdk-meta's templating DSL parses any `{{ NAME }}` token (with or without
inner whitespace) into a Var node and round-trips it as `{{ NAME }}`
(with whitespace) when no input is declared. The lifted body therefore
renders byte-different from the source file: `{{SDK_NAME}}` becomes
`{{ SDK_NAME }}`. This is the only file in the 52-file set that does not
round-trip byte-identical.

**Recommended action**: Phase 2 (consumer refactor) replaces the
runtime string-replace with the canonical DSL. Two paths to consider:
either teach the DSL a "passthrough" mode that emits Var nodes verbatim
(preserving original whitespace), or migrate the cursor template to use
the DSL's own input contract and let the `<Snippet>` component supply
the values. Until then, gonfalon's runtime substitution must continue
to accept both `{{SDK_NAME}}` and `{{ SDK_NAME }}` so the canonical
sdk-meta body stays usable.

## Inconsistent npm install package-name conventions

**Severity**: low

**SDKs affected**: js-client-sdk, react-client-sdk, vue-client-sdk,
node-client-sdk (all use `launchdarkly-<x>-sdk`); node-server-sdk,
react-native-client-sdk (use `@launchdarkly/<x>-sdk`)

**What we observed**: Newer SDK packages have moved to the
`@launchdarkly/` npm scope, but the older packages haven't, so the
install snippets diverge along package-name lines. This isn't a bug —
each command is correct for its SDK — but a Phase 2 consumer would
notice that the install-npm.txt files are nearly-but-not-quite a single
template parameterized by package name.

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

## Rendered files always end with a trailing newline

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
