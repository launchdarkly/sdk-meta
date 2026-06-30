# Snippets system — agent guide

Auto-loaded by Claude Code when an agent's cwd is anywhere under `snippets/`. Read this before authoring snippets, scaffolds, validators, or CI rows. This document describes the system **as it is** — not as it might become.

For the "how do I write or change a snippet" reference (frontmatter, `validation.*`, templating DSL, render targets, CLI essentials), see [docs/AUTHORING.md](docs/AUTHORING.md). This file covers everything else: repo map, scaffold internals, validator harnesses, CI matrix, conventions, releases.

## Don't author snippet content in unrelated repos

This repo (`launchdarkly/sdk-meta`) is the **canonical home** of every snippet body. Consumers (`gonfalon`, `ld-docs-private`, `ld-docs`) only carry **markers** that point back here:

- `ld-docs` / `ld-docs-private` — `<!-- SDK_SNIPPET:RENDER id="..." -->` markers in MDX. The `snippets render --target=ld-docs` adapter rewrites the fenced code block underneath each marker.
- `gonfalon` (ld-application) — `SDK_SNIPPET:RENDER` markers around JSX components; `snippets render --target=ld-application` rewrites the children.
- `gonfalon` raw-file consumers (e.g. `packages/sdk-info/`) — Vite `?raw` imports of files emitted by `snippets render --target=raw-files` from a YAML manifest.

If you find yourself editing snippet content in a consumer repo, **stop**. The change belongs here as a `.snippet.md` edit; the consumer either picks it up via `snippets render` (CI / sync action) or by re-imported raw files. The only things that legitimately live downstream are the markers themselves and the manifest YAMLs (for raw-files consumers).

## Repo layout

Everything in this guide is relative to `snippets/` (this directory).

```
snippets/
  cmd/snippets/         # CLI entry point (main.go); render | verify | validate | version
  internal/
    model/              # Frontmatter schema + loader (model.go: Frontmatter, Input, Validation, ParseFile, LoadAll)
    render/             # Templating DSL (template.go) + RenderRuntime / RenderForLDApplicationTemplate / RenderForJSXText (render.go)
    adapters/
      ldapplication/    # JSX-marker rewrite for gonfalon
      lddocs/           # MDX-marker rewrite for ld-docs
      rawfiles/         # Manifest-driven file emitter
    validate/           # validate.Run; stages snippets, builds/runs harnesses
    version/            # Version stamped by release-please ldflags
  embedded.go           # Embeds sdks/** into the binary; snippets.SDKsFS() returns the snapshot
  sdks/
    _shared/snippets/<group>/...      # Cross-SDK snippets (no sdk: frontmatter field)
    <sdk-name>/snippets/<group>/...   # Per-SDK snippets
  validators/
    shared/lib.sh                     # require_env, await_success_line, dump_redacted, fail_with_log
    languages/<runtime>/              # Per-language harness
      runner.yaml                     # mode, runs-on, image-prefix
      Dockerfile                      # docker mode only; build context is validators/
      harness/run.sh                  # entrypoint
      scaffold/ or test/              # Optional: pre-built project or driver code
  docs/AUTHORING.md     # Authoring reference: frontmatter, validation.*, DSL, render targets
```

Top-level workflows that touch this tree:

- `.github/workflows/snippets-validate.yml` — fans the matrix across SDKs and runs `snippets validate`.
- `.github/workflows/release-please.yml` — cuts the `snippets/vX.Y.Z` release.

## Snippet file format

Frontmatter schema, `validation.*` fields, the templating DSL, and example shapes for the common kinds live in [docs/AUTHORING.md](docs/AUTHORING.md). Parser invariants worth keeping in mind when extending the system: snippet files are loaded by `internal/model.LoadAll` (see `internal/model/model.go`); decoding uses `KnownFields(true)`, so adding a frontmatter key requires editing the `Frontmatter` struct; `id` is globally unique and collisions error at load; the parsed `Snippet{Path, Frontmatter, CodeLang, CodeBody}` is what every downstream tool consumes.

## Snippet groups (categories)

Every snippet sits in one group folder. The group is the middle segment of the `id` and is what `--group` filters on.

