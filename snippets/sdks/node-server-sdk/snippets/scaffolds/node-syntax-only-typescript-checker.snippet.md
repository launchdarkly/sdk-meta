---
id: node-server-sdk/scaffolds/node-syntax-only-typescript-checker
sdk: node-server-sdk
kind: scaffold
lang: javascript
file: index.js
description: |
  Checker half of the `node-syntax-only-typescript` scaffold pair.
  Staged as a companion so it lands at `index.js` while the wrappee
  body sits verbatim at `fragment.ts`.
---

```javascript
const fs = require('node:fs');
const { stripTypeScriptTypes } = require('node:module');
const { execFileSync } = require('node:child_process');

// Remove the type syntax with the TypeScript stripper in Node. The
// stripper replaces each type with spaces, so the line and column
// numbers in an error message match the fragment. The stripper throws
// on invalid syntax and on TypeScript that it cannot erase. Node then
// prints the location of the error and exits with status 1.
const js = stripTypeScriptTypes(fs.readFileSync('fragment.ts', 'utf8'), {
  sourceUrl: 'fragment.ts',
});

// Parse the result as an ES module. Do not run `--check` on the `.ts`
// file directly. When a `.ts` or `.mts` file has an `import` or
// `export` statement, Node 22 does not parse it and exits with 0.
fs.writeFileSync('fragment.mjs', js);
try {
  execFileSync(process.execPath, ['--check', 'fragment.mjs'], { stdio: 'pipe' });
} catch (err) {
  process.stderr.write(err.stderr || err.message);
  process.exit(1);
}

console.log('feature flag evaluates to true');
```
