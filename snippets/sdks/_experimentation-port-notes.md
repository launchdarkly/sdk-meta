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

## Bucket C: no validation block

**Severity**: low

**Snippets affected**: every `*/experimentation/{track-only,full}` entry.

**Why unbindable**: these are experimentation onboarding examples, not
standalone-runnable programs. Each one defines an app component plus
helper functions (`trackMetric`, `identifyUser` / `onUserEligible` /
`onUserBecomesEligible`) that call into the host app's own surfaces
(`YourComponent`, `applyVariant`, an `Activity`/`AppDelegate`
lifecycle). The existing per-SDK validator scaffolds
(`*-syntax-only`, `init-runner`) wrap a single fragment; they can't host
a multi-function example that references caller-supplied components.
They are also newly *proposed* snippets — there is no consumer
(gonfalon) wiring rendering them yet, so there is no marker/hash to
verify against either. Left as Bucket C: the canonical text lives in
sdk-meta, byte-checked by review, with no `validation:` block.

`snippets validate` only runs snippets that declare a `validation:`
block, and `.github/workflows/snippets-validate.yml` targets specific
snippets/groups per SDK (never `experimentation`), so these add no CI
surface.

**Recommended action**: when a consumer adopts the experimentation
onboarding flow, add an `experimentation-syntax-only` scaffold per
platform that stubs the caller-supplied components
(`YourComponent`/`applyVariant`) so the bodies can at least be
parse/type-checked, mirroring how `react-native-syntax-only` and
`swift-syntax-only` bind the sdk-docs fragments.

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
