---
id: android-client-sdk/sdk-docs/initialize-the-client-android-sdk-v5-x-java
sdk: android-client-sdk
kind: reference
lang: java
description: "Android SDK v5.x (Java) in section \"Initialize the client\""
# Bucket C: jvm validator pulls launchdarkly-java-server-sdk, not the android-client SDK (which lives in Google Maven as an aar). See _sdk-docs-port-notes.md.
---

```java
LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    // optional observability plugin, requires LaunchDarkly Android Client SDK v5.9+
    .plugins(Components.plugins().setPlugins(
      Collections.singletonList<Plugin>(Observability(this.getApplication()))
    ))
    // other options
    .build();

// You'll need this context later, but you can ignore it for now.
LDContext context = LDContext.create("example-context-key");

LDClient client = LDClient.init(this.getApplication(), ldConfig, context, 0);
```
