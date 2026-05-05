---
id: vue-client-sdk/sdk-docs/initialize-the-client-and-context-app-vue-vue-sdk-v2-x
sdk: vue-client-sdk
kind: reference
lang: html
description: "App.vue, Vue SDK v2.x in section \"Initialize the client and context\""
validation:
  scaffold: vue-client-sdk/scaffolds/vue-syntax-only
---

```html
<script setup>
import { ldInit } from 'launchdarkly-vue-client-sdk'

const [ldReady, ldClient] = ldInit({ context: { kind: 'user', key: 'example-context-key', name: 'Sandy' } })

ldClient.waitForInitialization(5).catch((error) => {

  // Client encountered an error or timeout, but is ready to use.
  // variation() calls will return their fallback values.

  if (error?.name.toLowerCase().includes('timeout')) {
    console.log(`===== timeout error: ${error}`)
  } else {
    console.log(`===== other init error: ${error}`)
  }
})

</script>

<template>
  <div v-if="ldReady">LaunchDarkly initialized.</div>
  <div v-else>LaunchDarkly initializing.</div>
</template>
```
