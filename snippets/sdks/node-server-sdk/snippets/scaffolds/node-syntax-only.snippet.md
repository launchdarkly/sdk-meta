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

// Strip simple TS bits (`as Type` assertions, `: Type =`
// declarations) before the parser sees them. Same approach as the
// node-client scaffold; see comment there for regex shape rationale.
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
