---
id: node-server-sdk/scaffolds/openfeature-node-toplevel
sdk: node-server-sdk
kind: scaffold
lang: javascript
file: index.mjs
description: |
  Resolves an OpenFeature provider doc fragment that is a set of
  top-level `import` statements. Unlike the parse-only
  `node-syntax-only-toplevel` scaffold, this one actually executes the
  module, so the imported package names must resolve against the
  OpenFeature SDK, the LaunchDarkly Node.js SDK, and the LaunchDarkly
  provider that the harness installs from `requirements`. A typo in a
  package or export name fails the import; this catches drift that a
  pure parse check would miss. The body imports do not connect to
  LaunchDarkly — only instantiating the provider does — so this runs
  without LaunchDarkly credentials.
inputs:
  body:
    type: string
    description: The wrappee snippet's import statements, staged verbatim and executed.
validation:
  runtime: node
  entrypoint: index.mjs
  requirements: |
    @openfeature/server-sdk
    @launchdarkly/node-server-sdk
    @launchdarkly/openfeature-node-server
---

```javascript
{{ body }}

console.log('feature flag evaluates to true');
```
