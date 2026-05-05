---
id: vue-client-sdk/sdk-docs/retrieve-flag-values-for-the-context-vue
sdk: vue-client-sdk
kind: reference
lang: html
description: "Vue in section \"Retrieve flag values for the context\""
validation:
  scaffold: vue-client-sdk/scaffolds/vue-syntax-only
---

```html
<script setup lang="ts">
  import { useLDFlag } from 'launchdarkly-vue-client-sdk'
  const featureFlagKey = 'my-boolean-flag'
  const myFlagValue = useLDFlag(featureFlagKey, false /* fallback flag value */)
</script>

<template>
  Feature flag "{{ featureFlagKey }}" has value "{{ myFlagValue }}".
</template>
```
