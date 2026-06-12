---
id: node-server-sdk/scaffolds/node-syntax-only-toplevel-checker
sdk: node-server-sdk
kind: scaffold
lang: javascript
file: index.js
description: |
  Checker half of the `node-syntax-only-toplevel` scaffold pair.
  Staged as a companion so it lands at `index.js` while the wrappee
  body sits verbatim at `fragment.ts`. Reading the fragment from disk
  (instead of embedding it in a template literal the way
  `node-syntax-only` does) keeps bodies containing backticks
  parseable. The TS-erasure regexes mirror `node-syntax-only`, plus an
  `implements X` clause eraser for class-implementation fragments;
  keep the shared parts in sync.
---

```javascript
const fs = require('node:fs');
const { execFileSync } = require('node:child_process');

const body = fs.readFileSync('fragment.ts', 'utf8');

// Strip simple TS bits (`as Type` assertions, `: Type =`
// declarations, `implements X` clauses on class declarations) before
// the parser sees them. The first two regexes mirror node-syntax-only;
// see the node-client scaffold for the regex shape rationale. The
// `implements` eraser is intentionally narrow: it only fires when the
// clause is followed by the class body's opening brace.
const erased = body
  .replace(/(\bconst|\blet|\bvar)\s+([A-Za-z_$][A-Za-z0-9_$]*)\s*:\s*[^=;]+=/g, '$1 $2 =')
  .replace(/([A-Za-z0-9_$\)])\s+as\s+[A-Za-z_$][A-Za-z0-9_$.<>\[\]\s|&,()]*/g, '$1')
  .replace(/\bimplements\s+[A-Za-z_$][A-Za-z0-9_$.]*(\s*,\s*[A-Za-z_$][A-Za-z0-9_$.]*)*\s*\{/g, '{');

fs.writeFileSync('fragment.mjs', erased);

try {
  execFileSync(process.execPath, ['--check', 'fragment.mjs'], { stdio: 'pipe' });
} catch (err) {
  process.stderr.write(err.stderr || err.message);
  process.exit(1);
}

console.log('feature flag evaluates to true');
```
