---
id: java-server-sdk/scaffolds/java-syntax-only
sdk: java-server-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/Snippet.java
description: |
  Parse-only validator for Java server SDK doc fragments.
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
