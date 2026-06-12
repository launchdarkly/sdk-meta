---
id: android-client-sdk/sdk-docs/features/testdata/flag-behavior-java
sdk: android-client-sdk
kind: reference
lang: java
description: Configuring test data flag behavior for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java
// This flag is true for the context with the key "example-context-key" and kind of "organization",
// and false for everyone else
td.update(td.flag("example-flag-key")
  .variation(false)
  .variationForKey(ContextKind.of("organization"), "example-context-key", true));

// This flag returns the string variation "green" for contexts who have the custom
// attribute "admin" with a value of true, and "red" for everyone else.
td.update(td.flag("example-flag-key")
  .variations(LDValue.of("red"), LDValue.of("green"))
  .variationIndexFunc(context ->
    context.getValue("admin").booleanValue() ? 1 : 0));
```
