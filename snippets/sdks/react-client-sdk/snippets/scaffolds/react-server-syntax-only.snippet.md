---
id: react-client-sdk/scaffolds/react-server-syntax-only
sdk: react-client-sdk
kind: scaffold
lang: tsx
file: snippet.tsx
description: |
  Parse-only validator for React Server Component doc fragments. These
  bodies import @launchdarkly/node-server-sdk and
  @launchdarkly/react-sdk/server -- a Node server environment the
  react-client browser validator cannot provide (lifting the imports to
  module scope would execute the Node server SDK inside Chromium).
  Routes through the `edge-ts` validator instead: the TypeScript
  compiler's transpileModule does a syntax check + type-strip with no
  module resolution and no type-checking, and the `.tsx` staging
  filename selects TSX parsing so the fragments' JSX parses. The doc
  fragments are whole modules, so the scaffold emits the body verbatim.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, parsed as a TSX module.
validation:
  runtime: edge-ts
  entrypoint: snippet.tsx
---

```tsx
{{ body }}
```
