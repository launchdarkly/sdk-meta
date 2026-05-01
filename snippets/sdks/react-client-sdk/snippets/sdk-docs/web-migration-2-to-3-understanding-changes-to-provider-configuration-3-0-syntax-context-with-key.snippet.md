---
id: react-client-sdk/sdk-docs/web-migration-2-to-3-understanding-changes-to-provider-configuration-3-0-syntax-context-with-key
sdk: react-client-sdk
kind: reference
lang: javascript
description: "3.0 syntax, context with key in section \"Understanding changes to provider configuration\""
---

```js
const context = {
  kind: "user",
  key: "example-user-key",
}

const LDProvider = await asyncWithLDProvider({
  clientSideID: "example-client-side-id",
  context: context,
})
```
