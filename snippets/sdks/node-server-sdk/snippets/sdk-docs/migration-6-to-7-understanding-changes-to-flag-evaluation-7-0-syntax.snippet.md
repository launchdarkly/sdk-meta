---
id: node-server-sdk/sdk-docs/migration-6-to-7-understanding-changes-to-flag-evaluation-7-0-syntax
sdk: node-server-sdk
kind: reference
lang: javascript
description: "7.0 syntax in section \"Understanding changes to flag evaluation\""
---

```js
client.variation('example-flag-key', context, false,
  (err, value) => {
    // check value and proceed accordingly
  });
```
