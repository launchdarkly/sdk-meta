---
id: java-server-sdk/scaffolds/openfeature-jvm-runner
sdk: java-server-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/tutorial/Main.java
description: |
  Runs an OpenFeature provider doc fragment that assumes a registered
  provider and a bound `provider`, `client`, and `context` already
  exist — the "evaluate a context" and "access the LaunchDarkly client"
  fragments. The scaffold registers a real LaunchDarkly provider, binds
  those names, runs the fragment, then evaluates a flag and prints the
  success line. Java has no local-variable shadowing, so fragments that
  declare their own `context` use the `openfeature-jvm-context-runner`
  variant instead. Requires LaunchDarkly credentials.
inputs:
  body:
    type: string
    description: The wrappee fragment, run with `provider`, `client`, and `context` in scope.
validation:
  runtime: jvm
  entrypoint: src/main/java/com/launchdarkly/tutorial/Main.java
---

```java
package com.launchdarkly.tutorial;

import dev.openfeature.sdk.OpenFeatureAPI;
import dev.openfeature.sdk.Client;
import dev.openfeature.sdk.EvaluationContext;
import dev.openfeature.sdk.ImmutableContext;
import com.launchdarkly.sdk.server.interfaces.LDClientInterface;
import com.launchdarkly.openfeature.serverprovider.Provider;

public class Main {
    public static void main(String[] args) throws Exception {
        Provider provider = new Provider(System.getenv("LAUNCHDARKLY_SDK_KEY"));
        OpenFeatureAPI.getInstance().setProvider(provider);
        Client client = OpenFeatureAPI.getInstance().getClient();
        EvaluationContext context = new ImmutableContext("example-user-key");

{{ body }}

        client.getBooleanValue(System.getenv("LAUNCHDARKLY_FLAG_KEY"), false, context);
        System.out.println("feature flag evaluates to true");
    }
}
```
