# Authoring snippets

How to write or change a `.snippet.md` file. For repo navigation, scaffold internals, validator harness internals, and CI wiring, see [../CLAUDE.md](../CLAUDE.md).

A snippet lives at `sdks/<sdk>/snippets/<group>/<name>.snippet.md` (or `sdks/_shared/snippets/<group>/<name>.snippet.md` for cross-SDK snippets). It is a single Markdown file with YAML frontmatter and one fenced code block:

```markdown
---
id: python-server-sdk/sdk-info/init
sdk: python-server-sdk
kind: init
lang: python
file: python-server-sdk/init.txt
description: Client initialization snippet for python-server-sdk.
validation:
  runtime: python
  entrypoint: init.py
  requirements: |
    launchdarkly-server-sdk
---

```python
import ldclient
# ...
```
```

The parser:

1. Reads YAML frontmatter fenced by `---` lines at the top of the file.
2. Captures the **first** fenced code block; anything between frontmatter and fence is prose and is ignored.
3. Decodes frontmatter with `KnownFields(true)` — unknown keys are a parse error. Typos (`Description:`, `entrypiont:`) fail loudly.
4. Enforces globally unique `id` across the whole tree.

## Snippet groups

Every snippet sits in one group folder; the group is the middle segment of `id` and is what `--group` filters on. See [../CLAUDE.md#snippet-groups-categories](../CLAUDE.md#snippet-groups-categories) for the full list and what consumes each one.

## Top-level frontmatter fields

| Field | Type | Required | Notes |
|---|---|---|---|
| `id` | string | yes | Globally unique. Convention: `<sdk>/<group>/<name>`. `_shared` snippets drop the sdk prefix (e.g. `id: cursor-prompt`). |
| `sdk` | string | – | Matches a directory under `sdks/`. Omit for `_shared/` snippets. |
| `kind` | string | yes | Categorical label. Live values: `install`, `init`, `initialize`, `import`, `context`, `implementation`, `flag-eval`, `reference`, `scaffold`, `manifest`. |
| `lang` | string | yes | Fenced-code language tag. Also the default `validation.runtime` when that field is empty. |
| `file` | string | – | Relative path consumers see and the validator writes the rendered body to. Omitted on `kind: reference` (the wrapping scaffold owns the path). |
| `description` | string | recommended | One-liner. Supports YAML `|` block scalars for multi-paragraph rationale (scaffolds). |
| `inputs` | map | – | Declared templating inputs. Each entry: `type` (e.g. `string`, `flag-key`), `description`, optional `runtime-default`. Use `inputs: {}` for scaffolds that take no parameters. |
| `validation` | map | – | Absence means "don't run through the validator." See below. |

Adding a new top-level key requires editing `internal/model/model.go`.

## `validation.*` fields

| Field | Notes |
|---|---|
| `checks` | Optional list of `Check` entries (multi-check shape, see below). When set it's authoritative — top-level fields below act as defaults each entry inherits. When unset and any of the top-level fields below is set, the loader synthesizes a single `[{kind: runtime, …}]` check from them — every existing snippet's behavior is unchanged. |
| `runtime` | Picks a harness under `validators/languages/<runtime>/` (e.g. `python`, `node`, `shell-install`, `ios-install`, `js-client`, `react-native-client`, `android-client`). Falls back to `lang:` if empty. Set explicitly when `lang` and runtime differ (e.g. `lang: javascript` + `runtime: node`; `lang: ts` + `runtime: js-client`). |
| `entrypoint` | File the harness invokes inside the staging dir. Defaults to `file:`. |
| `requirements` | Runtime-specific dependency descriptor. Python → `requirements.txt` contents (use `|`). Node → top-level dep name. .NET → NuGet package id. Other languages ship deps via a companion manifest scaffold. |
| `companions` | List of snippet IDs staged alongside; each is rendered with the same inputs and written to its own `file:`. Used for multi-file projects. |
| `scaffold` | ID of a `kind: scaffold` snippet to wrap this body in. **Mutually exclusive** with `runtime` / `entrypoint` / `requirements` / `companions` — those come from the scaffold. |
| `scaffold-inputs` | Map of named values passed into the scaffold's template beyond the implicit `body`. Lets one scaffold serve wrappees with slightly different setup. |
| `env` | Literal `KEY: VALUE` pairs forwarded into the harness process. No substitution. Used as a harness-side discriminator (e.g. `INSTALL_KIND: podfile`, `SNIPPET_MODE: flag-eval`). |
| `placeholders` | Literal source-text fragments **inside the rendered body** → env-var names. After rendering and scaffold composition, the dispatcher string-replaces each key with the named env var's value. Only the allow-list is honored: `LAUNCHDARKLY_SDK_KEY`, `LAUNCHDARKLY_FLAG_KEY`, `LAUNCHDARKLY_MOBILE_KEY`, `LAUNCHDARKLY_CLIENT_SIDE_ID`. |

Merge order for the harness env: scaffold's `validation.env` first, then wrappee's `validation.env` overrides. Inputs of type `flag-key`/`sdk-key`/`mobile-key`/`client-side-id` are walked before any Docker build and fail fast if the mapped env var is unset — real keys never live in the repo.

### `validation.checks` — multiple checks per snippet

The legacy shape (one runtime check derived from the top-level fields) covers most snippets. When a snippet needs more than one validator pass — say a parse-only check today and a runtime check after we wire it up — declare an explicit list:

```yaml
validation:
  scaffold: my-sdk/scaffolds/my-syntax-only   # parent default; both checks inherit
  checks:
    - kind: parse        # uses parent's scaffold + runtime
    - kind: runtime
      scaffold: my-sdk/scaffolds/init-runner  # overrides parent for this check only
      placeholders:
        YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
```

Each `Check` has a required `kind` plus the same optional fields as the top-level `validation` block (`scaffold`, `runtime`, `entrypoint`, `companions`, `requirements`, `scaffold-inputs`, `env`, `placeholders`). Fields not set on a Check inherit from the parent `validation`. `env` and `placeholders` merge key-by-key (Check overrides win on conflicts).

Recognized `kind` values today:

- **`parse`** — language's built-in syntax check (`node --check`, `php -l`, `./gradlew compileDebugKotlin`, etc.). No LD env credentials required — the dispatcher skips the `requireEnvForInputs` gate so a CI cell can run parse-checks without provisioning a key.
- **`typecheck`** — stronger compile-or-type-check pass (`tsc --noEmit`, `dart analyze --fatal-warnings`, `xcodebuild build`, `cargo check`). Languages whose `parse` already runs a typed compiler may treat this as a no-op alias.
- **`runtime`** — full end-to-end execution against a real LaunchDarkly environment, asserting on the EXAM-HELLO `feature flag evaluates to true` line. This is what every legacy snippet has been doing.

The dispatcher forwards `SNIPPET_CHECK=<kind>` to the harness via env, so `validators/languages/<runtime>/harness/run.sh` can switch on it. Harnesses that don't recognize a kind should exit non-zero with `unsupported check kind: <kind>` rather than silently falling back.

### `placeholders` vs `env`

These are easy to confuse:

- `placeholders` rewrites literal text **inside the rendered snippet body** (`'YOUR_SDK_KEY'` → the real key) before validation. Allow-list only. Consumers still see the literal placeholder.
- `env` sets process env vars on the harness (no body rewrite). Used to switch harness behavior per wrappee (`INSTALL_KIND: podfile`, `SNIPPET_MODE: flag-eval`).

## Templating DSL

Defined in `internal/render/template.go` and `render.go`. Intentionally tiny:

- `{{ name }}` — substitute the value of a declared input. Names match `[a-zA-Z][a-zA-Z0-9_]*`. Whitespace inside the braces is flexible; the parser captures the exact source form so foreign-template syntax round-trips byte-identically.
- `{{ name | filter }}` — substitute with a filter. The only filter today is `camelCase` (react-client-sdk uses it so `useFlags()` destructuring works on a kebab-cased flag key). Any other filter is a parse error.
- `{{ if name }}…{{ end }}` — emit the inner content only if `name` resolves to non-empty. Conditionals do not nest; the inner content may still contain `{{ name }}` substitutions. The conditional's `name` must be a declared input — undeclared names in a conditional are a render error.
- **Foreign-template passthrough**: any `{{ name }}` whose name is NOT in the declared inputs map is emitted verbatim, preserving surrounding whitespace. This is what lets Vue's mustache syntax and `_shared/cursor-prompt`'s literal `{{SDK_NAME}}` (no inner whitespace — gonfalon's runtime regex requires exactly that form) survive validation untouched.
- `{{ body }}` slot in scaffolds is just `{{ name }}` where `name == "body"`. The scaffold declares `inputs: { body: { type: string, description: ... } }` and the validator substitutes the wrappee's rendered body at that point.

### Render modes

The same body renders three ways without authors picking a mode:

| Mode | Producer | What it produces |
|---|---|---|
| `runtime` | `RenderRuntime` | Plain code with concrete input values substituted. The validator runs this. |
| `ld-application` | `RenderForLDApplicationTemplate` | JS template-literal with `${name}` interpolations and `${name ? \`…\` : ''}` ternaries. Used by gonfalon. |
| `JSX text` | `RenderForJSXText` | Same shape as `ld-application` but for JSX text contexts. |

`verify` renders identically to `ld-application` and compares byte-for-byte against the consumer file.

## Common shapes

**Install (shell-install runtime):**

```markdown
---
id: react-client-sdk/sdk-info/install-yarn
sdk: react-client-sdk
kind: install
lang: shell
file: react-client-sdk/install-yarn.txt
description: Install command for react-client-sdk (yarn).
validation:
  runtime: shell-install
---

```shell
yarn add @launchdarkly/react-sdk
```
```

The `ios-install` variant carries an `env:` discriminator because Podfile / Cartfile / Package.swift fragments aren't shell-sniffable:

```yaml
validation:
  runtime: ios-install
  env:
    INSTALL_KIND: podfile
```

**Reference (wrapped by a scaffold):** no `file:`, no runtime/entrypoint/requirements/companions — the scaffold owns those.

```markdown
---
id: ios-client-sdk/sdk-docs/evaluate-a-flag-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: 'Swift in section "Evaluate a flag"'
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
let client = LDClient.get()!
```
```

**Runnable init (scaffold + placeholders):** body is unchanged from what consumers see; substitution happens at validate time only.

```markdown
---
id: ios-client-sdk/sdk-info/init
sdk: ios-client-sdk
kind: init
lang: swift
file: ios-client-sdk/init.txt
description: Client initialization snippet for ios-client-sdk.
validation:
  scaffold: ios-client-sdk/scaffolds/init-runner
  placeholders:
    YOUR_MOBILE_KEY: LAUNCHDARKLY_MOBILE_KEY
---

```swift
import LaunchDarkly

let config = LDConfig(mobileKey: "YOUR_MOBILE_KEY", autoEnvAttributes: .enabled)
// ...
LDClient.start(config: config, context: context, startWaitSeconds: 5) { timedOut in
    // ...
}
```
```

## Render targets

A snippet's rendered body reaches consumers via one of three adapters:

| Target | Consumer | How it's wired |
|---|---|---|
| `ld-application` | gonfalon (ld-application) | `SDK_SNIPPET:RENDER` markers around JSX components; the adapter rewrites the marked element's children. |
| `ld-docs` | `ld-docs`, `ld-docs-private` | `<!-- SDK_SNIPPET:RENDER id="..." -->` markers in MDX; the adapter rewrites the fenced code block underneath each marker. |
| `raw-files` | gonfalon `packages/sdk-info/` (Vite `?raw` imports) | A YAML manifest of `(snippet-id, output-path)` pairs; the adapter emits each rendered body as a standalone file. |

This repo is the canonical home of every snippet body. Consumers carry only markers (or, for `raw-files`, the manifest YAML); never edit snippet content in a consumer repo.

### Marker syntax (ld-application / ld-docs)

The renderer never touches files outside a marker. Mark each generated region with one comment of the host syntax, immediately before the JSX element (or fenced code block, for MDX) whose body should be replaced:

- TS/JS expression context: `// SDK_SNIPPET:RENDER:<id> hash=<h> version=<v>`
- JSX children context: `{/* SDK_SNIPPET:RENDER:<id> hash=<h> version=<v> */}`
- Block comment anywhere: `/* SDK_SNIPPET:RENDER:<id> hash=<h> version=<v> */`
- MDX: `<!-- SDK_SNIPPET:RENDER id="<id>" hash=<h> version=<v> -->`

`hash` and `version` are filled in by `snippets render`. On first wiring, use `hash=0` as a placeholder — the next render rewrites it.

`hash` is load-bearing: it's recomputed from the rendered children on every run, and `verify` rejects any drift. `version` is informational — it records the binary version that last produced the children, and only advances when the rendered body actually changes (so a release without content changes doesn't churn every marker).

For `ld-application` markers, the element directly following the marker MUST be a capitalized JSX component tag (e.g. `<Snippet>`, `<CodeBlock>`). Lowercase HTML tags (`<pre>`, `<code>`) are not recognized — wrap them in a component first. The marker hash covers only the *children* between `>` and `</`; attributes (`lang`, `withCopyButton`, `label`, `className`, …) are the consumer's to choose and can be edited without re-rendering.

### Manifest syntax (raw-files)

```yaml
out: ./generated
entries:
  - id: python-server-sdk/sdk-info/install-pip
    path: python-server-sdk/install-pip.txt
  - id: python-server-sdk/sdk-info/init
    path: python-server-sdk/init.txt
```

Each `path` is resolved against `<consumer-root>/<out>/`. `--consumer` defaults to the manifest's directory; pass it explicitly when the manifest lives outside the consumer checkout.

## CLI essentials

Run all commands from the `snippets/` directory (where `go.mod` lives).

```sh
# Validate one snippet end-to-end against a real LD environment
set -a; source /tmp/ld-keys.env; set +a
go run ./cmd/snippets validate \
  --sdk=python-server-sdk \
  --snippet=python-server-sdk/sdk-info/init \
  --sdks=./sdks --validators=./validators

# Validate everything in one SDK
go run ./cmd/snippets validate --sdk=python-server-sdk --sdks=./sdks

# Filter by group
go run ./cmd/snippets validate --sdk=node-server-sdk --group=sdk-info \
  --sdks=./sdks --validators=./validators

# Re-render a consumer checkout from working-tree snippets
go run ./cmd/snippets render --target=ld-docs \
  --entrypoint=../ld-docs/src/content/sdk --sdks=./sdks

# Render via manifest (raw-files consumers)
go run ./cmd/snippets render --target=raw-files \
  --manifest=../gonfalon/packages/sdk-info/snippets.yaml --sdks=./sdks

# Read-only drift check against a consumer
go run ./cmd/snippets verify --target=ld-docs \
  --entrypoint=../ld-docs/src/content/sdk --sdks=./sdks
```

`/tmp/ld-keys.env` is the convention for local LD keys; it must export `LAUNCHDARKLY_SDK_KEY`, `LAUNCHDARKLY_FLAG_KEY`, and (where the SDK requires them) `LAUNCHDARKLY_MOBILE_KEY` / `LAUNCHDARKLY_CLIENT_SIDE_ID`. Missing keys produce a clear `snippet X input Y (type=...) requires LAUNCHDARKLY_... to be set` error before any Docker build kicks off.

`--sdks` defaults to a snapshot embedded in the binary; pass `--sdks=./sdks` while iterating so working-tree edits take effect immediately. Released binaries can be used without checking out `sdk-meta`.

Docker must be on `PATH` for any non-native runtime. First-run latency is dominated by the per-language Dockerfile pre-bake; subsequent runs hit the image cache.

For the full flag matrix on `validate` / `render` / `verify`, plus harness internals, see [../CLAUDE.md#cli-reference](../CLAUDE.md#cli-reference) and [../CLAUDE.md#validator-harnesses](../CLAUDE.md#validator-harnesses).
