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
  Maven Central, not the android client SDK (which lives in Google's
  Maven and ships as an `aar`). Fragments that reference
  `com.launchdarkly.sdk.android.*` types won't resolve here. See
  `_sdk-docs-port-notes.md` for the structural gap; the sdk-docs
  android fragments are documented as Bucket C until a real
  `android-client-validator` Docker image (mirroring the sdk-info
  init validator) lands.
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
