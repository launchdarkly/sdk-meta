---
id: vue-client-sdk/sdk-docs/access-the-underlying-javascript-sdk-javascript
sdk: vue-client-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"Access the underlying JavaScript SDK\""
validation:
  scaffold: vue-client-sdk/scaffolds/vue-syntax-only
---

```js
<script setup>
import { useLDClient } from 'launchdarkly-vue-client-sdk'

const [ldReady] = ldInit({ clientSideID, context })
const ldClient = useLDClient()
</script>

<template>
  <div v-if="ldReady">
    <p>All flags: {{ JSON.stringify(ldClient.allFlags()) }}</p>
  </div>
  <div v-else>LaunchDarkly client initializing...</div>
</template>
```
