---
id: react-client-sdk/sdk-info/flagEval-number
sdk: react-client-sdk
kind: flag-eval
lang: javascript
file: react-client-sdk/flagEval-number.txt
description: Flag evaluation example for react-client-sdk (number flag).
validation:
  scaffold: react-client-sdk/scaffolds/flag-eval-runner
  placeholders:
    feature-key: LAUNCHDARKLY_FLAG_KEY
---

```javascript
import { useNumberVariation } from '@launchdarkly/react-sdk';

// useNumberVariation evaluates a number feature flag and returns its value.
const flagValue = useNumberVariation('feature-key', 0);

// In your component, compare the flag value against your variation values
if (flagValue === 1) {

    // TODO: Put your feature here

} else {

    // TODO: Put your fallback behavior here

}
```
