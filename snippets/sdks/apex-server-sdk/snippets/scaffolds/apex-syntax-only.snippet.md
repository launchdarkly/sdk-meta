---
id: apex-server-sdk/scaffolds/apex-syntax-only
sdk: apex-server-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/Snippet.java
description: |
  Placeholder syntax check for Apex doc fragments — no Apex Docker harness yet, so we use the JVM validator on a wrapped Java file. Apex-specific syntax may not parse cleanly under javac; consider this scaffold a stub until a real Apex validator lands.
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
}
```
