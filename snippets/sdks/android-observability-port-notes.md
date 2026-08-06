# Android observability plugin port notes

Source: `ld-docs-private/fern/topics/sdk/observability/android.mdx` — the
LaunchDarkly Android observability plugin (EAP) reference page, 38 code
blocks. Snippets live in `android-client-sdk/snippets/observability/`.

## Coverage expansion (current-doc reconciliation)

The doc grew past the original 29-block port: it added a "Configure product
analytics event collection" section (`ObservabilityOptions.Analytics`) and a
"Manual instrumentation" section (record logs, record traces, `track`,
`trackScreenView`), and rewrote the distributed-tracing blocks to the
plugin's map-based API (`startSpan(name, properties = mapOf(...))`,
`recordLog(..., properties = ...)`, 1-arg `startSpan(name)`). All of these
are valid against 0.60.0 (the map/`properties` overloads are interface
default methods on `Observe`; `Analytics`, `track`, and `trackScreenView`
all ship in 0.60.0). Snippets were updated/added to match:

- Tracing snippets now use the map-based style. Java keeps the compilable
  `LDObserve.Companion.startSpan(name, new HashMap<>())` form (no `@JvmStatic`
  on the companion), where the doc's hand-written Java dropped `.Companion`
  and would not compile.
- New snippets: `import-java`/`import-kotlin`,
  `configure-product-analytics-kotlin`, `record-logs[-typed]-kotlin`,
  `record-traces[-typed]-kotlin`, `track-event-kotlin`,
  `track-screen-view-kotlin`.
- `ObservabilityOptions.Analytics` has no `pageViews` field on 0.60.0 or
  plugin `main`; the doc's Analytics block and its option list reference one,
  so the canonical snippet omits it (fields: `taps`, `trackEvents`,
  `screenViews`, `appLifecycle`, `appLaunch`).
- Scaffolds gained a `performDatabaseQuery()` stub (the record-traces
  fragments call it) and `java.util.HashMap`/`Map` imports (the Java tracing
  fragments construct `new HashMap<>()`).

## Validation

Real type-check against the published plugin
`com.launchdarkly:launchdarkly-observability-android:0.60.0` (standardized on
the current version; the docs' install block pinned a stale 0.21.0):

- **Kotlin + Java blocks** — the `android-client` validator (Kotlin 1.8,
  compileDebug{Kotlin,JavaWithJavac}) with the plugin + opentelemetry-api
  1.51.0 + okhttp added. New `kotlin-observability`, `kotlin-observability-coroutine`,
  and `java-observability` scaffolds. OTel types are imported **explicitly**,
  not by wildcard: the plugin aar ships a
  `com.launchdarkly.observability.replay.Attributes` that otherwise shadows
  `io.opentelemetry.api.common.Attributes` for the unqualified name the
  fragments use.
- **Compose / View masking blocks** — a dedicated `android-compose` validator
  (Kotlin 2.0.21 + Compose 1.7.x). The plugin's Compose masking API
  (`Modifier.ldMask()`) is built against Compose 1.7.x, which carries
  Kotlin-2.0 metadata the Kotlin-1.8 android-client validator can't read, so
  it gets its own project. The View masking block's `R.layout.activity_login`
  / `R.id.password` resolve against a layout baked into that validator.

## Snippet fixes (reconciled to the 0.60.0 API / valid code)

- `Instrumentations(activityLifecycle = ...)` -> the 0.60.0 params
  `userTaps`/`screens` (activityLifecycle was removed).
- `ReplayOptions(capturePeriodMillis = ...)` -> removed (no such param in 0.60.0).
- The "log with context" tracing fragment was truncated in the docs
  (`span.en`) -> `span.end()`.
- Java `Collections.singletonList<Plugin>(Observability(...))` (Kotlin syntax)
  -> `Collections.<Plugin>singletonList(new Observability(...))`.

## Plugin Java-interop gaps found (worth an SDK follow-up)

The Java examples only compile as non-idiomatic forms because the plugin
lacks Java-friendly annotations:

- `LDObserve` is a companion object with no `@JvmStatic`, so Java must call
  `LDObserve.Companion.startSpan(...)` (matches the SDK's own pure-Java e2e in
  observability-sdk#605).
- The `Observability` plugin constructor has no `@JvmOverloads` (and a data
  class `ObservabilityOptions` has no Java no-arg constructor), so Java init
  must pass the full 4-arg form
  `new Observability(app, key, ObservabilityOptions.builder().build(), null)`.
  Adding `@JvmOverloads` would let the docs use the clean 2-arg form.

## Not validated (rendered from canonical only)

- The two Gradle dependency install blocks (Android AAR coordinates can't be
  resolved by the shell-install plain-JVM gradle project; the android-client
  validator compiles against these exact artifacts).
- The two byte-buddy instrumentation Gradle config blocks and the W3C
  trace-context HTTP header example (build config / headers, not code).
