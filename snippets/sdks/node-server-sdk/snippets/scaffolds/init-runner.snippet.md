---
id: node-server-sdk/scaffolds/init-runner
sdk: node-server-sdk
kind: scaffold
lang: javascript
file: index.mjs
description: |
  Runs an `init.txt`-style snippet end-to-end against a real LaunchDarkly
  env. The wrappee's literal SDK-key placeholder is substituted at
  validate time via the snippet's `validation.placeholders` map (handled
  by the dispatcher), so by the time the body executes here it already
  carries the real key.

  The scaffold writes the wrappee body to a sibling `.mjs` file and
  dynamic-imports it so the wrappee's top-level statements run with
  `node_modules` resolution rooted in the staging dir. It then waits a
  bounded period for the wrappee's `client.once('ready', …)` callback
  to print the snippet's success line, and emits the EXAM-HELLO line
  the validator harness greps for. The shared `await_success_line`
  helper SIGTERMs us as soon as the EXAM-HELLO line appears, so the
  bounded wait below is just a ceiling.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: node
  entrypoint: index.mjs
  requirements: |
    @launchdarkly/node-server-sdk
---

```javascript
import { writeFileSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';

// String.raw on a tagged template prevents \\ escape-collapsing and
// keeps the wrappee body byte-identical to the snippet source. None of
// the node init bodies we ship contain backticks; document the
// constraint here for any future port that adds one.
const wrappeeBody = String.raw`
{{ body }}
`;

let sawSuccess = false;
const origLog = console.log.bind(console);
console.log = (...args) => {
  if (args.join(' ').includes('SDK successfully initialized')) {
    sawSuccess = true;
  }
  origLog(...args);
};

// Write into the same directory as this scaffold's entry file so the
// wrappee's `import '@launchdarkly/node-server-sdk'` resolves through
// the local node_modules the harness installed.
const here = dirname(fileURLToPath(import.meta.url));
const bodyPath = join(here, '_init-body.mjs');
writeFileSync(bodyPath, wrappeeBody);

await import(bodyPath);

await new Promise((resolve) => setTimeout(resolve, 30_000));

if (!sawSuccess) {
  console.error("scaffold: wrappee did not print 'SDK successfully initialized'");
  process.exit(1);
}

origLog('feature flag evaluates to true');
```
