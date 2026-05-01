---
id: vue-client-sdk/sdk-docs/determine-when-the-client-is-ready-vue
sdk: vue-client-sdk
kind: reference
lang: html
description: "Vue in section \"Determine when the client is ready\""
---

```html
<script setup lang="ts">
  import { useLDReady } from 'launchdarkly-vue-client-sdk'

  const ldReady = useLDReady()
</script>

<template>
  <div v-if="ldReady">... content that uses LaunchDarkly ...</div>
  <div v-else>LaunchDarkly client initializing...</div>
</template>
```
