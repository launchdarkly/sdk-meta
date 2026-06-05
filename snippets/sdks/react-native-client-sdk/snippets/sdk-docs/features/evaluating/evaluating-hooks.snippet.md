---
id: react-native-client-sdk/sdk-docs/features/evaluating/evaluating-hooks
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: Flag evaluation example using variation hooks for React Native.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only

---

```typescript
import { useBoolVariation, useNumberVariation, useStringVariation, useJsonVariation } from '@launchdarkly/react-native-client-sdk'

const boolResult = useBoolVariation('example-bool-flag-key', false);
const numResult = useNumberVariation('example-numeric-flag-key', 2);
const stringResult = useStringVariation('example-string-flag-key', '');
const jsonResult = useJsonVariation('example-json-flag-key', {});
```
