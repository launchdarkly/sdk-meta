---
id: java-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-evaluationdetail-5-0-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "5.0 syntax in section \"Understanding changes to EvaluationDetail\""
---

```java
// 5.x model: create an EvaluationDetail instance, maybe for testing
EvaluationDetail<String> myValue = EvaluationDetail.fromValue("x", 1, EvaluationReason.off());

// 5.x model: check the variation index of the result
EvaluationDetail<String> resultDetail =
  client.stringVariationDetail(flagKey, user, "default string value");
int variation = resultDetail.getVariationIndex();
if (variation == EvaluationDetail.NO_VARIATION) {
  // do something for the default value case
  // note that "if (resultDetail.isDefaultValue())" would also have worked
} else {
  doSomethingWithVariation(variation);
}
```
