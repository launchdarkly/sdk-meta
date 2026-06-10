---
id: cpp-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-cpp-c-v3-0
sdk: cpp-client-sdk
kind: reference
lang: c
description: Flag evaluation reason example for C++ (client-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
LDEvalDetail detail;
if (LDClientSDK_BoolVariationDetail(client, "example-flag-key", false, &detail)) {
    printf("Value was true!\n");
} else {
    LDEvalReason reason;
    if (LDEvalDetail_Reason(detail, &reason)) {
       printf("Value was false because of %d\n", LDEvalReason_Kind(reason));
    } else {
       printf("No reason provided to explain why flag was false!\n");
    }
}
```
