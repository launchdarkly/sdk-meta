---
id: node-server-sdk/sdk-docs/features/hooks/define-hook
sdk: node-server-sdk
kind: reference
lang: ts
description: Hook implementation and configuration for Node.js (server-side).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only-toplevel
---

```ts
export class ExampleHook implements integrations.Hook {

  getMetadata() {
    return { name: 'Example hook'}
  }

  // Implement at least one of `beforeEvaluation`, `afterEvaluation`

  // `beforeEvaluation` is called during the execution of a variation method
  // before the flag value has been determined

  // `afterEvaluation` is called during the execution of a variation method
  // after the flag value has been determined
}

const options: ld.LDOptions = {
  hooks: [new ExampleHook()]
};

const client = ld.init('YOUR_SDK_KEY', options);
```
