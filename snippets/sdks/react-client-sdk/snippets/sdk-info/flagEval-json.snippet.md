---
id: react-client-sdk/sdk-info/flagEval-json
sdk: react-client-sdk
kind: flag-eval
lang: javascript
file: react-client-sdk/flagEval-json.txt
description: Flag evaluation example for react-client-sdk (JSON flag).
validation:
  scaffold: react-client-sdk/scaffolds/flag-eval-runner
  placeholders:
    feature-key: LAUNCHDARKLY_FLAG_KEY
---

```javascript
import { useJsonVariation } from '@launchdarkly/react-sdk';

// useJsonVariation evaluates a JSON feature flag and returns its value.
const flagValue = useJsonVariation('feature-key', {});

// In your component, use the flag value to configure your feature
// TODO: Put your feature here
```
