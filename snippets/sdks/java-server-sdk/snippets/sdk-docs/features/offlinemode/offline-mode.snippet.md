---
id: java-server-sdk/sdk-docs/features/offlinemode/offline-mode
sdk: java-server-sdk
kind: reference
lang: java
description: Offline mode example for Java SDK v6.0.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
LDConfig config = new LDConfig.Builder()
  .offline(true)
  .build();
LDClient client = new LDClient("YOUR_SDK_KEY", config);
client.boolVariation("example-flag-key", context, false);

// This call to client.boolVariation always
// returns the default value (false)

```
