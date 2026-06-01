---
id: react-client-sdk/experimentation/full
sdk: react-client-sdk
kind: reference
lang: tsx
description: Full experimentation onboarding for react-client-sdk — initialize, identify on login/eligibility, evaluate, and track conversions.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```tsx
import { StrictMode, useEffect, useState, useCallback } from "react";
import { createRoot } from "react-dom/client";
import {
  createLDReactProvider,
  useLDClient,
  useFlags,
} from "@launchdarkly/react-sdk";

// Initialize once — extra clients can cause inconsistent experiment results.
// If you already know the user's key at startup, initialize with it directly and
// skip the identify() call below — that removes a window for mistakes. Use an
// anonymous context only when you don't yet know who the user is.
const LDProvider = createLDReactProvider("YOUR_CLIENT_SIDE_ID", {
  kind: "user",
  key: "EXAMPLE_CONTEXT_KEY", // use a consistent key so the same user gets the same experience
  anonymous: true,
});

function App() {
  const ldClient = useLDClient();
  const flags = useFlags();
  const [isIdentified, setIsIdentified] = useState(false);

  // Only needed when the user becomes known after startup (login/eligibility).
  // Reuse the anonymous key where possible so experiment assignment stays stable,
  // and wait for identify to finish before evaluating experiment flags.
  useEffect(() => {
    if (!ldClient) return;

    (async () => {
      await ldClient.identify({
        kind: "user",
        key: "EXAMPLE_CONTEXT_KEY", // use the logged-in user's ID so experiment assignment stays consistent
        anonymous: false,
        // any attributes that affect targeting or eligibility
      });
      setIsIdentified(true);
    })();
  }, [ldClient]);

  // Evaluate the experiment flag where the user encounters the experience.
  const experimentValue = isIdentified ? flags["YOUR_FLAG_KEY"] : "control";

  // Call this when the user completes a metric action.
  // Use the same user key you used when evaluating the flag — mismatched keys break conversion tracking.
  // The data argument is optional and accepts any shape your metric needs.
  const trackMetric = useCallback(
    (metricKey, data) => {
      ldClient?.track(metricKey, data);
    },
    [ldClient],
  );

  // The SDK batches and flushes events automatically. Don't add manual flush()
  // calls — they're unnecessary and actively harmful to performance in
  // long-running apps.
  // Don't skip or cache flag evaluations to reduce exposure counts — LaunchDarkly deduplicates them automatically.

  return (
    <button onClick={() => trackMetric("YOUR_METRIC_KEY" /* optional data */)}>
      {experimentValue}
    </button>
  );
}

createRoot(document.getElementById("root")).render(
  <StrictMode>
    <LDProvider>
      <App />
    </LDProvider>
  </StrictMode>,
);
```
