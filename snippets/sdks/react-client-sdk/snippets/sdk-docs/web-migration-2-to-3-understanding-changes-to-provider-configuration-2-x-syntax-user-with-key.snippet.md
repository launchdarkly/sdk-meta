---
id: react-client-sdk/sdk-docs/web-migration-2-to-3-understanding-changes-to-provider-configuration-2-x-syntax-user-with-key
sdk: react-client-sdk
kind: reference
lang: javascript
description: "2.x syntax, user with key in section \"Understanding changes to provider configuration\""
---

```js
const user = {
  key: "example-user-key",
}

const LDProvider = await asyncWithLDProvider({
  clientSideID: "example-client-side-id",
  user: user,
})
```
