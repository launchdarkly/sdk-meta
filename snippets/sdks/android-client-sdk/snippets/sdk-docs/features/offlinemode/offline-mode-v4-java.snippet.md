---
id: android-client-sdk/sdk-docs/features/offlinemode/offline-mode-v4-java
sdk: android-client-sdk
kind: reference
lang: java
description: Offline mode example for Android SDK v4.x (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only-v4-android

---

```java
LDConfig ldConfig = new LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .offline(true)
    .build();

LDClient client = LDClient.init(this.getApplication(), ldConfig, context, 0);

// Or to switch an already-instantiated client to offline mode:
client.setOffline();
```
