---
id: java-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-evaluationdetail-4-x-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "4.x syntax in section \"Understanding changes to EvaluationDetail\""
---

```java
// 4.x model: create an EvaluationDetail instance, maybe for testing
EvaluationDetail<String> myValue = new EvaluationDetail<>(EvaluationReason.off(), 1, "x");

// 4.x model: check the variation index of the result
EvaluationDetail<String> resultDetail =
  client.stringVariationDetail(flagKey, user, "default string value");
Integer variation = resultDetail.getVariationIndex();
if (variation == null) {
  // do something for the default value case
  // note that "if (resultDetail.isDefaultValue())" would also have worked
} else {
  doSomethingWithVariation(variation.intValue());
}
```
