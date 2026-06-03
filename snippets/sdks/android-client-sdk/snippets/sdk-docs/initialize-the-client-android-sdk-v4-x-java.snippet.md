---
id: android-client-sdk/sdk-docs/initialize-the-client-android-sdk-v4-x-java
sdk: android-client-sdk
kind: reference
lang: java
description: "Android SDK v4.x (Java) in section \"Initialize the client\""
# TODO(validator): body uses the v4.x `new LDConfig.Builder()` 0-arg
# constructor, but the android-client validator's pre-baked
# app/build.gradle pulls in `launchdarkly-android-client-sdk:5.x`
# where Builder requires `AutoEnvAttributes`. Needs a v4-pinned
# android-client-v4 validator (clone of android-client with
# LD_ANDROID_SDK_VERSION pinned to a 4.x release). Land in a
# follow-up validator-update PR.
---

```java
LDConfig ldConfig = new LDConfig.Builder()
    .mobileKey("example-mobile-key")
    .build();

// You'll need this context later, but you can ignore it for now.
LDContext context = LDContext.create("example-context-key");

LDClient client = LDClient.init(this.getApplication(), ldConfig, context, 0);
```
