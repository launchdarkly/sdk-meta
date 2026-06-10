---
id: dotnet-server-sdk/sdk-docs/features/evaluation-reasons/print-reason
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Reason-object inspection example for .NET (server-side).
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only-members

---

```csharp
void PrintReason(EvaluationReason reason)
{
  switch (reason.Kind)
  {
    case EvaluationReasonKind.Off:
      Console.WriteLine("it's off");
      break;
    case EvaluationReasonKind.Fallthrough:
      Console.WriteLine("fell through");
      break;
    case EvaluationReasonKind.TargetMatch:
      Console.WriteLine("targeted");
      break;
    case EvaluationReasonKind.RuleMatch:
      Console.WriteLine("matched rule " + reason.RuleIndex + "/" + reason.RuleId);
      break;
    case EvaluationReasonKind.PrerequisiteFailed:
      Console.WriteLine("prereq failed: " + reason.PrerequisiteKey);
      break;
    case EvaluationReasonKind.Error:
      Console.WriteLine("error: " + reason.ErrorKind);
      break;
  }
  // or, if all you want is a simple descriptive string:
  Console.WriteLine(reason.ToString());
}
```
