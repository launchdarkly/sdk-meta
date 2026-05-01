---
id: js-client-sdk/sdk-docs/migration-3-to-4-flag-listener-changes-javascript-sdk-v3-x
sdk: js-client-sdk
kind: reference
lang: javascript
description: "JavaScript SDK v3.x in section \"Flag listener changes\""
---

```js
  // The exact signature may have varied, but typically:
    client.on('change', (changedFlags) => {
    // changedFlags is a key value pair where the flag key is mapped
    // to a diff object.
  });
```
