---
id: js-client-sdk/experimentation/button-copy
sdk: js-client-sdk
kind: reference
lang: javascript
description: Creates a <button> element whose label reflects the assigned flag variation and tracks clicks. Drop in wherever you want to A/B test button copy.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```javascript
// Creates a <button> element whose label reflects the assigned variation.
// Call this after the client is ready and attach the returned element to the DOM
// wherever your original button lives.
//
// Prerequisites:
//   - A string flag whose key matches YOUR_FLAG_KEY. Set each variation's value to
//     the button label you want users to see (e.g. "Get started", "Start for free").
//     The flag value is used as the button label directly.
//   - A click metric whose key matches YOUR_METRIC_KEY attached to your experiment.
export function createExperimentButton({ onClick } = {}) {
  // The flag value is the button label. The default is shown when the flag is off
  // or the SDK hasn't finished initializing yet.
  // Don't cache the result — LaunchDarkly deduplicates exposure events automatically.
  const label = ldClient.variation('YOUR_FLAG_KEY', 'Get started');

  const button = document.createElement('button');
  button.textContent = label;

  button.addEventListener('click', () => {
    // Track the click so LaunchDarkly can attribute it to the right variation.
    // Mismatched contexts (evaluating as user A, tracking as user B) break conversion attribution.
    ldClient.track('YOUR_METRIC_KEY');
    onClick?.();
  });

  return button;
}
```
