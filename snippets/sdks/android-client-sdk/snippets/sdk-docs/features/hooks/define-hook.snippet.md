---
id: android-client-sdk/sdk-docs/features/hooks/define-hook
sdk: android-client-sdk
kind: reference
lang: java
description: Hook implementation and configuration for the Android SDK.
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java

public class ExampleHook extends Hook {

  public ExampleHook(String name) {
    super(name);
  }

  // Implement at least one of:
  //
  // * `beforeEvaluation` - called during the execution of a variation method
  // before the flag value has been determined
  //
  // * `afterEvaluation` - called during the execution of a variation method
  // after the flag value has been determined
  //
  // * `beforeIdentify` - called during the execution of the identify process
  // before the operation completes, but after any context modifications are performed
  //
  // * `afterIdentify` - called during the execution of the identify process
  // after the operation completes
  //
  // * `afterTrack` - called during the execution of the track process
  // after the event has been enqueued
}

List<Hook> hookList = new ArrayList<>();
ExampleHook exampleHook = new ExampleHook("Example hook");
hookList.add(exampleHook);

LDConfig ldConfig = new LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("example-mobile-key")
    .hooks(
      Components.hooks()
        .setHooks(hookList)
    )
    .build();

LDContext context = LDContext.create("example-context-key");

LDClient client = LDClient.init(this.getApplication(), ldConfig, context, 0);

```
