---
id: android-client-sdk/sdk-docs/initialize-the-client-android-sdk-v5-x-java
sdk: android-client-sdk
kind: reference
lang: java
description: "Android SDK v5.x (Java) in section \"Initialize the client\""
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java
LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    // optional observability plugin, requires LaunchDarkly Android Client SDK v5.9+
    .plugins(Components.plugins().setPlugins(
      Collections.<Plugin>singletonList(
        new Observability(this.getApplication(), "example-mobile-key", ObservabilityOptions.builder().build())
      )
    ))
    // other options
    .build();

// You'll need this context later, but you can ignore it for now.
LDContext context = LDContext.create("example-context-key");

LDClient client = LDClient.init(this.getApplication(), ldConfig, context, 0);
```
