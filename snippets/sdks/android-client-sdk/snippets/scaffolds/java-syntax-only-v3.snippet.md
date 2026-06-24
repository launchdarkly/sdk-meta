---
id: android-client-sdk/scaffolds/java-syntax-only-v3
sdk: android-client-sdk
kind: scaffold
lang: java
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.java
description: |
  Parse-and-type-check validator for Android Java doc fragments that
  target the v3.x API surface (`client.alias(newUser, previousUser)` —
  the alias method was removed at 4.0, and LDUser-based calls cannot
  compile against the v5 aar the `android-client` validator container
  pre-bakes).

  Unlike `java-syntax-only-v4` (which routes through the `jvm`
  validator), this scaffold stays on the `android-client` container's
  parse path: the jvm harness requires a server-side SDK key that the
  android CI row does not provision, while the android parse path
  needs no key at all. The stub surface is fully self-contained —
  nested `LDClient` / `LDUser` classes declare just the members the
  v3-era fragments reference, so the compile never touches the real
  aar's API surface and gradle's `compileDebugJavaWithJavac` provides
  the syntax + type check.

  The nested classes shadow any same-named types from the package
  wildcard imports the sibling scaffolds use; this file deliberately
  imports nothing from the SDK so the stubs are authoritative.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: android-client
  entrypoint: app/src/main/java/com/launchdarkly/hello_android/Snippet.java
  env:
    SNIPPET_CHECK: parse
---

```java
package com.launchdarkly.hello_android;

public class Snippet {
    // Stub of the v3-era android user type. Only referenced as an
    // opaque argument by the doc fragments, so no members are needed.
    static class LDUser {
    }

    // Stub of the v3-era android LDClient surface. Only the members
    // the doc fragments call are declared.
    static class LDClient {
        void alias(LDUser newUser, LDUser previousUser) {
        }
    }

    // Ambient names the doc fragments assume an Activity host
    // provides.
    @SuppressWarnings("unused")
    private static final LDClient client = null;
    @SuppressWarnings("unused")
    private static final LDUser newUser = null;
    @SuppressWarnings("unused")
    private static final LDUser previousUser = null;

    @SuppressWarnings({"unused", "ConstantConditions"})
    private void wrappee() {
        try {
{{ body }}
        } catch (Throwable ignored) { }
    }
}
```
