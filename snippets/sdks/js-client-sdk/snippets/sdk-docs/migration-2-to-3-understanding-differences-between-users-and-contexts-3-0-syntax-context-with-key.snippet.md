---
id: js-client-sdk/sdk-docs/migration-2-to-3-understanding-differences-between-users-and-contexts-3-0-syntax-context-with-key
sdk: js-client-sdk
kind: reference
lang: javascript
description: "3.0 syntax, context with key in section \"Understanding differences between users and contexts\""
---

```js
const context = {
  kind: 'user',
  key: 'example-user-key'
};
const client = LDClient.initialize('example-client-side-id', context);
```
