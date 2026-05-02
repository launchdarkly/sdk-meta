---
id: node-client-sdk/scaffolds/init-runner
sdk: node-client-sdk
kind: scaffold
lang: javascript
file: index.mjs
description: |
  Runs an `init.txt`-style snippet end-to-end against a real LaunchDarkly
  env. Same shape as the node-server scaffold: write the wrappee body to
  a sibling `.mjs` file, dynamic-import it, then assert the client
  `LaunchDarkly.initialize(...)` returned a client whose
  `waitForInitialization` call settled successfully.

  The snippet itself doesn't print on success (it has empty try/catch
  branches), so we tail-append a second `await client.waitForInitialization(5)`
  on the same lexical `client` the body declared. If init has settled
  successfully the awaiting promise resolves immediately; if init failed
  in the body's swallowed catch, the second await rejects, the
  rejection is caught here, and we exit non-zero.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: node
  entrypoint: index.mjs
  requirements: |
    launchdarkly-node-client-sdk
---

```javascript
{{ body }}

try {
  await client.waitForInitialization(5);
  console.log('feature flag evaluates to true');
  await client.close();
} catch (err) {
  console.error('scaffold: waitForInitialization rejected:', err && err.message ? err.message : err);
  process.exit(1);
}
```
