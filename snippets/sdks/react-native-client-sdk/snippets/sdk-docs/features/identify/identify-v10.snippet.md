---
id: react-native-client-sdk/sdk-docs/features/identify/identify-v10
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: Identify example for the React Native SDK v10.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only
---

```typescript
import { useLDClient } from '@launchdarkly/react-native-client-sdk';

const client = useLDClient();
const context: LDContext = {'key': 'example-user-key', 'kind': 'user'};
    client
      .identify(context)
      .catch((e: any) => console.error(`error identifying ${context.key}: ${e}`));

```
