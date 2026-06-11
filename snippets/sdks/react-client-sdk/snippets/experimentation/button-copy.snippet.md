---
id: react-client-sdk/experimentation/button-copy
sdk: react-client-sdk
kind: reference
lang: tsx
description: Standalone ExperimentButton component — reads its label from a string flag and tracks clicks. Drop in wherever you want to A/B test button copy.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```tsx
import { useCallback } from "react";
import { useLDClient, useStringVariation } from "@launchdarkly/react-sdk";

interface ExperimentButtonProps {
  onClick?: () => void;
  className?: string;
}

// Drop-in replacement for any <button> you want to A/B test.
//
// Prerequisites:
//   1. Wrap your app in <LDProvider> (see the experimentation/full snippet).
//   2. Create a string flag whose key matches YOUR_FLAG_KEY. Set each variation's
//      value to the button label you want users to see (e.g. "Get started",
//      "Start for free"). The flag value is used as the button label directly.
//   3. Create a click metric in LaunchDarkly whose key matches YOUR_METRIC_KEY
//      and attach it to your experiment.
export function ExperimentButton({ onClick, className }: ExperimentButtonProps) {
  // useStringVariation looks the flag up by its key, subscribes to it, and
  // re-renders the component whenever its value changes, so live targeting
  // rule updates reach the button without a page reload. The default is shown
  // when the flag is off or the SDK hasn't finished initializing yet.
  const label = useStringVariation("YOUR_FLAG_KEY", "Get started");
  const ldClient = useLDClient();

  const handleClick = useCallback(() => {
    // Track the click so LaunchDarkly can attribute conversions to the right variation.
    // Use the same user context that was active when the flag was evaluated —
    // mismatched contexts break conversion attribution.
    ldClient.track("YOUR_METRIC_KEY");
    onClick?.();
  }, [ldClient, onClick]);

  return (
    <button className={className} onClick={handleClick}>
      {label}
    </button>
  );
}
```
