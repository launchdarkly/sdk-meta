---
id: react-client-sdk/sdk-info/flagEval-string
sdk: react-client-sdk
kind: flag-eval
lang: javascript
file: react-client-sdk/flagEval-string.txt
description: Flag evaluation example for react-client-sdk (string flag).
validation:
  scaffold: react-client-sdk/scaffolds/flag-eval-runner
  placeholders:
    feature-key: LAUNCHDARKLY_FLAG_KEY
---

```javascript
import { useStringVariation } from '@launchdarkly/react-sdk';

// useStringVariation evaluates a string feature flag and returns its value.
const flagValue = useStringVariation('feature-key', 'fallback-value');

// In your component, compare the flag value against your variation values
if (flagValue === 'example-variation-value') {

    // TODO: Put your feature here

} else {

    // TODO: Put your fallback behavior here

}
```
