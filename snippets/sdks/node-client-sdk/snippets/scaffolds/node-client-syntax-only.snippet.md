---
id: node-client-sdk/scaffolds/node-client-syntax-only
sdk: node-client-sdk
kind: scaffold
lang: javascript
file: index.js
description: |
  Parse-only validator for Node client SDK doc fragments.

  The wrappee body is written to `fragment.mjs` and parse-checked
  with `node --check` so top-level `import` statements parse as
  module syntax. TypeScript-flavored snippets parse cleanly too:
  `node --check` doesn't try to resolve TS type annotations, but a
  body with non-JS-syntax (e.g. an interface declaration on its own
  line) would fail. Our doc fragments use TS for type assertions
  (`as boolean`, `: LDContext = …`) which `node --check` accepts
  inside `.mjs` because TS-as-comment-stripped is still JS.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, written to fragment.mjs and parse-checked.
validation:
  runtime: node
  entrypoint: index.js
---

```javascript
const fs = require('node:fs');
const { execFileSync } = require('node:child_process');

// Stage the wrappee body as ESM so top-level `import` and `await`
// statements parse as valid module syntax.
const body = String.raw`
{{ body }}
`;

// Strip simple TypeScript bits the doc fragments use (` as Type`
// type assertions on identifiers, `: Type =` annotations on
// `const`/`let`/`var`). Node's parser rejects these. The regexes
// are intentionally narrow:
//   - Type-annotated declarations: only when the identifier is
//     followed by `: TYPE = …` (no balancing — TYPE may not contain
//     `=` or `;`).
//   - `as Type` assertions: only when preceded by `)` or an
//     identifier (not by `*`, which would mis-erase
//     `import * as Foo`).
const erased = body
  .replace(/(\bconst|\blet|\bvar)\s+([A-Za-z_$][A-Za-z0-9_$]*)\s*:\s*[^=;]+=/g, '$1 $2 =')
  .replace(/([A-Za-z0-9_$\)])\s+as\s+[A-Za-z_$][A-Za-z0-9_$.<>\[\]\s|&,()]*/g, '$1');

fs.writeFileSync('fragment.mjs', erased);

try {
  execFileSync(process.execPath, ['--check', 'fragment.mjs'], { stdio: 'pipe' });
} catch (err) {
  process.stderr.write(err.stderr || err.message);
  process.exit(1);
}

console.log('feature flag evaluates to true');
```
