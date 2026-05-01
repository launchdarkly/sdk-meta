---
id: electron-client-sdk/sdk-docs/electron-initialize-the-client-typescript
sdk: electron-client-sdk
kind: reference
lang: typescript
description: "TypeScript in section \"Initialize the client\""
---

```ts
import * as LDElectron from 'launchdarkly-electron-client-sdk'

// You'll need this user later, but you can ignore it for now.
const user: LDElectron.LDUser = { key: 'example' }
const options: LDElectron.LDOptions = {}
const client = LDElectron.initializeInMain('example-client-side-id', user, options)
```
