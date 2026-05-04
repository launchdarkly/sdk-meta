---
id: react-native-client-sdk/scaffolds/init-runner-welcome
sdk: react-native-client-sdk
kind: scaffold
lang: tsx
file: src/welcome.tsx
description: |
  Stub welcome companion for the react-native-client init scaffold.
  The harness copies `src/welcome.tsx` from the staged snippet into the
  pre-baked project; the init scaffold doesn't actually use a Welcome
  component but the file has to exist for the cp to succeed.
---

```tsx
import React from 'react';
import { Text, View } from 'react-native';

export default function Welcome() {
  return <View><Text>welcome stub (unused by init scaffold)</Text></View>;
}
```
