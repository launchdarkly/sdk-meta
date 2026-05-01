---
id: android-client-sdk/sdk-docs/migration-2-to-3-gson-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: "Kotlin in section \"Gson\""
---

```kotlin
import com.google.gson.*
import com.launchdarkly.sdk.json.LDGson

val gson: Gson = GsonBuilder()
  .registerTypeAdapterFactory(LDGson.typeAdapters())
  // any other GsonBuilder options go here
  .create()
```
