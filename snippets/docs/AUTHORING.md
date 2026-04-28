# Authoring snippets

A snippet is a single Markdown file with YAML frontmatter and one fenced code
block. See `sdks/python-server-sdk/snippets/getting-started/` for the canonical
examples.

## Frontmatter fields used in the first slice

| Field | Required | Notes |
|---|---|---|
| `id` | yes | globally unique; convention `<sdk>/<group>/<name>` |
| `sdk` | yes | matches a directory under `sdks/` |
| `kind` | yes | one of: `bootstrap` (project-scaffolding command, e.g. `npx create-vue@latest`), `install` (dependency-add command, e.g. `pip install ldclient-py`), `hello-world` (the actual code file the consumer drops in), `run` (the command that runs the result, e.g. `python main.py`). One snippet has exactly one kind; the set is closed for now. New kinds get added here as the snippet system grows. |
| `lang` | yes | language tag for the fenced code block |
| `description` | recommended | one-liner shown to the consumer |
| `inputs` | optional | each input has `type`, `description`, `runtime-default` |
| `ld-application.slot` | optional | logical slot label for the ld-application adapter |
| `validation.entrypoint` | optional | filename to run inside the validator (e.g. `main.py`) |
| `validation.requirements` | optional | line written into `requirements.txt` |

## Templating

Inside the code block:

- `{{ name }}` — substitutes an input value
- `{{ if name }}...{{ end }}` — emits the inner content only when `name` has a non-empty value

Conditionals do not nest. The inner content of a conditional may still contain `{{ name }}`.

## Render modes

Same template renders three different ways:

| Mode | What it produces |
|---|---|
| `ld-application` | JS template-literal expression with `${name}` interpolations and `${name ? \`...\` : ''}` ternaries inside JSX. |
| `runtime` (validator) | Plain code with concrete values substituted. |
| `verify` | Same as `ld-application`, then compared byte-for-byte against the consumer file. |

## Render markers in consumer files

The generator never edits files outside a marker comment. Mark each generated
region with one comment of the host syntax, immediately before the JSX element
whose body should be replaced:

- TS/JS expression context: `// SDK_SNIPPET:RENDER:<id> hash=<h> version=<v>`
- JSX children context: `{/* SDK_SNIPPET:RENDER:<id> hash=<h> version=<v> */}`
- Block comment anywhere: `/* SDK_SNIPPET:RENDER:<id> hash=<h> version=<v> */`

`hash` and `version` are filled in by `snippets render`. On first wiring, use
`hash=0` as a placeholder — the next render rewrites it. (Any hex string
works; `0` just minimizes finger-typing.)

The element directly following a marker MUST be a capitalized JSX component
tag (e.g. `<Snippet>`, `<CodeBlock>`). Lowercase HTML tags (`<pre>`, `<code>`)
are not recognized; wrap the content in a component first if you need to mark
it. The marker hash covers only the *children* of the element (between `>`
and `</`), matching the `scope=content` contract — attributes (`lang`,
`withCopyButton`, `label`, `className`, …) are the consumer's to choose and
can be edited without re-running `snippets render`.

## CLI quick reference

```sh
# Validate a snippet end-to-end in Docker, against a real LD environment
export LAUNCHDARKLY_SDK_KEY=...     # server-side key
export LAUNCHDARKLY_FLAG_KEY=...    # flag the snippet evaluates
snippets validate --sdk=python-server-sdk

# Rewrite all marked regions in an ld-application checkout
snippets render --target=ld-application --out=/path/to/ld-application

# Confirm consumer file matches what we'd render (no edits)
snippets verify --target=ld-application --out=/path/to/ld-application
```

## Validator inputs

Validation runs against a real LaunchDarkly environment. The caller supplies:

| Env var | Mapped to inputs of `type` |
|---|---|
| `LAUNCHDARKLY_SDK_KEY` | passed to the snippet via the same env var inside the Docker container; the snippet reads it the same way the `hello-*` sample apps do |
| `LAUNCHDARKLY_FLAG_KEY` | substituted into any input of `type: flag-key` at render-for-validation time |

Snippet frontmatter does **not** carry default values for `flag-key` / `sdk-key` /
`mobile-key` / `client-side-id` types — those always come from the caller's
environment so real keys never end up committed to this repo.
