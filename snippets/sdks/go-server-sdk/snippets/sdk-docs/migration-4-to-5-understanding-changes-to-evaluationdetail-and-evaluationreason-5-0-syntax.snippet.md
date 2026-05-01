---
id: go-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-evaluationdetail-and-evaluationreason-5-0-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "5.0 syntax in section \"Understanding changes to EvaluationDetail and EvaluationReason\""
---

```go
import (
  "gopkg.in/launchdarkly/go-sdk-common.v2/ldreason"
)

value, detail, _ := client.BoolVariationDetail(flagKey, user, false)

// 5.x model: check if the variation index is 0
if detail.VariationIndex.IsDefined() && detail.VariationIndex.IntValue() == 0 {
    log.Printf("it is variation zero")
}

// 5.x model: check if the reason is "prerequisite failed"
if detail.Reason.GetKind() == ldreason.EvalReasonPrerequisiteFailed {
    log.Printf("the prerequisite %s failed", detail.Reason.GetPrerequisiteKey())
}
```
