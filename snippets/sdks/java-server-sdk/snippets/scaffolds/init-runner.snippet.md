---
id: java-server-sdk/scaffolds/init-runner
sdk: java-server-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/tutorial/Main.java
description: |
  Runs an `init.txt`-style snippet end-to-end against a real LaunchDarkly
  env. The wrappee body declares its own `public class Main` with a
  `main` method, so this scaffold drops the body verbatim under the
  `com.launchdarkly.tutorial` package and lets the jvm harness compile
  it. A separate runner companion (different package, different class)
  is the maven entrypoint: it invokes `Main.main(args)` via reflection,
  asserts the wrappee printed `SDK successfully initialized`, and emits
  the EXAM-HELLO success line the validator harness greps for.

  Layout:
    - src/main/java/com/launchdarkly/tutorial/Main.java (this scaffold) —
      the snippet body verbatim, with the literal SDK-key placeholder
      already substituted via `validation.placeholders`.
    - src/main/java/com/launchdarkly/tutorial/runner/Runner.java
      (companion `init-runner-main`) — top-level runner that the harness
      treats as the maven mainClass.

  Two top-level public `Main` classes can't coexist in the same package
  in Java; the runner sits under a sub-package to keep both as public
  classes in their own files.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: jvm
  entrypoint: src/main/java/com/launchdarkly/tutorial/runner/Runner.java
  companions:
    - java-server-sdk/scaffolds/init-runner-main
---

```java
package com.launchdarkly.tutorial;

{{ body }}
```
