---
id: android-client-sdk/scaffolds/android-syntax-only
sdk: android-client-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/Snippet.java
description: |
  Parse-only validator for Android client SDK doc fragments. Uses the JVM validator since Android is a JVM language.
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
