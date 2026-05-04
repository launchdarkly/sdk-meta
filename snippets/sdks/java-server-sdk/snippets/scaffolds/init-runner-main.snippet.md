---
id: java-server-sdk/scaffolds/init-runner-main
sdk: java-server-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/tutorial/runner/Runner.java
description: |
  Runner half of the Java init scaffold pair. Lives under a different
  package (`com.launchdarkly.tutorial.runner`) so both this and the
  wrappee's `Main` class can be public top-level classes. Listed as a
  companion of `init-runner` and used as the maven mainClass via the
  parent scaffold's `validation.entrypoint`.
---

```java
package com.launchdarkly.tutorial.runner;

import java.io.ByteArrayOutputStream;
import java.io.PrintStream;

public class Runner {
    public static void main(String[] args) throws Exception {
        ByteArrayOutputStream buf = new ByteArrayOutputStream();
        PrintStream tee = new PrintStream(new java.io.OutputStream() {
            @Override public void write(int b) {
                System.err.write(b);
                buf.write(b);
            }
        }, true);
        PrintStream origOut = System.out;
        System.setOut(tee);
        try {
            com.launchdarkly.tutorial.Main.main(args);
        } finally {
            System.setOut(origOut);
        }
        String captured = buf.toString();
        if (!captured.contains("SDK successfully initialized")) {
            System.err.println("scaffold: wrappee did not print 'SDK successfully initialized'");
            System.exit(1);
        }
        System.out.println("feature flag evaluates to true");
    }
}
```
