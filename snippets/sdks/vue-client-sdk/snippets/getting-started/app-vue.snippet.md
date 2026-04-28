---
id: vue-client-sdk/getting-started/app-vue
sdk: vue-client-sdk
kind: hello-world
lang: html
file: src/App.vue
description: src/App.vue evaluates a feature flag with useLDFlag.
inputs:
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source.
ld-application:
  slot: app-vue
# Validator pending — same as main-js.
---

In `src/App.vue`:

```html
<script setup>
import { useLDFlag, useLDReady } from 'launchdarkly-vue-client-sdk'

const ldReady = useLDReady()
const flagValue = useLDFlag('{{ featureKey }}', false)
</script>

<template>
  <div v-if="ldReady">Feature Flag {{ featureKey }} is {{ flagValue }}</div>
  <div v-else>LaunchDarkly client initializing...</div>
</template>
```
