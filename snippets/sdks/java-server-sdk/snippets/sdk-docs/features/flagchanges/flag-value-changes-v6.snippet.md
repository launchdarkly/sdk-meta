---
id: java-server-sdk/sdk-docs/features/flagchanges/flag-value-changes-v6
sdk: java-server-sdk
kind: reference
lang: java
description: Flag value change subscription example for Java SDK v6.0.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only-members

---

```java
void logWheneverOneFlagChangesForOneContext(LDClient client, String flagKey, LDContext context) {
    client.getFlagTracker().addFlagValueChangeListener(flagKey, context, event -> {
        System.out.printf("Flag \"%s\" for context \"%s\" has changed from %s to %s\n", event.getKey(),
            context.getKey(), event.getOldValue(), event.getNewValue());
    });
}
```
