---
id: node-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-private-attributes-3-0-syntax-two-attributes-marked-private
sdk: node-client-sdk
kind: reference
lang: javascript
description: "3.0 syntax, two attributes marked private in section \"Understanding changes to private attributes\""
---

```js
const options = { privateAttributes: ['email', 'address'] };

const client = LaunchDarkly.initialize('example-client-side-id', context, options);
```
