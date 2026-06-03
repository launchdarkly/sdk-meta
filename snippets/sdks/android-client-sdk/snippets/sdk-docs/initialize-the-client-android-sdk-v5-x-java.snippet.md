---
id: android-client-sdk/sdk-docs/initialize-the-client-android-sdk-v5-x-java
sdk: android-client-sdk
kind: reference
lang: java
description: "Android SDK v5.x (Java) in section \"Initialize the client\""
# TODO(snippet-bug): body has Kotlin syntax in a `java` code block:
# (1) `Observability(this.getApplication())` — Java requires `new`
# before a constructor call; (2)
# `Collections.singletonList<Plugin>(...)` — Java type arguments
# on a method call must precede the method name
# (`Collections.<Plugin>singletonList(...)`), not follow it. Fix in
# the snippet-bugs PR.
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
