---
id: android-client-sdk/observability/configure-plugin-options-java
sdk: android-client-sdk
kind: reference
lang: java
file: app/src/main/java/com/launchdarkly/hello_android/SnippetObsJava.java
description: "Android observability plugin: Configure the plugin options (Java)"
validation:
  scaffold: android-client-sdk/scaffolds/java-observability
---

```java
String mobileKey = "example-mobile-key";

LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey(mobileKey)
    .plugins(
      Components.plugins().setPlugins(
        Collections.<Plugin>singletonList(
          new Observability(
            this.getApplication(),
            mobileKey,
            ObservabilityOptions.builder()
              .resourceAttributes(Attributes.of(
                AttributeKey.stringKey("serviceName"), "example-service"
              ))
              .build(),
            null
          )
        )
      )
    )
    .build();
```
