---
id: react-native-client-sdk/sdk-docs/features/contextconfig/context-example
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: Context example for React Native SDK v10.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only

---

```typescript
import { type LDContext } from '@launchdarkly/react-native-client-sdk';

// key and kind are the only required attributes

let context: LDContext = {
  key: 'example-user-key',
  kind: 'user',
  firstName: 'Sandy',
  lastName: 'Smith',
  email: 'sandy@example.com',
  address: {
    street: '123 Main St',
    city: 'Springfield'
  }
};
```
