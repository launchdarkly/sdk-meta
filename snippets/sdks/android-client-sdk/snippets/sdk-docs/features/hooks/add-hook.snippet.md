---
id: android-client-sdk/sdk-docs/features/hooks/add-hook
sdk: android-client-sdk
kind: reference
lang: java
description: Adding a hook to an existing client for the Android SDK.
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java
List<Hook> hookList = new ArrayList<>();
ExampleHook exampleHook = new ExampleHook("Example hook");
hookList.add(exampleHook);

LDClient client = LDClient.init(this.getApplication(), ldConfig, context, 0);
client.addHook(exampleHook);

```
