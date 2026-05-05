---
id: react-client-sdk/sdk-info/flagEval
sdk: react-client-sdk
kind: flag-eval
lang: javascript
file: react-client-sdk/flagEval.txt
description: Flag evaluation example for react-client-sdk.
validation:
  scaffold: react-client-sdk/scaffolds/flag-eval-runner
---

```javascript
import { useBoolVariation } from '@launchdarkly/react-sdk';

const flagKey = 'featureKey';

// useBoolVariation evaluates a boolean feature flag and returns its value.
const flagValue = useBoolVariation(flagKey, false);

// In your component, find where your feature is instantiated
if (flagValue) {

    // TODO: Put your feature here

} else {

    // TODO: Put your fallback behavior here

}
```
