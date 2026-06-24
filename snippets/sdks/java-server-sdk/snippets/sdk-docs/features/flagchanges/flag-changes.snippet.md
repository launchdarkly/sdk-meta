---
id: java-server-sdk/sdk-docs/features/flagchanges/flag-changes
sdk: java-server-sdk
kind: reference
lang: java
description: Flag change subscription example for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only-members

---

```java
void logWheneverAnyFlagChanges(LDClient client) {
    client.getFlagTracker().addFlagChangeListener(event -> {
        System.out.printf("Flag \"%s\" has changed\n", event.getKey());
    });
}
```
