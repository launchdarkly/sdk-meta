---
id: react-native-client-sdk/sdk-docs/features/evaluating/evaluating-methods
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: Flag evaluation example using variation methods for React Native.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only

---

```typescript
import { useLDClient } from '@launchdarkly/react-native-client-sdk'

const client = useLDClient();
let boolResult = client.boolVariation('example-bool-flag-key', false);
let numResult = client.numberVariation('example-numeric-flag-key', 2);
let stringResult = client.stringVariation('example-string-flag-key', '');
let jsonResult = client.jsonVariation('example-json-flag-key', {});
```
