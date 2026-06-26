---
id: electron-client-sdk/sdk-docs/features/identify/identify-js
sdk: electron-client-sdk
kind: reference
lang: javascript
description: Identify example for Electron (JavaScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only
---

```javascript
const newUser = { key: 'example-user-key', name: 'Sandy' };

client.identify(newUser, null, (err, newFlags) => {
  console.log('value of flag for this user is: ' + newFlags['example-flag-key']);
  console.log('this should be the same: ' + client.variation('example-flag-key'));
});

// or:
client.identify(newUser).then((newFlags) => {
  // as above
});
```
