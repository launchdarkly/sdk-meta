# experimentation port notes

Notes for the `experimentation/` snippet group added for the frontend
client SDKs: `react-client-sdk` (React Web), `js-client-sdk`,
`react-native-client-sdk`, `ios-client-sdk`, and `android-client-sdk`.

Each SDK gets two snippets — `track-only` and `full` — ported from the
"Proposed onboarding SDK snippets for experimentation" Confluence page
(pageId 4880073181). The reviewer comments left on that page are applied
inline rather than carried as TODOs; the substantive ones are recorded
below for traceability.

This file is the analogue of `_aiconfigs-observability-port-notes.md`,
`_sdk-info-port-notes.md`, and `_sdk-docs-port-notes.md`.

## Validation

The 8 frontend-web + mobile-RN + iOS experimentation snippets bind
through the existing per-SDK `*-syntax-only` scaffolds in parse-only
mode. Each scaffold wraps the wrappee body inside a never-invoked
function under an `if (false)` (or equivalent) guard, so unresolved
caller surfaces (`<YourComponent>`, the host app's `Activity`/
`AppDelegate` lifecycle, etc.) don't have to actually exist — the
validator only asserts that the body parses (and, where the toolchain
type-checks too, that the body type-checks).

Per-SDK summary:

- **js-client-sdk** (`experimentation/full`, `experimentation/track-only`)
  — bind to `scaffolds/js-syntax-only`. The harness's existing
  `IMPORT_LIFT_TARGET` / `BODY_BEGIN` / `BODY_END` awk pre-step lifts
  top-level `import`s and strips `export` keywords so the body's
  module-scope-only directives are legal inside the scaffold's
  `async function _wrappee()` wrapper. The full variant's top-level
  `await ldClient.start()` works because `_wrappee` is async. No
  scaffold or harness changes needed.

- **react-client-sdk** (`experimentation/full`, `experimentation/track-only`)
  — bind to `scaffolds/react-syntax-only`. `@launchdarkly/react-sdk`
  is already pre-baked in the validator image. The same IMPORT_LIFT
  awk pre-step handles the body's `import` directives. The full
  variant's hook calls (`useLDClient`, `useFlags`) live inside a
  function component declared in the body; that compiles fine as a
  dead-code function declaration under `if (false)`.

- **react-native-client-sdk** (`experimentation/full`, `experimentation/track-only`)
  — bind to `scaffolds/react-native-syntax-only` (Babel-parse pass
  via `SNIPPET_MODE=syntax-only`). The scaffold previously hard-coded
  `import React from 'react';` at file scope; both wrappee bodies
  import React themselves, which would have redeclared it after the
  IMPORT_LIFT step. The scaffold's React import was dropped (the
  scaffold's `App` returns `null` and uses no JSX, so the import was
  unused). `<YourComponent />` in the bodies parses fine — Babel
  doesn't resolve component references.

- **ios-client-sdk** (`experimentation/full`, `experimentation/track-only`)
  — bind to `scaffolds/swift-syntax-only`. The bodies declare
  file-scope `func startLaunchDarkly()`, `func onUserBecomesEligible(_:)`,
  `func trackMetric(_:_:)`; spliced into the scaffold they become
  nested local functions inside `AppDelegate._wrappee()`, which
  Swift permits. The full variant calls `applyVariant(variant)` from
  inside `onUserBecomesEligible`; the scaffold was extended with a
  file-scope `func applyVariant(_ variant: String) {}` no-op stub so
  xcodebuild's type-checker resolves the reference. The harness's
  existing Python import-lift pre-step moves the body's
  `import LaunchDarkly` up to file scope.

## Still Bucket C: android-client-sdk

**Snippets affected**: `android-client-sdk/experimentation/{track-only,full}`.

**Why unbindable today**: the existing
`android-client-sdk/scaffolds/android-syntax-only` scaffold routes
through the `jvm` validator (Java + Maven against
`launchdarkly-java-server-sdk`); the experimentation bodies are
Kotlin and reference `com.launchdarkly.sdk.android.*` types (the
android client SDK ships as an `aar` to Google's Maven, not a plain
jar to Maven Central), plus AndroidX types like `AppCompatActivity`,
`Application`, and `Bundle`. There is no parse-only mode of the
`android-client` validator (the existing harness drives a
MainActivity through a Robolectric lifecycle and asserts on a
TextView), and adding one would require a kotlinc-only /
`compileDebugKotlin` dispatch path plus a kotlin-aware syntax-only
scaffold — larger than this slice. Same structural gap as the
`android-client-sdk/sdk-docs/*` fragments documented in
`_sdk-docs-port-notes.md`.

## Reviewer comments applied inline

Deliberate changes from the raw page text, per the page's review
comments:

- **React Web** — unified both variants on the current
  `@launchdarkly/react-sdk` (`createLDReactProvider`, `useLDClient`,
  `useFlags`). The page's `track-only` still used the legacy
  `launchdarkly-react-client-sdk` / `asyncWithLDProvider`; the author
  comment ("Updated this snippet to use the latest version") only
  reached the `full` variant. Matches the canonical
  `react-client-sdk/sdk-info/init`.
- **JavaScript** — replaced `start()` + `await waitForInitialization()`
  with `await start()` ("We can just await start. You only need
  waitForInitialization if you are needing to wait somewhere where you
  aren't starting.").
- **JavaScript / React Native** — renamed `setExperimentContext` to
  `identifyUser` and reframed the guidance to prefer initializing
  directly with a known deterministic key, using an anonymous context
  only when the key is unknown and reusing it on transition ("I find
  setExperimentContext very confusing… if you already have your
  deterministic key, it usually makes sense to initialize directly with
  that and remove the opportunity for mistakes from the identify.").
- **All client variants** — removed manual `flush()` calls and the
  mobile background-flush handlers (React Native `AppState` listener,
  iOS `applicationDidEnterBackground`, Android `onStop()`), and
  strengthened the comment wording. FDv2 client SDKs already flush on
  background, and manual flushing is "actively harmful in long-running
  applications" ("Otherwise this kind of stuff builds up and makes it
  harder for us to see problems.").
- **iOS** — moved the top-level `LDClient.start(...)` into a
  `startLaunchDarkly()` setup function ("You can't invoke member
  functions at top level (e.g. LDClient.start()) in swift. You'll need
  to call it from within a setup function.").
- **Android** — fixed the AutoEnv import in `track-only` from
  `com.launchdarkly.sdk.android.LDConfig.Builder.AutoEnvAttributes` to
  `com.launchdarkly.sdk.android.AutoEnvAttributes` ("The import path is
  incorrect for the auto env, but correct in the other example.").

## Open questions

- **Experimentation with AI Config SDKs**: a reviewer asked whether
  experimentation is available when a user is on the AI Config SDKs
  (AgentControl), referencing the experimentation prerequisites support
  matrix, and whether dedicated AI-SDK experimentation snippets are
  warranted. That's a scoping question for the experimentation/AI
  Configs teams, not a change to these frontend snippets; left out of
  this slice.
- **Anonymous context kind**: the same review thread recalled a
  write-up about using a dedicated context kind for the anonymous
  bootstrap context. If that guidance is confirmed, the anonymous
  bootstrap blocks here should adopt that kind.
