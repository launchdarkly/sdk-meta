---
id: node-server-sdk/sdk-docs/features/allflags/allflags
sdk: node-server-sdk
kind: reference
lang: javascript
description: All flags example for Node.js (server-side) SDK v7.x and later.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
client.allFlagsState(context, options, (err, flagsState) => {
  // this object can be converted to JSON using toJSON()
  // or can be queried for flag values using allValues() or getFlagValue(flag-key)
});
```
