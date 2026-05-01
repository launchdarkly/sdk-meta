---
id: node-server-sdk/sdk-docs/migration-6-to-7-understanding-changes-to-alias-events-7-0-syntax-javascript
sdk: node-server-sdk
kind: reference
lang: javascript
description: "7.0 syntax, JavaScript in section \"Understanding changes to alias events\""
---

```js
const context = {
  kind: 'multi',
  user: { key: "example-user-key" },
  device: { key: "example-device-key" }
}

client.identify(context)
```
