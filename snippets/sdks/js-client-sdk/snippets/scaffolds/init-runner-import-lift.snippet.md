---
id: js-client-sdk/scaffolds/init-runner-import-lift
sdk: js-client-sdk
kind: scaffold
lang: javascript
file: src/app.ts
description: |
  Same end-to-end runner as js-client-sdk/scaffolds/init-runner, for
  init bodies that carry their own module imports (for example the
  observability and session replay plugin imports). ESM forbids
  `import` inside the async wrapper, so the body sits between the
  //BODY_BEGIN and //BODY_END markers and the js-client harness lifts
  its `import` lines up to //IMPORT_LIFT_TARGET at module scope. The
  scaffold supplies the `createClient` import the init body assumes.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: js-client
  entrypoint: src/app.ts
---

```javascript
import { createClient } from '@launchdarkly/js-client-sdk';
//IMPORT_LIFT_TARGET

(async () => {
  let _scaffoldSucceeded = false;
  const _origLog = console.log.bind(console);
  console.log = (...args) => {
    if (args.join(' ').includes('SDK successfully initialized')) {
      _scaffoldSucceeded = true;
    }
    _origLog(...args);
  };

//BODY_BEGIN
{{ body }}
//BODY_END

  if (_scaffoldSucceeded) {
    document.body.textContent = 'feature flag evaluates to true';
  } else {
    document.body.textContent = 'scaffold: wrappee did not print success';
  }
})();
```
