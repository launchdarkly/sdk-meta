---
id: java-server-sdk/sdk-docs/migration-4-to-5-gson-java
sdk: java-server-sdk
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
