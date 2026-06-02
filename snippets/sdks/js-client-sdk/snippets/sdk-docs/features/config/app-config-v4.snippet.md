---
id: js-client-sdk/sdk-docs/features/config/app-config-v4
sdk: js-client-sdk
kind: reference
lang: javascript
description: Application metadata configuration example for JavaScript.
---

```js
const options = {
    applicationInfo: {
	    id: 'authentication-service',
	    version: '1.0.0',
    },
};
// Create client
const client = createClient('example-client-side-id', context, options);

// Then start the client
client.start();
```
