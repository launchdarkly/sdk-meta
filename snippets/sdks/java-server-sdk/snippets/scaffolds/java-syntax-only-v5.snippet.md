---
id: java-server-sdk/scaffolds/java-syntax-only-v5
sdk: java-server-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/Snippet.java
description: |
  Parse-and-type-check validator for Java server SDK doc fragments
  that target the v5.x API surface (`client.alias(user, previousUser)`
  — the alias method was removed at 6.0, so v5-era fragments cannot
  compile against the v7 SDK jar the `jvm` validator's pom pins).

  Routes through the same `jvm` validator as `java-syntax-only`,
  following the same stub approach as the android
  `java-syntax-only-v4` scaffold: nested stub classes declare just
  the v5 surface the fragments reference, and this file deliberately
  imports nothing from `com.launchdarkly.sdk.server` so the stubs
  are authoritative regardless of the pinned SDK version on the
  classpath.
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
    // Stub of the v5-era user type. Only referenced as an opaque
    // argument by the doc fragments, so no members are needed.
    static class LDUser {
    }

    // Stub of the v5-era LDClient surface. Only the members the doc
    // fragments call are declared.
    static class LDClient {
        void alias(LDUser user, LDUser previousUser) {
        }
    }

    // Ambient names the doc fragments assume earlier snippets
    // created.
    @SuppressWarnings("unused")
    private static final LDClient client = null;
    @SuppressWarnings("unused")
    private static final LDUser user = null;
    @SuppressWarnings("unused")
    private static final LDUser previousUser = null;

    public static void main(String[] args) {
        System.out.println("feature flag evaluates to true");
    }

    @SuppressWarnings({"unused", "ConstantConditions"})
    private void wrappee() {
        try {
{{ body }}
        } catch (Throwable ignored) { }
    }
}
```
