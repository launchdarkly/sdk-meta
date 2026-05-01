---
id: js-client-sdk/sdk-info/flagEval
sdk: js-client-sdk
kind: flag-eval
lang: javascript
file: js-client-sdk/flagEval.txt
description: Flag evaluation example for js-client-sdk.
---

```javascript
const flagKey = 'featureKey';

const flagValue = ldclient.variation(flagKey, false);

if (flagValue) {

    // TODO: Put your feature here

} else {

    // TODO: Put your fallback behavior here

}
```
