---
id: android-client-sdk/sdk-docs/features/flagchanges/flag-changes-java
sdk: android-client-sdk
kind: reference
lang: java
description: Flag change listener registration for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
String flagKey = "yourFlagKey";

FeatureFlagChangeListener listener = new FeatureFlagChangeListener() {
    @Override
    public void onFeatureFlagChange(String flagKey) {
        try {
            boolean newValue = LDClient.get().boolVariation(flagKey, false);
        } catch (LaunchDarklyException e) {
            // LDClient.get() throws if the client has not been initialized
        }
    }
};

LDClient.get().registerFeatureFlagListener(flagKey, listener);
```
