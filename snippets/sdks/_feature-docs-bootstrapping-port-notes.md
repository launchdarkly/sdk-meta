# Port notes: /sdk/features/bootstrapping

Source: `ld-docs-private` `fern/topics/sdk/features/bootstrapping.mdx`.
12 code blocks extracted into `sdk-docs/features/bootstrapping/`
snippets across 4 SDKs (Electron, JavaScript, Node.js client-side,
React Web). All but two (the JavaScript SDK's Ruby-template blocks)
are bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **Electron JavaScript** (`electron-client-sdk/.../bootstrapping-js`):
  the sample called `LaunchDarkly.initialize(...)`, but the Electron
  SDK exports only `initializeInMain` / `initializeInRenderer` (no
  `initialize`). Changed to `initializeInMain`, matching the
  TypeScript sibling block and the SDK's typings.
- **JavaScript SDK v3.x** (`js-client-sdk/.../bootstrapping-v3`) and
  **v3.x TypeScript** (`bootstrapping-v3-ts`): bare `...` elision
  lines inside the `onPageLoad` function body are a syntax error in
  JavaScript/TypeScript; converted to `// ...` comments.
- **JavaScript SDK v3.x TypeScript** (`bootstrapping-v3-ts`), also:
  the block imported `LDClient` as a named import and then called
  `LDClient.initialize(...)` — in the v3 typings `LDClient` is an
  interface (a type, not a value), so the call cannot work. Split
  into a namespace import (`import * as LDClient`) for the
  `initialize` call plus named imports for the types, matching the
  namespace-import pattern the config page uses for the same SDK.
- **Node.js (client-side) v3 JavaScript / TypeScript**
  (`node-client-sdk/.../bootstrapping-js`, `bootstrapping-ts`): same
  bare-`...` elision-line fix as the JavaScript v3 blocks.
- **Node.js (client-side) v3 TypeScript** (`bootstrapping-ts`), also:
  the block imported from `'launchdarkly-js-client-sdk'` — the
  browser SDK's package — instead of `'launchdarkly-node-client-sdk'`,
  and had the same `LDClient`-interface-as-value problem as the
  JavaScript v3 TypeScript block. Fixed the package name and split
  the import the same way.
- **JavaScript SDK v3.x Ruby template** (`bootstrapping-ruby-v3`):
  the client was assigned to `const ldclient` but every subsequent
  call (`client.waitForInitialization(5)`) used `client`. Renamed the
  declaration to `client`, matching the v3 block on the evaluation
  reasons page.

## Validation routing added in this port

- `react-client-sdk/scaffolds/react-server-syntax-only` — new
  parse-only scaffold for React Server Component fragments, routed
  through the existing `edge-ts` validator (transpileModule syntax
  check; no module resolution). The fragments import
  `@launchdarkly/node-server-sdk` and `@launchdarkly/react-sdk/server`,
  a Node server environment the react-client browser validator cannot
  provide. The scaffold stages the body as `snippet.tsx` so the
  TypeScript compiler parses JSX. Binding it here also let the two
  pre-existing react reference-page snippets that carried a
  `TODO(scaffold)` for exactly this case
  (`bootstrap-react-server-wrapper-javascript`,
  `react-server-component-support-create-a-wrapped-server-side-sdk-client-react-web-sdk-v4-0`)
  be wired up; their TODOs are resolved in this port.
- `node-client-sdk/scaffolds/node-client-syntax-only` — extended the
  TypeScript-erasure pass with a third narrow rewrite: parameter type
  annotations on lines that open a `function` declaration
  (`function onPageLoad(flags: LDFlagSet)`), which `node --check`
  rejects. Anchored to function-declaration lines so `key: value`
  pairs inside object literals are never touched.
- `electron-client-sdk/scaffolds/electron-syntax-only` — deliberately
  unchanged. The TypeScript fragment's
  `import * as LaunchDarkly from 'launchdarkly-electron-client-sdk'`
  passes through the scaffold's dead-function path as is. Lifting such
  imports to module scope (the IMPORT_LIFT triad other JS-family
  scaffolds use) breaks fragments that use the imported namespace as a
  value: the electron package is not installed in the js-client
  validator image, so the lifted import becomes an undefined global
  reference that throws in Chromium before the success line renders.
  Inside the never-executed wrappee function the import stays inert.

## Known non-binds

- `js-client-sdk/.../bootstrapping-ruby-v4` and
  `bootstrapping-ruby-v3` — the bodies embed a Ruby ERB template
  directive (`<%= client.all_flags_state(...).to_json %>`) inside the
  JavaScript, so they are not parseable by any JavaScript-family
  validator. Wiring them up needs a mixed-host harness that
  pre-renders (or stubs out) the ERB directive before handing the
  remainder to a JS parser. These snippets are byte-equality-verified
  through marker hashes, which is the strongest check available for a
  fragment without a runnable host language — the same blocker class
  as the bare shell env-var fragments documented in
  `_sdk-docs-port-notes.md`.
