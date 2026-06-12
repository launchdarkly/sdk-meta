# Port notes: /sdk/features/inspectors

Source: `ld-docs-private` `fern/topics/sdk/features/inspectors.mdx`.
1 code block extracted into `sdk-docs/features/inspectors/` (the page's
React Native and React Web sections defer to the JavaScript sample and
carry no code blocks of their own). Bound to a validator.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX.

- **JavaScript Flag Used inspector** (`js-client-sdk/.../flag-used`):
  the published sample passed the options as `options: { ... }` inside
  the `LDClient.initialize(...)` argument list — labeled arguments are
  not JavaScript syntax, so the block could never parse. The third
  argument is the options object itself; the `options:` label was
  dropped so the inspectors configuration is passed directly.

## Validation routing added in this port

None. The snippet binds to the existing
`js-client-sdk/scaffolds/js-syntax-only` scaffold.

## Known non-binds

None.
