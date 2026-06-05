---
id: java-server-sdk/scaffolds/java-syntax-only
sdk: java-server-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/Snippet.java
description: |
  Parse-only validator for Java server SDK doc fragments.

  The wrappee body is dropped inside `wrappee()` (where unresolved
  `client.boolVariation(...)` calls don't fail compilation thanks to
  the stub `client` field). Java forbids `import` statements inside a
  method body, so any top-level `import …;` lines in the wrappee are
  lifted out at validate-time by the harness's pre-stage rewrite —
  `Snippet.java`'s body section reaches the compiler with imports
  already at file scope. The IMPORT_LIFT_MARKER comment is the cue
  the rewrite uses to splice extracted imports above the class
  declaration.
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

import com.launchdarkly.sdk.*;
import com.launchdarkly.sdk.server.*;
import com.launchdarkly.sdk.server.migrations.*;
// Common JDK types config/timeout fragments reference without their own
// import line (the docs assume it); provide it so they resolve.
import java.time.Duration;
// IMPORT_LIFT_MARKER

public class Snippet {
    // Stub instance the wrappee body refers to. Never used at runtime;
    // present so the body's `client.boolVariation(…)` calls resolve a
    // symbol during compilation.
    @SuppressWarnings("unused")
    private static final LDClient client = null;

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
