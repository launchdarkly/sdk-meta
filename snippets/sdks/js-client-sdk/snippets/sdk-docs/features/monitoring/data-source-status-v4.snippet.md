---
id: js-client-sdk/sdk-docs/features/monitoring/data-source-status-v4
sdk: js-client-sdk
kind: reference
lang: javascript
description: Data source status event listener for JavaScript SDK v4.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```javascript
client.on('dataSourceStatus', (context, error) => {
  // handle error

  // error.state may be one of:
  // Initializing, Valid, Interrupted, SetOffline, Closed
});
```
