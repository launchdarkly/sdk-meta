---
id: react-native-client-sdk/getting-started/welcome-tsx
sdk: react-native-client-sdk
kind: hello-world
lang: tsx
file: src/welcome.tsx
description: Welcome component that evaluates the flag and renders the value.
inputs:
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source.
ld-application:
  slot: welcome-tsx
---

Create a new file `src/welcome.tsx`:

```tsx
import {useEffect} from 'react';
import {Text, View} from 'react-native';

import {useBoolVariation, useLDClient} from '@launchdarkly/react-native-client-sdk';

export default function Welcome() {
    const flagValue = useBoolVariation('{{ featureKey }}', false);
    const ldc = useLDClient();

    useEffect(() => {
        ldc
            .identify({kind: 'user', key: 'example-user-key', name: 'Sandy'})
            .catch((e: any) => console.error('error: ' + e));
    }, []);

    return (
        <View style={{height: '100%', backgroundColor: flagValue ? '#00844B' : '#373841'}}>
            <Text style={{marginTop: 160, color: 'white', textAlign: 'center'}}>The {{ featureKey }} feature flag
                evaluates to <Text
                    style={{fontWeight: 'bold'}}>{flagValue ? 'true' : 'false'}</Text></Text>
        </View>
    );
}
```
