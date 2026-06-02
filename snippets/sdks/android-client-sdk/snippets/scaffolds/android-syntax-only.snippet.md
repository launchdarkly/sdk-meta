---
id: android-client-sdk/scaffolds/android-syntax-only
sdk: android-client-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/Snippet.java
description: |
  Parse-only validator for Android client SDK doc fragments. Uses
  the JVM validator since Android is a JVM language.

  NOTE: the JVM validator fetches `launchdarkly-java-server-sdk` from
  Maven Central, not the android client SDK. For fragments that
  reference `com.launchdarkly.sdk.android.*` types, use
  `android-client-sdk/scaffolds/kotlin-syntax-only` instead — it
  routes through the `android-client` Docker validator
  (`./gradlew compileDebugKotlin`) which has the real
  `launchdarkly-android-client-sdk` aar + AndroidX on the classpath.
  This `android-syntax-only` (java/JVM-routed) scaffold remains for
  fragments that don't touch android-client-specific types.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: jvm
  entrypoint: src/main/java/com/launchdarkly/Snippet.java
---

```java
package com.launchdarkly;

public class Snippet {
    public static void main(String[] args) {
        System.out.println("feature flag evaluates to true");
    }

    @SuppressWarnings("unused")
    private void wrappee() {
        try {
{{ body }}
        } catch (Throwable ignored) { }
    }
}
```
