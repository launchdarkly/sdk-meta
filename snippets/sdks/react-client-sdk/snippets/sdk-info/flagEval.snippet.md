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
import { useFlags } from 'launchdarkly-react-client-sdk';

// We added your flag key. The React SDK uses camelCase for flag keys automatically
// useFlags is a custom hook which returns all feature flags
const { featureKey } = useFlags();

// In your component, find where your feature is instantiated
if (featureKey) {

    // TODO: Put your feature here

} else {

    // TODO: Put your fallback behavior here

}
```
