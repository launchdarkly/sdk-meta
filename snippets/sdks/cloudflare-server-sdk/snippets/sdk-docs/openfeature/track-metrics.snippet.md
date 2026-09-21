---
id: cloudflare-server-sdk/sdk-docs/openfeature/track-metrics
sdk: cloudflare-server-sdk
kind: reference
lang: javascript
file: cloudflare-server-sdk/sdk-docs/openfeature/track-metrics.mjs
description: "Cloudflare OpenFeature provider in section \"Track metrics\""
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only
---

```js
client.track('example-event-key', context, { "optionalCustomData": "optionalCustomValue"});

ctx.waitUntil(provider.getClient().flush());
```
