---
id: java-server-sdk/scaffolds/openfeature-jvm-context-runner
sdk: java-server-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/tutorial/Main.java
description: |
  Runs an OpenFeature "construct a context" fragment, which declares its
  own `context` local. The scaffold registers a real LaunchDarkly
  provider and binds `client`, but leaves `context` for the fragment to
  declare; afterward it evaluates a flag with the fragment's `context`
  and prints the success line. Separate from `openfeature-jvm-runner`
  because Java forbids re-declaring a local that the scaffold already
  bound. Requires LaunchDarkly credentials.
inputs:
  body:
    type: string
    description: The wrappee fragment; declares `context` and runs with `client` in scope.
validation:
  runtime: jvm
  entrypoint: src/main/java/com/launchdarkly/tutorial/Main.java
---

```java
package com.launchdarkly.tutorial;

import java.util.HashMap;
import java.util.Map;
import dev.openfeature.sdk.OpenFeatureAPI;
import dev.openfeature.sdk.Client;
import dev.openfeature.sdk.EvaluationContext;
import dev.openfeature.sdk.ImmutableContext;
import dev.openfeature.sdk.Value;
import com.launchdarkly.openfeature.serverprovider.Provider;

public class Main {
    public static void main(String[] args) throws Exception {
        Provider provider = new Provider(System.getenv("LAUNCHDARKLY_SDK_KEY"));
        OpenFeatureAPI.getInstance().setProvider(provider);
        Client client = OpenFeatureAPI.getInstance().getClient();

{{ body }}

        client.getBooleanValue(System.getenv("LAUNCHDARKLY_FLAG_KEY"), false, context);
        System.out.println("feature flag evaluates to true");
    }
}
```
