---
id: js-client-sdk/sdk-docs/features/identify/identify-multi
sdk: js-client-sdk
kind: reference
lang: javascript
description: Identify example for the JavaScript SDK v3+ (multi-context).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```javascript
const newMultiContext = {
  kind: 'multi',
  user: {
    key: 'example-user-key',
    name: 'Anna',
    role: 'doctor'
  },
  device: {
    key: 'example-device-key',
    platform: 'Android'
  }
}

client.identify(newMultiContext, hash, function() {
  console.log("New context's flags available");
});
```
