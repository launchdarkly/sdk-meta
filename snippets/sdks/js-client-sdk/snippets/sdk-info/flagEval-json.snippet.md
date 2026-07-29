---
id: js-client-sdk/sdk-info/flagEval-json
sdk: js-client-sdk
kind: flag-eval
lang: javascript
file: js-client-sdk/flagEval-json.txt
description: Flag evaluation example for js-client-sdk (JSON flag).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```javascript
const flagKey = 'featureKey';

const flagValue = ldclient.variation(flagKey, {});

// Use the flag value to configure your feature
// TODO: Put your feature here
```
