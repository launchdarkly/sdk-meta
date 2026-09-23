---
id: cloudflare-server-sdk/sdk-docs/openfeature/evaluate-a-context
sdk: cloudflare-server-sdk
kind: reference
lang: javascript
file: cloudflare-server-sdk/sdk-docs/openfeature/evaluate-a-context.mjs
description: "Cloudflare OpenFeature provider in section \"Evaluate a context\""
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only
---

```js
const flagValue = await client.getBooleanValue("example-flag-key", false, context);
```
