---
id: android-client-sdk/sdk-docs/migration-2-to-3-gson-java
sdk: android-client-sdk
kind: reference
lang: java
description: "Java in section \"Gson\""
---

```java
import com.google.gson.*;
import com.launchdarkly.sdk.json.LDGson;

Gson gson = new GsonBuilder()
  .registerTypeAdapterFactory(LDGson.typeAdapters())
  // any other GsonBuilder options go here
  .create();
```
