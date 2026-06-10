---
id: java-server-sdk/scaffolds/java-syntax-only-members
sdk: java-server-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/Snippet.java
description: |
  Class-member-scope sibling of `java-syntax-only`. That scaffold
  splices the body inside `wrappee()`, which breaks for fragments
  that are themselves method declarations (Java has no local
  methods). This variant splices the body at class scope instead.

  Same `jvm` validator, so the body compiles against the real
  `launchdarkly-java-server-sdk` from Maven Central.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced at class scope.
validation:
  runtime: jvm
  entrypoint: src/main/java/com/launchdarkly/Snippet.java
---

```java
package com.launchdarkly;

import com.launchdarkly.sdk.*;
import com.launchdarkly.sdk.server.*;

@SuppressWarnings("unused")
public class Snippet {
    public static void main(String[] args) {
        System.out.println("feature flag evaluates to true");
    }

{{ body }}
}
```
