---
id: js-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v3
sdk: js-client-sdk
kind: reference
lang: javascript
description: Flag evaluation reason example for JavaScript SDK v3.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```javascript
const options = { evaluationReasons: true };
const client = LDClient.initialize('example-client-side-id', context, options);

try {
    client.waitForInitialization(5);

    // proceed with successfully initialized client:

    const detail = client.variationDetail('example-flag-key', false);

    const value = detail.value;
    const index = detail.variationIndex;
    const reason = detail.reason;

} catch(err) {
    // Client failed to initialize or timed out
    // variation() calls return fallback values until initialization completes
}
```
