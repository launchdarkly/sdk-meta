---
id: go-server-sdk/sdk-docs/features/evaluation-reasons/print-reason
sdk: go-server-sdk
kind: reference
lang: go
description: Reason-object inspection example for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "github.com/launchdarkly/go-sdk-common/v3/ldreason"
)

func PrintReason(reason ldreason.EvaluationReason) {
  switch reason.GetKind() {
  case ldreason.EvalReasonOff:
    fmt.Println("it's off")
  case ldreason.EvalReasonFallthrough:
    fmt.Println("fell through")
  case ldreason.EvalReasonTargetMatch:
    fmt.Println("targeted")
  case ldreason.EvalReasonRuleMatch:
    fmt.Printf("matched rule %d/%s\n", reason.GetRuleIndex(), reason.GetRuleID())
  case ldreason.EvalReasonPrerequisiteFailed:
    fmt.Printf("prereq failed: %s\n", reason.GetPrerequisiteKey())
  case ldreason.EvalReasonError:
    fmt.Printf("error: %s\n", reason.GetErrorKind())
  }
  // or, if all you want is a simple descriptive string:
  fmt.Println(reason)
}
```
