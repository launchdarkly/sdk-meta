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
import React, { useCallback } from 'react';
import { Pressable, Text } from 'react-native';
import { useLDClient, useStringVariation } from '@launchdarkly/react-native-client-sdk';

interface ExperimentButtonProps {
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
export function ExperimentButton({ onPress }: ExperimentButtonProps) {
  // useStringVariation subscribes to the flag, so the label updates as soon as
  // initialization completes or targeting changes — no manual change listener
  // or parent re-render needed. The default is shown when the flag is off or
  // the SDK hasn't finished initializing yet.
  // Don't cache the result — LaunchDarkly deduplicates exposure events automatically.
  const label = useStringVariation('YOUR_FLAG_KEY', 'Get started');
  const client = useLDClient();

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
