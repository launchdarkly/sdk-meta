---
id: node-server-sdk/scaffolds/node-syntax-only-typescript
sdk: node-server-sdk
kind: scaffold
lang: ts
file: fragment.ts
description: |
  Parse-only validator for Node server SDK doc fragments that use
  TypeScript syntax that the regex erasure in `node-syntax-only` does
  not remove, such as `import type` declarations and typed function
  signatures.

  This variant stages the body verbatim as `fragment.ts` (this file).
  The `node-syntax-only-typescript-checker` companion removes the type
  syntax with the TypeScript stripper in Node
  (`module.stripTypeScriptTypes`, Node 22.13 and later). It then runs
  `node --check` on the result as ESM. Invalid TypeScript fails the
  strip step. Syntax that the stripper cannot erase, such as `enum` and
  `namespace` declarations, also fails. `--check` does not run the
  code, so unresolved names (`aiConfig`, `awsClient`) pass.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, staged verbatim as fragment.ts.
validation:
  runtime: node
  entrypoint: index.js
  companions:
    - node-server-sdk/scaffolds/node-syntax-only-typescript-checker
---

```ts
{{ body }}
```
