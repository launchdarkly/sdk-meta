---
id: java-server-sdk/sdk-docs/features/testdata/flag-behavior-v6
sdk: java-server-sdk
kind: reference
lang: java
description: Configuring test data flag behavior for Java SDK v6.0.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
// This flag is true for the context with the key "example-context-key" and kind of "organization",
// and false for everyone else
td.update(td.flag("example-flag-key")
  .variationForKey(ContextKind.of("organization"), "example-context-key", true)
  .fallthroughVariation(false));

// This flag returns the string variation "green" for contexts who have the custom
// attribute "admin" with a value of true, and "red" for everyone else.
td.update(td.flag("example-string-flag-key")
  .variations(LDValue.of("red"), LDValue.of("green"))
  .fallthroughVariation(0)
  .ifMatch("admin", LDValue.of(true))
  .thenReturn(1));
```
