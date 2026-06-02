---
id: android-client-sdk/sdk-docs/initialize-the-client-android-sdk-v4-x-java
sdk: android-client-sdk
kind: reference
lang: java
description: "Android SDK v4.x (Java) in section \"Initialize the client\""
# TODO(validate): jvm validator pulls launchdarkly-java-server-sdk, not the android-client SDK (which lives in Google Maven as an aar). See _sdk-docs-port-notes.md.
---

```java
LDConfig ldConfig = new LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .build();

// You'll need this context later, but you can ignore it for now.
LDContext context = LDContext.create("example-context-key");

LDClient client = LDClient.init(this.getApplication(), ldConfig, context, 0);
```
