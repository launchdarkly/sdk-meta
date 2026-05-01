---
id: js-client-sdk/sdk-docs/migration-3-to-4-flag-listener-changes-javascript-sdk-v4-0
sdk: js-client-sdk
kind: reference
lang: javascript
description: "JavaScript SDK v4.0 in section \"Flag listener changes\""
---

```js
  // General change event - fires when any flags change
  client.on('change', (context, changedKeys) => {
    // context: The LDContext for which flags changed
    // changedKeys: Array of flag keys that changed

    // Still need to call variation() to get current values
    changedKeys.forEach(flagKey => {
      const flagValue = client.variation(flagKey, defaultValue);
    });
  });

  // Specific flag change event - fires when a specific flag changes
  client.on('change:example-flag-key', (context) => {
    // Only fires when 'my-flag' changes
    const flagValue = client.variation('example-flag-key', false);
  });
```
