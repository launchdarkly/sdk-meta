# feature-docs evaluating port notes

Notes for porting the `/sdk/features/evaluating` (Evaluating flags) MDX page
(in launchdarkly/ld-docs-private) onto sdk-meta canonical snippets. Follows the
pattern established for the configuration section (see
`_feature-docs-config-port-notes.md`).

Source page (1):

| MDX page | URL | Code blocks | SDKs |
|---|---|---|---|
| `evaluating.mdx` | `/sdk/features/evaluating` | 45 | ~28 (client + server + edge) |

"Evaluation reasons" (`evaluation-reasons.mdx`) is a separate page and is NOT in
scope here.

## Group + ID layout

Snippets land under a nested folder beneath each SDK's existing `sdk-docs` group:

```
sdks/<sdk-id>/snippets/sdk-docs/features/evaluating/<slug>.snippet.md
```

Resulting IDs look like `go-server-sdk/sdk-docs/features/evaluating/evaluating-v6`.
The second ID segment is `sdk-docs`, so CI's existing per-SDK rows (which run
`validate` without a `--group` filter) pick these up alongside the config and
landing-page fragments — no new CI rows needed.

These nested feature-page snippets are deliberately separate from the pre-existing
flat `sdk-docs/evaluate-*` snippets, which belong to the per-SDK reference pages
(`server-side|client-side/<sdk>/index.mdx`) and show slightly different examples.
The flat snippets were left untouched.

## Slug conventions

Each `<CodeBlock title='…'>` becomes one snippet; the slug is the page name
(`evaluating`) plus a suffix capturing whatever distinguishes the block from its
accordion siblings:

| Distinguisher | Slug suffix |
|---|---|
| Single block for the SDK | (none — slug is `evaluating`) |
| Language tab (Java/Kotlin, Swift/Objective-C) | `-java`, `-kotlin`, `-swift`, `-objc` |
| TypeScript vs JavaScript sibling | `-ts` / `-js` |
| SDK version | `-v6`, `-v4`, `-v3`, `-v2`, `-v1x`, `-v7-scopedclient` |
| C++ native / C-binding / C SDK v2.x (+ C++ binding) | `-cpp-native-v3-0`, `-cpp-c-v3-0`, `-c-sdk-v2`, `-c-sdk-v2-cpp` |
| React Native hooks vs methods | `-hooks`, `-methods` |
| Electron all-flags companion block | `-all-flags` |
| JS SDK v4.x typed-method block | `-v4` |

## Snippet shape

All `kind: reference` snippets that declare only `validation.scaffold:` (no
runtime/entrypoint/requirements — the scaffold owns those), bound to each SDK's
existing `*-syntax-only` scaffold. Bodies reference `client`/`context`/`user`
which the scaffolds stub.

## Scaffold stub extensions made here

The evaluation bodies pass a bare `context` (or `user`) to the variation methods;
several syntax-only scaffolds previously stubbed only `client` (the config bodies
never needed `context`). Extended (additive, harmless to existing bodies):

- `java-server-sdk/scaffolds/java-syntax-only` — added an `LDContext context` stub.
- `haskell-server-sdk/scaffolds/haskell-syntax-only` — added a `context :: Context` stub.
- `haskell-server-sdk/scaffolds/haskell-syntax-only-v3` — added a `user :: User`
  stub (the v3.x SDK's variation functions take `User`).

## Sample fixes (corrected versus the source MDX)

- **Erlang (`evaluating-v2`, `evaluating-v1x`)**: the source map literal had a
  trailing comma (`#{key => <<"…">>,}`), which is a syntax error in Erlang. Removed
  the trailing comma so the example compiles.

## Validator coverage / bind rate

Of the 45 blocks, 44 are bound to a validator. The one unbound block:

- **iOS Objective-C (`evaluating-objc`)**: there is no Objective-C parse scaffold
  (the iOS validator is the macOS-only native xcodebuild harness, and the existing
  `swift-syntax-only` scaffold is Swift-only). iOS Swift (`evaluating-swift`) IS
  bound to `swift-syntax-only` and validated on the macOS CI row.

Local note: `dotnet-server`/`dotnet-client` syntax-only scaffolds compile against
the LaunchDarkly AI add-on; a local NuGet restore mismatch can surface a
`LaunchDarkly.Sdk.Server.Ai.DataModel` namespace error that does not occur in CI
(the config dotnet snippets exhibit the same locally yet pass CI). These are
validated in CI.

## Out of scope / follow-ups

- `evaluation-reasons.mdx` — separate page, natural next section.
