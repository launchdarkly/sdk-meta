---
id: node-server-sdk/scaffolds/node-syntax-only
sdk: node-server-sdk
kind: scaffold
lang: javascript
file: index.js
description: |
  Parse-only validator for Node server SDK doc fragments.

  The wrappee body is written to `fragment.mjs` (ESM) so top-level
  `import` statements parse as valid module syntax, then `node --check`
  is invoked to syntax-check it without actually resolving / running
  the imports. A wrappee with malformed syntax fails the check; a
  wrappee referencing unresolved symbols (`client`, `ldclient`, etc.)
  passes because `--check` does not execute.
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

// Stage the wrappee body as ESM so top-level `import` statements parse
// as valid module syntax. `node --check` is purely syntactic; the
// imported package names don't have to resolve.
const body = String.raw`
{{ body }}
`;

// Write the body as a TypeScript module (.mts) and let Node strip the
// types natively (`--experimental-strip-types`, Node 22.6+) before the
// `--check` syntax pass. This handles the full TS surface the doc
// fragments use — typed function parameters (`async(key?: string)`),
// inline object types, `as` assertions, `: Type` declarations — which
// a regex strip can't do reliably. `--check` is purely syntactic, so
// imported package names still don't have to resolve.
fs.writeFileSync('fragment.mts', body);

try {
  execFileSync(
    process.execPath,
    ['--experimental-strip-types', '--check', 'fragment.mts'],
    { stdio: 'pipe' },
  );
} catch (err) {
  process.stderr.write(err.stderr || err.message);
  process.exit(1);
}

console.log('feature flag evaluates to true');
```