- `sdk-info` — install commands and the canonical init snippet. Consumed by gonfalon's sdk-info panels and ld-docs's quickstart pages. Typically `kind: install` or `kind: init`.
- `sdk-docs` — fragments referenced from ld-docs MDX pages. Mostly `kind: reference` snippets wrapped by syntax-only scaffolds, plus some `kind: flag-eval` runners.
- `observability` — `import` / `initialize` / `install` snippets for the LaunchDarkly observability plugins. Validated through `init-runner-observability` scaffolds that hand-supply the imports normally provided by the matching `observability/import` snippet.
- `ai-configs` — `install` / `context` / `implementation` snippets for AI Configs. Currently node-server-sdk only.
- `experimentation` — experimentation-specific fragments (where present).
- `scaffolds` — `kind: scaffold` snippets only. Not consumed by docs; only by the validator (via wrappees' `validation.scaffold`).
- `_shared/snippets/<group>/` — cross-SDK snippets whose `id` does not carry a sdk prefix. Example: `_shared/snippets/sdk-info/cursor-prompt.snippet.md`.

## Scaffolds

A scaffold is a `kind: scaffold` snippet whose body wraps a wrappee body via a `{{ body }}` slot. The scaffold owns the runtime, entrypoint, requirements, and companions; the wrappee just declares `validation.scaffold: <id>` (and, optionally, `placeholders:` and `scaffold-inputs:`).

### Composition order (validate.stageSnippet)

1. Resolve the wrappee's effective validation snippet via `validation.scaffold:`. Without a scaffold, the snippet renders itself.
2. Render the wrappee body with its own runtime inputs.
3. Apply `validation.placeholders` to that rendered body — literal source-text keys replaced with the value of the named env var (allow-list only).
4. Build the scaffold's input map: scaffold's env-typed inputs + wrappee `scaffold-inputs:` + the special `body` slot set to the substituted wrappee body.
5. Render the scaffold and write to `stageDir/<scaffold.file>`.
6. Stage every companion (from the scaffold) at its own `file:`.

### init-runner + companion pairs

When a runtime needs more than one file at staging time, the scaffold uses `companions:` to point at sibling `kind: scaffold` snippets.

- **Go**: `go-server-sdk/scaffolds/init-runner` has `file: wrappee/init.go` and body `{{ body }}` only; `entrypoint: main.go`; `companions: [go-server-sdk/scaffolds/init-runner-main]`. The companion ships the runner `main.go` that `exec.Command("go", "run", "wrappee/init.go")`s the wrappee, asserts its success line, then emits EXAM-HELLO. Two `package main` files in separate dirs are legal under a single `go.mod` at staging root.
- **iOS**: `ios-client-sdk/scaffolds/init-runner` writes `AppDelegate.swift`; companion `init-runner-viewcontroller` writes `ViewController.swift`. The harness's `cp $SNIPPET_DIR/ViewController.swift` step requires the file. Parallel pair for the syntax-only flavor (`swift-syntax-only` + `swift-syntax-only-viewcontroller`).
- **.NET (server)**: `dotnet-server-sdk/scaffolds/init-runner` ships `Program.cs`; companion `init-runner-csproj` ships `HelloDotNet.csproj` with `Sdk="Microsoft.NET.Sdk.Web"` so `WebApplication.CreateBuilder` resolves.
- **Vue / React Native**: pairs use `init-runner-app` / `init-runner-welcome` companions for the `App.vue` / welcome screens.

### Syntax-only scaffolds and IMPORT_LIFT

ESM and Swift forbid top-level `import` inside a function body; Go requires `import (...)` at file scope. Doc fragments routinely combine "add this import" with "call this method," so several scaffolds embed markers the harness rewrites before invoking the compiler.

`react-client-sdk/scaffolds/react-syntax-only`, `js-client-sdk/scaffolds/js-syntax-only`, and `react-native-client-sdk/scaffolds/react-native-syntax-only` follow this shape:

```
//IMPORT_LIFT_TARGET

async function _wrappee() {
  if (false) {
//BODY_BEGIN
{{ body }}
//BODY_END
  }
}
```

The matching `validators/languages/<runtime>/harness/run.sh` greps for `//IMPORT_LIFT_TARGET` and, when present, runs an awk pass that walks `//BODY_BEGIN` … `//BODY_END`, treats `import` lines (including multi-line `import { ... } from '...';` continuations) as lift candidates, strips `export default` / `export` prefixes, and emits lifted imports just after `//IMPORT_LIFT_TARGET` at module scope. The rest stays inside `_wrappee()`.

Go solves the same problem differently: `go-syntax-only` splits the body into top-level decls (`import (…)`, `func`, `var/const/type/package`) vs function-body residue, splices top-level into a synthesized `package wrappee\n` file, wraps the rest in `func _wrappee()`, then runs `go/parser.ParseFile`. C# uses a `// USING_LIFT_MARKER` comment with a harness pre-stage rewrite to lift `using …;` directives.

### Parse-only stub surfaces

Syntax-only scaffolds typically never execute the body — they just want it to compile/parse. To make unqualified names like `client`, `context`, `config` resolve without a real SDK environment:

- **C++** (`cpp-client/server-syntax-only`): wrappee is a `template <int = 0> void _wrappee()` that is never instantiated; both native and C-binding SDK headers are included at file scope; a polymorphic `struct _AnyClient` declares variadic-template members (`BoolVariation`, `Identify`, `StartAsync`, …) plus an `operator LDClientSDK() const` so the same `client` name resolves whether the body uses the native or C-binding API. `struct _Maxwait` exposes implicit conversions to `unsigned int`, `chrono::milliseconds`, and `chrono::seconds`. `using namespace launchdarkly; using namespace launchdarkly::client_side;` (or `::server_side`) lets unqualified names resolve.
- **.NET** (`dotnet-server-sdk/scaffolds/csharp-syntax-only`): `private static dynamic client = null;` accepts both the v6 `User` API and v7+ `Context` API. Body lives inside `private void Wrappee()` wrapped in `try { … } catch (Exception) { /* never reached */ }`. Requirements pin both `LaunchDarkly.ServerSdk` and `LaunchDarkly.Observability`.
- **iOS** (`swift-syntax-only`): stub `client: LDClient!` on AppDelegate; body lives inside `_wrappee()` (never invoked); AppDelegate prints EXAM-HELLO from `didFinishLaunching`.
- **Python** (`python-syntax-only`): parses via `ast.parse` inside a raw triple-quoted string. No execution. Companion `with-test-data` initializes against `TestData.data_source()` with a pre-populated `your.feature.key=True` flag, exposes `client` and `context` to the body, then evaluates the flag for EXAM-HELLO.

### Placeholders contract

The dispatch order matters for scaffold authors: `validation.placeholders` runs **after** rendering the wrappee and **before** splicing into the scaffold's `{{ body }}` slot. So a placeholder key (e.g. `'YOUR_SDK_KEY'`, `'SDK_KEY'`) only needs to be unique in the wrappee body, not in the composed output — the scaffold's own scaffold-supplied init code can use the same literal text. This is what lets gonfalon keep shipping `'YOUR_SDK_KEY'` placeholders in docs while validation substitutes real keys.

For `env`-keyed harness dispatch (e.g. `INSTALL_KIND` on `ios-install`, `SNIPPET_MODE: flag-eval | syntax-only` on the react harnesses), see [docs/AUTHORING.md#validation-fields](docs/AUTHORING.md#validation-fields) — `validation.env` carries literal KEY=VALUE pairs to the harness with no substitution, and harness scripts switch on the key.

### init-runner-observability scaffolds

Six companion scaffolds (js-client, node-server, python-server, react-client, react-native-client, vue-client) solve the same problem: an `observability/initialize` body declares the LD client + observability plugin but assumes the matching `observability/import` snippet ran at module scope. Each scaffold hand-supplies that module-scope `import` block, splices `{{ body }}` at module scope (or inside an async IIFE for js-client), awaits init via the SDK's appropriate signal (`waitForInitialization`, `is_initialized()` polling, a `Sentinel` rendered inside the body's `LDProvider`, a DOM sentinel for Vue), then emits EXAM-HELLO. Best-effort `flush()` + `close()` after success (Node, Python). Vue / RN also use `companions:` for `init-runner-app` / `init-runner-welcome` screens.

The PR that introduced these scaffolds also extended `js-syntax-only` and `react-native-syntax-only` with IMPORT_LIFT markers so observability `import` snippets — which carry top-level `import …;` directives — can pass the syntax-only path.

### Adding a new scaffold (checklist)

1. **Pick the runtime** — a directory under `validators/languages/<runtime>/` must already exist (with `Dockerfile`, `harness/run.sh`, `runner.yaml`) or you'll need to build it. Reuse where possible.
2. **Single-file or paired** — if the runtime needs >1 file at staging time (manifest, view controller, app shell), make a companion scaffold and list its ID in `validation.companions:`. Set `validation.entrypoint:` on the parent if the entrypoint isn't its own `file:`.
3. **Wire `{{ body }}`** — declare `inputs: { body: { type: string, description: ... } }` and embed `{{ body }}` once in the code block. For runtimes where the body language forbids imports inside functions, add the `//IMPORT_LIFT_TARGET` + `//BODY_BEGIN` / `//BODY_END` triad and either extend an existing awk lift pre-step or add one in `harness/run.sh`.
4. **Stub the runtime surface for parse-only scaffolds** — declare stubs for `client`, `context`, `config` so the body's references resolve at compile/parse time without running.
5. **Emit EXAM-HELLO** — end with the canonical success line (`feature flag evaluates to true` printed to stdout, or written into a known DOM node / `Tests` label). Gate the line on a real init signal (`waitForInitialization`, `is_initialized()`, sentinel render).
6. **Use `placeholders:`, not template markers** — when the body should retain `'YOUR_SDK_KEY'`-style literals for docs rendering but use real keys at validate time, the **wrappee** (not the scaffold) declares `validation.placeholders:` and the dispatcher substitutes pre-splice.
7. **Per-snippet env via `validation.env:`** — set literal KEY=VALUE there when one harness needs to switch behavior per wrappee, and have `harness/run.sh` dispatch on `$KEY`.
8. **Requirements** — Python ships `validation.requirements: |` as the literal `requirements.txt`. Other languages ship a companion manifest scaffold.
9. **Wrappee binding** — the wrappee declares `validation.scaffold: <sdk>/scaffolds/<name>` and, when needed, `validation.scaffold-inputs: { ... }`.

## Validator harnesses

Each language validator lives at `validators/languages/<runtime>/` with this skeleton:

```
runner.yaml      # mode, runs-on, image-prefix
Dockerfile       # docker mode only; build context = validators/
harness/run.sh   # entrypoint
test/ or scaffold/   # optional driver code / pre-built project
```

Shared helpers: `validators/shared/lib.sh`.

### runner.yaml

```yaml
mode: docker            # or: native
runs-on: ubuntu-latest  # or: macos-latest
image-prefix: sdk-snippets/python-validator   # docker mode only
```

- `mode: docker` — Go validator builds the Dockerfile, runs the container with `/snippet` bind-mounted to the staged snippet.
- `mode: native` — Go validator execs `harness/run.sh` on the host and passes `$SNIPPET_DIR` plus the LD env vars. Used by `ios-client` (xcodebuild + Simulator can't run in Linux). No `image-prefix`.
- `runs-on` — hint for the CI matrix; the Go validator itself does not read it.
- `image-prefix` — required for docker mode; local image tag.

### Dockerfile pattern

1. `FROM` a runtime-appropriate base (`python:3.11-slim`, `node:20-bookworm-slim`, `node:22-bookworm`, `mcr.microsoft.com/playwright:v1.59.1-noble`, `eclipse-temurin:17-jdk-noble`).
2. Install OS packages: `apt-get update ... && apt-get install ... && rm -rf /var/lib/apt/lists/*`. `android-client` uses `Acquire::Retries=5` and curl `--retry 3` for flaky noble mirrors.
3. Layer additional toolchains as needed. `shell-install` adds Go via the upstream tarball, pip via `python3-pip`, and `corepack prepare pnpm@9 --activate` + a `yarn@1` pin compatible with Node 20.
4. Pre-bake a project dir with the SDK and dev deps pre-installed so per-validate cycles stay fast. Examples:
   - `js-client` writes `package.json` / `tsdown.config.ts` / `tsconfig.json` / placeholder `src/app.ts` + `index.html` under `/opt/hello-js`, then `npm install` and a pre-warm `npm run build`.
   - `react-native-client` writes `package.json` / `babel.config.js` / `jest.setup.js` / `tsconfig.json` + placeholder `App.tsx` and `src/welcome.tsx` under `/opt/hello-react-native` and runs `npm install`.
   - `android-client` clones `launchdarkly/hello-android` to `/opt/hello-android`, patches `app/build.gradle` (SDK version, Robolectric + junit + androidx test deps, `testOptions { unitTests.includeAndroidResources = true; all { testLogging { showStandardStreams = true } } }`), and copies in `HelloAppTest.kt`.
5. Copy the shared lib and harness, then set the entrypoint:

```dockerfile
COPY shared /harness-shared
COPY languages/<runtime>/harness /harness
ENTRYPOINT ["/harness/run.sh"]
```

Build context is `validators/` (see the comment in `python/Dockerfile`), which is why paths are `shared/` and `languages/<runtime>/...`. Pin SDK + toolchain versions as `ARG`s (`LD_JS_CLIENT_SDK_VERSION`, `LD_RN_SDK_VERSION`, `LD_ANDROID_SDK_VERSION`, `CMDLINE_TOOLS_VERSION`, `ROBOLECTRIC_VERSION`, ...).

### harness/run.sh contract

Every harness starts with `set -eu`, sources `lib.sh`, then `require_env`s the inputs it needs. Inputs passed by the Go validator:

- `SNIPPET_ENTRYPOINT` — path under `/snippet` (docker) or `$SNIPPET_DIR` (native) to the snippet's entry file.
- `SNIPPET_DIR` — native-mode only.
- `SNIPPET_MODE` — optional dispatch hint (e.g. `flag-eval`, `syntax-only`).
- `INSTALL_KIND` — `ios-install`-specific discriminator (swift-package / podfile / cartfile).
- `LAUNCHDARKLY_SDK_KEY` (server-side), `LAUNCHDARKLY_MOBILE_KEY` (mobile), `LAUNCHDARKLY_CLIENT_SIDE_ID` (browser), `LAUNCHDARKLY_FLAG_KEY` (everywhere except shell-install).

**Success contract**: EXAM-HELLO hello-world snippets emit a line matching `feature flag evaluates to [Tt]rue` (Python prints `True`, others lowercase; quoting around the flag key varies). The harness either:

- Backgrounds the snippet and calls `await_success_line "$LOG" "$PID" "$deadline"` (succeeds on first regex hit, SIGTERMs the process, prints the matched line), or
- Waits for the runner to exit and `grep`s the captured log — used when the runner's output (jest, xcodebuild) already contains the regex verbatim and a false-positive on a failure printout would otherwise match.

On failure: `fail_with_log "$LOG" "<message>"` prints the message + the redacted log and exits 1.

**shell-install is the outlier**: no flag evaluation. Post-install assertions instead — `node_modules/<pkg>/`, `pip show <pkg>`, `grep <modpath> go.mod`, `bower_components/`.

### shared/lib.sh helpers

Source via `. /harness-shared/lib.sh` (docker) or relative path (native).

- `require_env <NAME>...` — POSIX-shell loop; prints `<name> not set` to stderr and `exit 1` if any var is empty.
- `await_success_line <log-file> <pid> <deadline-epoch>` — polls every 0.2s for the EXAM-HELLO regex. On match: SIGTERMs the pid, prints the matched line and `validator: ok`, returns 0. On pid-exit or deadline: returns 1.
- `dump_redacted <log-file>` — substitutes `LAUNCHDARKLY_SDK_KEY` / `LAUNCHDARKLY_MOBILE_KEY` / `LAUNCHDARKLY_CLIENT_SIDE_ID` with `<REDACTED_<var>>` via a built sed program. `_sed_escape_pattern` backslash-escapes `][\.*^$|/&` so credentials containing regex metas or the `|` delimiter can't break out.
- `fail_with_log <log-file> <message>` — prints `validator: <msg>` and `--- snippet output (keys redacted) ---` to stderr, calls `dump_redacted`, exits 1.

### Per-harness shapes worth knowing

- **shell-install** — single multi-toolchain image. `cat $BODY` → awk-extract the first token (`LEAD`) and second token (`SUB`) → `case` over `npm|pnpm|yarn|pip|pip3|go|bower`. Each branch sets up pre-state (`npm init -y`, `python3 -m venv .venv`, `go mod init`), runs the body, asserts post-state. Unknown leading tokens are a hard error so a new install style must teach the harness.
- **js-client** — `npm run build` (tsdown bundles `src/app.ts` → `dist/app.js` with `platform: 'browser'`, `noExternal: ['@launchdarkly/js-client-sdk']` so the LD SDK is inlined). Then `node /harness/check.js` launches headless Chromium via Playwright, navigates to `file:///opt/hello-js/index.html`, mirrors console + pageerror to stdout, polls `page.textContent('body')` for the EXAM-HELLO regex.
- **react-native-client** — `timeout --signal=TERM 180s npm test` (jest with `@react-native/jest-preset`). Pre-baked `App.test.tsx` mocks `@launchdarkly/react-native-client-sdk` to force `initialConnectionMode: 'polling'` (RN streaming needs `XMLHttpRequest` which doesn't exist in Node), renders `<App/>`, `waitFor`s a flattened body-text regex, logs the flat text. `jest.setup.js` pins `AppState.currentState = 'active'`. Greps `$LOG` after jest exits (jest's failure printout contains the regex verbatim, which would falsely match `await_success_line`).
- **python** — `pip install --quiet -r requirements.txt` if present; then `PYTHONUNBUFFERED=1 timeout 30s python "$SNIPPET_ENTRYPOINT"` in background + `await_success_line` with a 25s deadline.
- **ios-client** (native) — sources `lib.sh` via relative path. Copies `validators/languages/ios-client/scaffold/` into a `mktemp -d`, drops in `AppDelegate.swift` and `ViewController.swift`. A Python in-script pass lifts any `import` lines that ended up inside a function body back to file scope. `xcodegen generate` against `project.yml`, picks an iPhone simulator dynamically via `xcrun simctl list devices available --json`, then `xcodebuild test -destination "platform=iOS Simulator,name=$SIM_NAME" CODE_SIGNING_ALLOWED=NO`. Passes LD env into the simulator via `SIMCTL_CHILD_LAUNCHDARKLY_*` prefixes.
- **android-client** — copies snippet `.kt` files into the pre-baked scaffold's package dir. A Python in-script pass lifts function-body `import` lines back to file scope and substitutes `BaseApplication` → `MainApplication` (snippet's literal Application class → the scaffold's class expected by `@Config(application = MainApplication::class)`). `timeout 600s ./gradlew --no-daemon testDebugUnitTest --tests='*HelloAppTest*' --console=plain` in background + `await_success_line` with a 590s deadline. Robolectric drives the activity lifecycle on the JVM (no emulator).

### Adding a new language validator

1. Create `validators/languages/<runtime>/runner.yaml` with `mode`, `runs-on`, and (docker only) `image-prefix`.
2. For docker mode, write `Dockerfile` using `validators/` as the build context. Pin SDK + toolchain versions as `ARG`s. Pre-bake a placeholder project matching the snippet's file layout so per-validate stays fast. End with the `COPY shared` / `COPY languages/<runtime>/harness` / `ENTRYPOINT` triple.
3. For native mode, omit `image-prefix`; the harness sources `lib.sh` via `$(cd "$(dirname "$0")/../../../shared" && pwd)/lib.sh` and reads snippet files from `$SNIPPET_DIR`.
4. Write `harness/run.sh`. Start with `set -eu; . /harness-shared/lib.sh; require_env <LD-key-var> LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT`. Pick the LD key var matching the SDK audience: server → `LAUNCHDARKLY_SDK_KEY`, mobile → `LAUNCHDARKLY_MOBILE_KEY`, browser → `LAUNCHDARKLY_CLIENT_SIDE_ID`.
5. Stage snippet files from `/snippet/...` into the pre-baked project, run under `timeout` with stdout/stderr to `$LOG`, then `await_success_line "$LOG" "$PID" "$deadline"` (background, regex hit) or grep after exit when the runner echoes the regex itself.
6. Harness-internal driver code (jest test, Robolectric test, Xcode scaffold) goes under `<runtime>/test/` or `<runtime>/scaffold/` and is `COPY`ed in via the Dockerfile.
7. Install-style snippets that don't evaluate a flag follow `shell-install`'s shape: sniff the body's leading token, dispatch over package managers, assert install side-effects. Reserve `INSTALL_KIND` for cases (like `ios-install`) where the body can't be sniffed.

## CLI reference

Entry point: `cmd/snippets/main.go`. Subcommands:

```
snippets render   --target=<adapter> [...]
snippets verify   --target=<adapter> [...]
snippets validate --sdk=<id> [...]
snippets version
```

No subcommand (or `-h` / `--help` / `help`) prints the usage block; unknown subcommand exits 2.

### `render`

Rewrites consumer files in place.

| Flag | Notes |
|---|---|
| `--target` | Required. One of `ld-application`, `ld-docs`, `raw-files`. |
| `--entrypoint` | Repeatable. Required for `ld-application` and `ld-docs`. Directory inside the consumer checkout to walk for marker files. |
| `--manifest` | Required for and exclusive to `--target=raw-files`. Path to the YAML manifest of `(snippet-id, output-path)` pairs. |
| `--consumer` | Consumer-checkout root the manifest's `out:` resolves against. Defaults to the manifest's directory. |
| `--sdks` | Path to a working `sdks/` tree. Default: the embedded snapshot (see below). |

Adapter behavior:

- `ld-application` (`internal/adapters/ldapplication.Render`) — walks each `--entrypoint`, finds files with `SDK_SNIPPET:RENDER` markers, rewrites JSX children inside marked components. Used by gonfalon.
- `ld-docs` (`internal/adapters/lddocs.Render`) — same marker flow, rewrites fenced code-block bodies in MDX. Used by ld-docs / ld-docs-private.
- `raw-files` (`internal/adapters/rawfiles.Render`) — reads the manifest and writes each rendered body to `<consumer>/<manifest.out>/<entry.path>`. Used by gonfalon `packages/sdk-info/` and similar `?raw`-import consumers.

Output: `rewrote <path>` per changed file (marker targets) or `wrote <path>` per emitted file (raw-files); `no changes` / `no files written` when idle.

Realistic invocations:

```
snippets render --target=ld-application --entrypoint=./packages/sdk-info --sdks=../sdk-meta/sdks
snippets render --target=ld-docs --entrypoint=./src/content/sdk --entrypoint=./src/content/guides
snippets render --target=raw-files --manifest=./packages/sdk-info/snippets.yaml
```

### `verify`

Read-only dry run of `render` for marker-driven targets (`raw-files` is intentionally not supported). Same flags as `render` minus `--manifest` / `--consumer`.

Fails if rendered bytes differ from disk or a marker's hash does not match its region's current contents. Prints `ok` on success, exits 1 on drift. Never writes; never executes snippet code. Used as a CI guard against drift.

### `validate`

Stages snippet bodies, builds the per-language validator (Docker or native), runs each snippet against a real LaunchDarkly environment.

| Flag | Notes |
|---|---|
| `--sdk` | Required. Matches `sdks/<sdk>/`. Filters snippets to those whose frontmatter `sdk:` matches. |
| `--snippet` | Optional. Restricts to a single snippet id. |
| `--snippet-skip` | Optional. Skips one snippet id. CI uses this to split a slow snippet onto its own matrix row. |
| `--group` | Optional. Filters on the middle segment of the snippet id (`sdk-info`, `sdk-docs`, ...). |
| `--sdks` | Path to a working `sdks/` tree. Default: embedded snapshot. |
| `--validators` | Path to the `validators/` directory. Default: `./validators`. Must be on disk (Docker `COPY` reads from it). |
| `--jobs` | Max concurrent batch-harness invocations (batch-mode validators only). Default `0` = `NumCPU`. |
| `--image-cache` | Cross-run Docker layer cache for `mode: docker` validators: `gha` (GitHub Actions cache) or a registry ref prefix (`type=registry`). Empty (default) = plain `docker build`, no buildx. CI sets it only on non-fork builds. |

Behavior:

1. Reads `LAUNCHDARKLY_SDK_KEY`, `LAUNCHDARKLY_FLAG_KEY`, `LAUNCHDARKLY_MOBILE_KEY`, `LAUNCHDARKLY_CLIENT_SIDE_ID` from env. Snippets with typed inputs / placeholders requiring a key fail fast if it's empty.
2. Loads all snippets, applies filters, discards scaffolds and snippets without `validation.runtime` or `validation.scaffold`.
3. Resolves the effective validation snippet (scaffold-bound snippets defer to their scaffold for runtime / entrypoint / companions / requirements).
4. Stages the rendered body (+ companions, `requirements.txt`, scaffolded `body` slot) into a temp dir, runs the language harness:
   - `mode: docker` — builds `validators/.../Dockerfile` with build context = entire `validators/` tree, tags with a content hash of `shared/` + the runner dir, then `docker run --rm -v <stage>:/snippet:ro` with `SNIPPET_ENTRYPOINT` and LD env vars forwarded.
   - `mode: native` — execs `<runner>/harness/run.sh` with `SNIPPET_DIR` and `SNIPPET_ENTRYPOINT` exported.
5. Errors out with `no validatable snippets found ...` if nothing matched.

Example:

```
snippets validate --sdk=python-server-sdk --group=sdk-info --sdks=./sdks --validators=./validators
snippets validate --sdk=flutter-client-sdk --snippet=flutter-client-sdk/sdk-info/install
```

### Embedded `sdks` FS vs `--sdks`

`resolveSDKsFS` in `main.go`:

```go
func resolveSDKsFS(sdksFlag string) fs.FS {
    if sdksFlag == "" {
        return snippets.SDKsFS()
    }
    return os.DirFS(sdksFlag)
}
```

- Released binaries ship with an embedded snapshot of `sdks/` (`snippets.SDKsFS()`) so downstream consumers can `go run` / drop in the binary without checking out `sdk-meta`.
- Local authoring passes `--sdks=./sdks` so edits to the working tree are picked up immediately without rebuilding.
- `--sdks` is accepted by `render`, `verify`, and `validate`.
- `validators/` is always read from disk — Docker's build context needs a filesystem path; no embedded equivalent.

## CI: `.github/workflows/snippets-validate.yml`

Triggers on PRs touching `snippets/**` or the workflow itself, plus `workflow_dispatch`. Concurrency group cancels prior runs on the same ref. Permissions: `id-token: write` (OIDC) and `contents: read`.

One job `validate` with `fail-fast: false` and an `include:`-only matrix.

### Matrix row shape

| Field | Required | Notes |
|---|---|---|
| `sdk` | yes | Matches `snippets/sdks/<sdk>/`. |
| `runs-on` | yes | `ubuntu-latest` everywhere except the three `ios-client-sdk` rows (`macos-latest`). |
| `key-type` | yes | `server` | `client` | `mobile` | `none`. Determines which LD key the verify-hello-app wrapper injects. |
| `snippet` | – | Restrict to a single bound snippet path. |
| `snippet_skip` | – | Exclude one bound snippet. Note: matrix key uses underscore; CLI flag uses `--snippet-skip` (dash). |
| `group` | – | Restrict to one group (`sdk-info`, `sdk-docs`, ...). |
| `label` | – | Overrides the job name. **Required** when multiple rows share an `sdk` so artifacts don't collide. |

Job name renders as `${{ matrix.label || matrix.sdk }}`. Rows are grouped in the file by purpose: server-key SDKs, JS client, mobile-key dotnet, shell-install-only (dotnet-server-sdk with an unused `key-type: server`), client/mobile init-binding SDKs, the three iOS rows, the android row, and the sdk-docs-syntax-only block at the bottom.

### `verify-hello-app` wrapper

`launchdarkly/gh-actions/actions/verify-hello-app@verify-hello-app-v2.0.1` — the same action the `hello-*` example repos use:

1. Assumes the AWS role in `vars.AWS_ROLE_ARN_EXAMPLES` via GitHub OIDC.
2. Pulls LaunchDarkly Sandbox credentials from AWS Secrets Manager and injects one of `LAUNCHDARKLY_SDK_KEY` / `LAUNCHDARKLY_CLIENT_SIDE_ID` / `LAUNCHDARKLY_MOBILE_KEY` based on the `use_*_key` flag.
3. Injects `LAUNCHDARKLY_FLAG_KEY` from SSM `/sdk/common/hello-apps/boolean-flag-key`.
4. Runs `command:` and asserts the output contains the EXAM-HELLO line.

Key-type mapping:

```yaml
with:
  use_server_key: ${{ matrix.key-type == 'server' }}
  use_client_key: ${{ matrix.key-type == 'client' }}
  use_mobile_key: ${{ matrix.key-type == 'mobile' }}
  role_arn: ${{ vars.AWS_ROLE_ARN_EXAMPLES }}
```

Exactly one `use_*_key` is true per row (when `key-type != 'none'`).

### `key-type: none` bypass

Skips `verify-hello-app` entirely and runs a plain shell step `Validate (no key)`. Used for rows whose only bound snippets are install fragments — they exercise package managers (`pod install`, `carthage bootstrap`, `swift package resolve`) in an empty workdir and never produce the EXAM-HELLO marker the action greps for. The two validate steps are mutually exclusive (`if: matrix.key-type != 'none'` vs `if: matrix.key-type == 'none'`); `Record result` treats success on either as ok.

### Optional matrix fields → CLI flags

The optional fields flow through env vars into the `go run ./cmd/snippets validate` invocation in both validate steps:

```yaml
env:
  MATRIX_SDK: ${{ matrix.sdk }}
  MATRIX_SNIPPET: ${{ matrix.snippet }}
  MATRIX_SNIPPET_SKIP: ${{ matrix.snippet_skip }}
  MATRIX_GROUP: ${{ matrix.group }}
```

The shell appends `--snippet=`, `--snippet-skip=`, `--group=` only when the corresponding env var is non-empty. `--sdks=./sdks --validators=./validators` are always passed.

`label` is consumed only by the job name and the artifact-suffix step (`Compute artifact suffix` slugifies `LABEL` via `tr ' ()/' '----' | tr -s '-' | sed 's/-$//'`, falling back to `sdk` when empty). Two rows sharing one `sdk` MUST set distinct `label` values or their artifacts collide.

### Artifact upload + summary

Per-cell `Record result` writes `/tmp/result/status` (`ok` or `fail`) and copies `/tmp/validate.log` to `/tmp/result/log`. `Upload result` uses `actions/upload-artifact@v4` with name `result-<slug>`, `if-no-files-found: ignore`, 14-day retention.

The `summary` job (`needs: validate`, `if: always()`) downloads `result-*` artifacts via `actions/download-artifact@v4` with `pattern: result-*` into `results/`, renders a markdown table in `$GITHUB_STEP_SUMMARY` with columns SDK / Result / Excerpt. Failing cells get the last 3 non-empty log lines (pipes and backticks escaped). A final step re-walks artifacts and `exit 1`s if any status != `ok` — a green summary job means every cell passed.

### CI recipes

- **Add a new SDK row** (server-key, plain): append `- { sdk: <name>, runs-on: ubuntu-latest, key-type: server }` to the appropriate block. The SDK directory must already exist and the CLI must already pick it up.
- **Add a credentials-less row**: use `key-type: none`; the workflow auto-routes to `Validate (no key)`.
- **Split one SDK across multiple rows**: repeat `sdk:` with distinct `label:` values, partition with `snippet:` / `snippet_skip:` / `group:`. `label` is required for artifact uniqueness.
- **Change a row's key-type**: edit just `key-type`. The `use_*_key` booleans recompute automatically; nothing else to touch unless you're crossing the `none` boundary.
- **Thread a new optional env var**: (1) add `MATRIX_FOO: ${{ matrix.foo }}` to **both** the `Validate` and `Validate (no key)` `env:` blocks; (2) add `[ -n "$MATRIX_FOO" ] && args="$args --foo=$MATRIX_FOO"` to both `command`/`run` shells; (3) set `foo:` on the rows that need it. Keep the two steps in sync — divergence between them is a common bug.

## Local testing

Run all `go` and `snippets` commands from the `snippets/` subdir (this directory). That's where `go.mod`, `cmd/snippets/`, `internal/`, and `validators/` live, and where `--validators=./validators` resolves.

### LD keys for local validate runs

The `/tmp/ld-keys.env` file is the convention; source it before invoking the CLI:

```bash
set -a; source /tmp/ld-keys.env; set +a
go run ./cmd/snippets validate --sdk=python-server-sdk --sdks=./sdks
```

The file must export `LAUNCHDARKLY_SDK_KEY`, `LAUNCHDARKLY_FLAG_KEY`, and (where the SDK requires them) `LAUNCHDARKLY_MOBILE_KEY` / `LAUNCHDARKLY_CLIENT_SIDE_ID`. Missing keys produce a clear `snippet X input Y (type=...) requires LAUNCHDARKLY_... to be set` error before any Docker build kicks off.

### Go-tooling commands

```bash
# Build the binary
go build -o snippets ./cmd/snippets

# Unit tests (model parser, render DSL, adapters; NOT language harnesses)
go test ./...
```

`go test ./...` covers `internal/model` (parser), `internal/render` (DSL), and adapter render/verify behavior. Language harnesses run only under `snippets validate`. For the day-to-day `validate` / `render` / `verify` invocations, see [docs/AUTHORING.md#cli-essentials](docs/AUTHORING.md#cli-essentials).

### Docker prerequisites

`snippets validate` requires Docker available on `PATH` for any non-native runtime. The Go validator builds the image with build context `./validators` and tags it from `image-prefix`, so first-run latency is dominated by the Dockerfile's pre-bake steps. Subsequent runs are fast (image cache + content-hash tag).

## Conventions

### Commit messages

Conventional Commits with `snippets` scope: `feat(snippets): …`, `fix(snippets): …`, `chore(snippets): …`, `docs(snippets): …`. The release-please config drives `snippets/CHANGELOG.md` and the `snippets/vX.Y.Z` release from these.

### Port-notes files

When porting snippets from gonfalon, drop a `<sdk>/PORT-NOTES.md` (or per-snippet sibling notes) so the next agent has a paper trail of what changed and why. Don't litter the snippet bodies themselves with porting commentary — describe behavior in code comments, history in the notes file.

### When the existing validator can't validate a snippet

Every snippet that can be validated MUST be wired up. Routes when the obvious scaffold doesn't fit:

1. **Extend the scaffold's stub surface** so the body's references resolve. Most `*-syntax-only` scaffolds add file-scope stubs for `client`, `context`, etc. — extend them when a body needs an SDK type the stub doesn't expose. Adding a stub is preferred over deferring validation.
2. **Add a per-snippet scaffold** when one snippet's shape is irreconcilable with the rest of the SDK's bodies (e.g. an Application-host body alongside plain expression-fragment bodies). The wrappee's `validation.scaffold:` picks the right one.
3. **Add a new validator** when the runtime has no harness yet (HTML, XML, YAML, BrightScript). Lightweight Docker images are fine — `xmllint`, `prettier --check`, `brighterscript` lex+parse, etc. See "Adding a new language validator" below.
4. **Fix the snippet** when it's mistagged (e.g. `lang: php` for a shell command body), references API the current SDK no longer exposes, or has actual syntax errors. Inline the fix in the same PR that wires up validation; if the change is policy/product-sensitive, open a separate snippet-bugs PR that lands first.

Don't park snippets as "no validation possible" without an explicit, documented blocker — and even then, the blocker should describe what concretely needs to happen to wire it up.

### Code-style guardrails (apply broadly)

- No spec numbers in code (`// spec section 3.2.1`) — paraphrase. Spec numbers go stale.
- No PR / review / issue refs in code (`// fixes after review #4`) — describe the behavior. The history is in git.
- No unicode arrows in code or docs — use `-->`.
- Avoid the word "polling" in LD streaming code — it's a product term for the polling data source. Use "continue" / "keep reading" / "next iteration".
- Dart code: use case-pattern null narrowing (`if (x case final v?)`) instead of `!`; run `dart format` before committing (CI enforces it separately from analyze/test).
- Rust code: run `cargo fmt` before committing (CI enforces `rustfmt` separately from build/test).

## Releases (`.github/workflows/release-please.yml`)

The repo runs release-please in monorepo mode with two packages, `api-js` and `snippets`. The `snippets` package is configured `"draft": true`, so the first release-please call creates a draft GitHub release without a git tag. The `release-snippets` job then:

1. Pushes the `snippets/vX.Y.Z` tag manually (drafts don't auto-tag).
2. Cross-compiles via plain `go build` for linux/amd64, linux/arm64, darwin/amd64, darwin/arm64. (goreleaser was abandoned because its monorepo support is Pro-only.) Each archive is named `snippets_<version-without-v>_<goos>_<goarch>.tar.gz` with `snippets`, `README.md`, `LICENSE` at the archive root (no wrapper dir).
3. Embeds the bare version via `-X .../internal/version.Version=<version>` ldflags.
4. Generates `checksums.txt` (sha256sum of all archives).
5. Attests each archive with `actions/attest-build-provenance@v2` (SLSA, signed via OIDC, stored on GitHub).
6. Uploads archives + checksums to the draft release with `gh release upload --clobber`.
7. Flips the release out of draft with `gh release edit --draft=false`. After this step the release is immutable under repo policy, so every asset must be uploaded **first**.

Downstream consumers (e.g. the `sync-snippets` action) resolve the latest release by the `snippets/v*` tag prefix, strip `snippets/v` to get the bare version, fetch `snippets_<version>_<goos>_<goarch>.tar.gz`, and verify with `gh attestation verify <archive> --repo launchdarkly/sdk-meta`. Producer and consumer agree on the bare-version filename — that's why the build strips the leading `v`.

## Source-of-truth quick links

When in doubt, the code is authoritative; this guide paraphrases it. Frequent lookup targets:

- Frontmatter schema, loader, parser: `internal/model/model.go` (`Frontmatter`, `Input`, `Validation`, `ParseFile`, `LoadAll`).
- Template DSL parser + nodes: `internal/render/template.go` (`Parse`, `Node`, `Literal`, `Var`, `Cond`, `tokenRe`).
- Render modes (runtime / ld-application / JSX text), filters, foreign-template passthrough: `internal/render/render.go`.
- Validation runner (filter, stage, harness exec): `internal/validate/validate.go`.
- CLI entry, usage block, flag wiring: `cmd/snippets/main.go`.
- Shared harness helpers: `validators/shared/lib.sh`.
- CI matrix and key-type wiring: `.github/workflows/snippets-validate.yml`.
- Release tag + artifact contract: `.github/workflows/release-please.yml`.
- Snippet author reference (frontmatter, `validation.*`, DSL, render targets, CLI essentials): `docs/AUTHORING.md`.
