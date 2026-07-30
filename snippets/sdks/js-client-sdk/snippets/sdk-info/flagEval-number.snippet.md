---
id: js-client-sdk/sdk-info/flagEval-number
sdk: js-client-sdk
kind: flag-eval
lang: javascript
file: js-client-sdk/flagEval-number.txt
description: Flag evaluation example for js-client-sdk (number flag).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```javascript
const flagKey = 'featureKey';

const flagValue = ldclient.variation(flagKey, 0);

if (flagValue === 1) {

    // TODO: Put your feature here

} else {

    // TODO: Put your fallback behavior here

}
```
