---
id: js-client-sdk/sdk-docs/features/aliasing/alias-v2
sdk: js-client-sdk
kind: reference
lang: javascript
description: Alias event example for JavaScript SDK v2.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
const previousUser = client.getUser();
client.alias(newUser, previousUser);
```
