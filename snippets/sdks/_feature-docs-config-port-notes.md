# feature-docs config port notes

Notes for porting the `/sdk/features/config/*` MDX pages (in
launchdarkly/ld-docs-private) onto sdk-meta canonical snippets.

Source pages (4):

| MDX page | URL | Code blocks | SDKs |
|---|---|---|---|
| `migration-config.mdx` | `/sdk/features/migration-config` | 11 | 11 (server + edge) |
| `app-config.mdx` | `/sdk/features/app-config` | 32 | 21 |
| `index.mdx` | `/sdk/features/config` | 47 | 27 |
| `service-endpoint-configuration.mdx` | `/sdk/features/service-endpoint-configuration` | 99 | 24 |

This file is the analogue of `_sdk-docs-port-notes.md` and
`_sdk-info-port-notes.md`.

## Group + ID layout

Snippets land under a new nested folder beneath each SDK's existing
`sdk-docs` group:

```
sdks/<sdk-id>/snippets/sdk-docs/features/config/<slug>.snippet.md
```

Resulting snippet IDs look like
`python-server-sdk/sdk-docs/features/config/migration-config`. The
second segment of the ID is `sdk-docs`, so CI's existing
`--group=sdk-docs` matrix row picks these up alongside the per-SDK
landing-page fragments — one validation rotation covers both
families.

Per-CodeBlock naming (because each marker maps to exactly one
fence; the MDX adapter does not handle multi-CodeBlock accordions
under a single marker):

| Source shape | Slug convention |
|---|---|
| One CodeBlock per accordion (e.g. all of `migration-config.mdx`) | `<page>` (e.g. `migration-config`) |
| Multiple CodeBlocks for SDK version variants (e.g. .NET v4.0 / v3.0 in `index.mdx`) | `<page>-v<N>-<N>` (e.g. `index-v4-0`, `index-v3-0`) |
| Multiple CodeBlocks for language flavors (e.g. Android Java / Kotlin) | `<page>-<lang>` (e.g. `index-java`, `index-kotlin`) |
| Multiple CodeBlocks for native / C-binding variants (C++ SDKs) | `<page>-<binding>` (e.g. `index-native`, `index-c-binding`) |

## Snippet shape

These are all `kind: reference` snippets with no `validation:`
block — the bodies reference symbols (`_client`, `client`, `ldclient`,
the SDK's `Migration*` types) that aren't supplied in a standalone
runtime, mirroring the `kind: reference` fragments under each SDK's
existing `sdk-docs/` group.

The fence in the .snippet.md file is plain `\`\`\`<lang>`; MDX-side
presentation attributes (`maxLines=0`, etc.) stay on the consumer
side. The renderer only rewrites bytes BETWEEN fences, so the
consumer's fence-line attributes are preserved across re-renders.

## Scope (this PR)

All four `/sdk/features/config/*` pages, 189 snippet files:

| MDX page | Snippets | SDKs touched |
|---|---:|---:|
| `migration-config.mdx` | 11 | 11 |
| `app-config.mdx` | 32 | 21 |
| `index.mdx` | 47 | 27 |
| `service-endpoint-configuration.mdx` | 99 | 24 |

### New sdk-meta SDK directories created here

The snippet system needed `sdks/<id>/` directories for four edge SDKs
that weren't yet hosted. Each gets a minimal `sdk.yaml` and its first
snippets in this port:

- `akamai-server-edgekv-sdk` (referenced by all 4 pages)
- `cloudflare-server-sdk` (referenced by all 4 pages)
- `fastly-server-sdk` (referenced by `index.mdx` only)
- `vercel-server-sdk` (referenced by all 4 pages)

`akamai-server-base-sdk` (a separate Akamai package) is NOT added —
none of the four config pages import from it.

### MDX-side fix-ups landed alongside the markers

- **`service-endpoint-configuration.mdx` Electron accordion**: two
  consecutive blocks were both titled "TypeScript, connecting to
  federal instance" — but the second's body shows the EU endpoints
  (`eu.launchdarkly.com`). The duplicate title is a doc-side typo;
  this PR renames the second block to "TypeScript, connecting to EU
  instance" so it matches its body. One-line MDX change.

- **JS accordion bodies in `app-config.mdx` / `index.mdx`**: the
  original `<CodeBlock>` bodies were indented 2 spaces (visual
  nesting inside the wrapping `<CodeBlocks>`). The canonical form
  in sdk-meta is flush-left (consistent with every other SDK
  accordion in these pages), so after re-render the bodies move
  flush-left. CommonMark-valid; surrounding fence lines retain
  their 2-space indent, which the renderer doesn't touch.

## Slug conventions used

Each `<CodeBlock title='…'>` becomes one snippet. The snippet
filename / slug derives from the page name plus a suffix that
captures whatever distinguishes this block from its accordion
siblings:

| Distinguisher in title | Slug suffix |
|---|---|
| Single block in accordion | (no suffix — slug is just the page name, e.g. `migration-config`) |
| SDK version (`v4.0`, `v3.x`) | `v<N>` or `v<N>-<M>` (e.g. `v4`, `v4-0`, `v3`) |
| Language tab (Java/Kotlin, Swift/Objective-C) | `<lang>` (e.g. `java`, `kotlin`, `swift`, `objc`) |
| TypeScript vs JavaScript sibling | `ts` / `js` |
| C++ native vs C-binding vs C SDK v2.x | `cpp-native`, `cpp-c`, `c-sdk` |
| Endpoint variant (Relay Proxy / federal / EU) | `relay`, `federal`, `eu` |
| Combinations (e.g. C++ native + EU) | hyphen-joined: `cpp-native-eu` |
| Roku `index.mdx` constructor vs setters reference | `create`, `options` |
| Akamai `index.mdx` default vs custom feature store | `default`, `custom-store` |

Final snippet IDs all live under
`<sdk-id>/sdk-docs/features/config/<slug>` so CI's existing
`--group=sdk-docs` matrix row picks them up alongside the per-SDK
landing-page fragments.

## Out of scope for this PR

None — all four pages are covered. Future feature-docs ports
(`/sdk/features/evaluating`, `/sdk/features/identify`, etc.) can
follow the same pattern: new snippets land under
`<sdk>/sdk-docs/features/<topic>/<slug>` and add markers to the
relevant MDX file.
