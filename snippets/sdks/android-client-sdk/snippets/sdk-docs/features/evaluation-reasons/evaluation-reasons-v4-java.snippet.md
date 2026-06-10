---
id: android-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v4-java
sdk: android-client-sdk
kind: reference
lang: java
description: Flag evaluation reason example for Android SDK v4.x (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only-v4

---

```java
LDConfig ldConfig = new LDConfig.Builder()
  .mobileKey("example-mobile-key")
  .evaluationReasons(true)
  .build();
LDClient client = LDClient.init(this.getApplication(), ldConfig, context, secondsToBlock);

EvaluationDetail<Boolean> detail =
  client.boolVariationDetail("example-flag-key", false);
  // or stringVariationDetail for a string-valued flag, etc.

boolean value = detail.getValue();
Integer index = detail.getVariationIndex();
EvaluationReason reason = detail.getReason();
```
