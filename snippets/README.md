# sdk-snippets

Single source of truth for LaunchDarkly SDK code snippets rendered into the
LD application UI and the LD docs site.

First slice: `python-server-sdk` "Getting started" flow.

## Layout

```
cmd/snippets/          CLI entrypoint (render, verify, validate)
internal/              generator Go code
sdks/<id>/             one directory per SDK
  sdk.yaml             SDK descriptor
  snippets/            .snippet.md files (YAML frontmatter + Markdown + code blocks)
validators/languages/  per-language Docker validator harnesses
```

## CLI

```
snippets render   --target=ld-application --out=<app-checkout>
snippets verify   --target=ld-application --out=<app-checkout>
snippets validate --sdk=python-server-sdk
```

`render` rewrites the code regions between `SDK_SNIPPET:RENDER` markers.
`verify` recomputes hashes and fails if any managed region has drifted.
`validate` runs each snippet in a Docker container against a real LaunchDarkly
environment. Required env vars (matching the convention used by the `hello-*`
sample apps):

```sh
export LAUNCHDARKLY_SDK_KEY=<server-side SDK key>
export LAUNCHDARKLY_FLAG_KEY=<flag key the snippet evaluates>
snippets validate --sdk=python-server-sdk
```

Neither variable is committed to this repo. They are forwarded into the per-
snippet Docker run.

## Adapter targets

| Target | Renders into |
|---|---|
| `ld-application` | the LaunchDarkly application's "Get Started" flows (TSX/JSX) |
| `ld-docs` | the LaunchDarkly documentation site (MDX) — not yet implemented |

The adapter is selected via `--target`. Today only `ld-application` is wired
up; `ld-docs` will land in a later slice.
