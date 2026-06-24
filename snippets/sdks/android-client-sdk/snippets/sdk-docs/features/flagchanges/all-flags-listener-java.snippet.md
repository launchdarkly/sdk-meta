---
id: android-client-sdk/sdk-docs/features/flagchanges/all-flags-listener-java
sdk: android-client-sdk
kind: reference
lang: java
description: All-flags update listener example for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
LDAllFlagsListener listener = new LDAllFlagsListener() {
    @Override
    public void onChange(List<String> flagKeys) {
        // Get new values for flagKeys or other operations
    }
};

// register all flags listener
LDClient.get().registerAllFlagsListener(listener);
// when done with all flags listener it should be unregistered
LDClient.get().unregisterAllFlagsListener(listener);
```
