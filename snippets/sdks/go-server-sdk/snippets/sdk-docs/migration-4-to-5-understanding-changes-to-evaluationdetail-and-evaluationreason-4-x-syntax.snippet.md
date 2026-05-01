---
id: go-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-evaluationdetail-and-evaluationreason-4-x-syntax
sdk: go-server-sdk
kind: reference
lang: go
description: "4.x syntax in section \"Understanding changes to EvaluationDetail and EvaluationReason\""
---

```go
import (
  ld "gopkg.in/launchdarkly/go-server-sdk.v4"
)

value, detail, _ := client.BoolVariationDetail(flagKey, user, false)

// 4.x model: check if the variation index is 0
if detail.VariationIndex != nil && *detail.VariationIndex == 0 {
    log.Printf("it is variation zero")
}

// 4.x model: check if the reason is "prerequisite failed"
if pf, ok := detail.Reason.(ld.EvaluationReasonPrerequisiteFailed); ok {
    log.Printf("the prerequisite %s failed", pf.PrerequisiteKey)
}
```
