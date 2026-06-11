---
id: react-native-client-sdk/experimentation/button-copy
sdk: react-native-client-sdk
kind: reference
lang: tsx
description: Standalone ExperimentButton component — reads its label from a string flag and tracks presses. Drop in wherever you want to A/B test button copy.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only
---

```tsx
import React, { useCallback, useEffect, useState } from 'react';
import { Pressable, Text } from 'react-native';
import { ReactNativeLDClient } from '@launchdarkly/react-native-client-sdk';

interface ExperimentButtonProps {
  client: ReactNativeLDClient;
  onPress?: () => void;
}

// Drop-in replacement for any pressable element you want to A/B test.
//
// Prerequisites:
//   1. Initialize an LDClient and wrap your app in <LDProvider>
//      (see the experimentation/full snippet).
//   2. Create a string flag whose key matches YOUR_FLAG_KEY. Set each variation's
//      value to the button label you want users to see (e.g. "Get started",
//      "Start for free"). The flag value is used as the button label directly.
//   3. Create a press metric in LaunchDarkly whose key matches YOUR_METRIC_KEY
//      and attach it to your experiment.
export function ExperimentButton({ client, onPress }: ExperimentButtonProps) {
  // Initialise from the current flag value and re-render whenever it changes,
  // so live targeting updates reach the button without remounting.
  const [label, setLabel] = useState<string>(
    () => client.variation('YOUR_FLAG_KEY', 'Get started') as string,
  );

  useEffect(() => {
    const updateLabel = () => {
      setLabel(client.variation('YOUR_FLAG_KEY', 'Get started') as string);
    };
    client.on('change:YOUR_FLAG_KEY', updateLabel);
    return () => { client.off('change:YOUR_FLAG_KEY', updateLabel); };
  }, [client]);

  const handlePress = useCallback(() => {
    // Track the press so LaunchDarkly can attribute it to the right variation.
    // Mismatched contexts (evaluating as user A, tracking as user B) break conversion attribution.
    client.track('YOUR_METRIC_KEY');
    onPress?.();
  }, [client, onPress]);

  return (
    <Pressable onPress={handlePress}>
      <Text>{label}</Text>
    </Pressable>
  );
}
```
