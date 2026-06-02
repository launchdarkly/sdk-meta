---
id: js-client-sdk/sdk-docs/features/config/app-config-v3
sdk: js-client-sdk
kind: reference
lang: javascript
description: Application metadata configuration example for JavaScript.
---

```js
const options = {
    application: {
	    id: 'authentication-service',
	    version: '1.0.0',
    },
};
const client = LDClient.initialize('example-client-side-id', context, options);
```
