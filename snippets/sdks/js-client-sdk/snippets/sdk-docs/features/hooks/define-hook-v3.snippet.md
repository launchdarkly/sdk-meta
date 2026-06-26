---
id: js-client-sdk/sdk-docs/features/hooks/define-hook-v3
sdk: js-client-sdk
kind: reference
lang: js
description: Hook implementation and configuration for the JavaScript SDK v3.6+.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
export class ExampleHook implements Hook {

  getMetadata() {
    return { name: 'Example hook'}
  }

  // Implement at least one of:
  //
  // * `beforeEvaluation` - called during the execution of a variation method
  // before the flag value has been determined
  //
  // * `afterEvaluation` - called during the execution of a variation method
  // after the flag value has been determined
  //
  // * `beforeIdentify` - called during the execution of the identify process
  // before the operation completes, but after any context modifications are performed
  //
  // * `afterIdentify` - called during the execution of the identify process
  // after the operation completes
  //
  // * `afterTrack` - called during the execution of the track process
  // after the event has been enqueued
}

const options = {
  hooks: [new ExampleHook()]
};

const client = LDClient.initialize('example-client-side-id', context, options);
```
