---
id: java-server-sdk/sdk-docs/features/hooks/define-hook
sdk: java-server-sdk
kind: reference
lang: java
description: Hook implementation and configuration for the Java SDK.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
import com.launchdarkly.sdk.server.integrations.Hook;

class ExampleHook extends Hook {

  public ExampleHook(String name) {
    super(name);
  }

  // Implement at least one of `beforeEvaluation`, `afterEvaluation`

  // `beforeEvaluation` is called during the execution of a variation method
  // before the flag value has been determined

  // `afterEvaluation` is called during the execution of a variation method
  // after the flag value has been determined
}

ExampleHook exampleHook = new ExampleHook("example-hook");

LDConfig config = new LDConfig.Builder()
    .hooks(
        Components.hooks().setHooks(Collections.singletonList(exampleHook)))
    .build();

LDClient client = new LDClient("YOUR_SDK_KEY", config);
```
