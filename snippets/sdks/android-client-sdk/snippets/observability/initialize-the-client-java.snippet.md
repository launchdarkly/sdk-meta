---
id: android-client-sdk/observability/initialize-the-client-java
sdk: android-client-sdk
kind: reference
lang: java
file: app/src/main/java/com/launchdarkly/hello_android/SnippetObsJava.java
description: "Android observability plugin: Initialize the client (Java)"
validation:
  scaffold: android-client-sdk/scaffolds/java-observability
---

```java
String mobileKey = "example-mobile-key";

LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey(mobileKey)
    .plugins(Components.plugins().setPlugins(
      Collections.<Plugin>singletonList(
        new Observability(this.getApplication(), mobileKey, ObservabilityOptions.builder().build(), null)
      )
    ))
    // other options
    .build();

// You'll need this context later, but you can ignore it for now.
LDContext context = LDContext.create("example-context-key");

LDClient client = LDClient.init(this.getApplication(), ldConfig, context, 0);
```
