---
id: js-client-sdk/sdk-info/flagEval-string
sdk: js-client-sdk
kind: flag-eval
lang: javascript
file: js-client-sdk/flagEval-string.txt
description: Flag evaluation example for js-client-sdk (string flag).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```javascript
const flagKey = 'featureKey';

const flagValue = ldclient.variation(flagKey, 'fallback-value');

if (flagValue === 'example-variation-value') {

    // TODO: Put your feature here

} else {

    // TODO: Put your fallback behavior here

}
```
