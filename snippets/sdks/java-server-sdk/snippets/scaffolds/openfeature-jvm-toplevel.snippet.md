---
id: java-server-sdk/scaffolds/openfeature-jvm-toplevel
sdk: java-server-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/tutorial/Main.java
description: |
  Resolves an OpenFeature provider doc fragment that is a set of
  top-level `import` statements. The jvm harness lifts the body's
  `import …;` lines out to file scope at the IMPORT_LIFT_MARKER, so
  javac resolves each imported name against the OpenFeature SDK and the
  LaunchDarkly provider baked into the validator image. An import that
  names a class the packages don't ship fails compilation. The body
  does not connect to LaunchDarkly, so this needs no credentials beyond
  what the harness already requires.
inputs:
  body:
    type: string
    description: The wrappee's import statements; lifted to file scope by the harness.
validation:
  runtime: jvm
  entrypoint: src/main/java/com/launchdarkly/tutorial/Main.java
---

```java
package com.launchdarkly.tutorial;

// IMPORT_LIFT_MARKER

public class Main {
    public static void main(String[] args) {
{{ body }}
        System.out.println("feature flag evaluates to true");
    }
}
```
