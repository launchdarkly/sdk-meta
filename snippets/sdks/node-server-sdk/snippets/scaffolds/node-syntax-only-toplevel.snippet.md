---
id: node-server-sdk/scaffolds/node-syntax-only-toplevel
sdk: node-server-sdk
kind: scaffold
lang: ts
file: fragment.ts
description: |
  Parse-only validator for Node server SDK doc fragments that must be
  staged as a standalone module — class/interface declarations and
  bodies containing backtick characters (for example Markdown-style
  code references inside comments). `node-syntax-only` embeds the
  wrappee body inside a JS template literal, so a body containing a
  backtick terminates the literal and breaks the scaffold itself.

  This variant stages the body verbatim as `fragment.ts` (this file)
  and ships the checker program as the `node-syntax-only-toplevel-checker`
  companion, which reads the fragment from disk, erases the simple
  TypeScript surface the docs use (type-annotated declarations,
  `as Type` assertions, `implements X` clauses), and runs
  `node --check` on the result as ESM. `--check` never executes the
  code, so unresolved names (`ld`, `integrations`, `client`) pass.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, staged verbatim as fragment.ts.
validation:
  runtime: node
  entrypoint: index.js
  companions:
    - node-server-sdk/scaffolds/node-syntax-only-toplevel-checker
---

```ts
{{ body }}
```
