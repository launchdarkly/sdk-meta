---
id: cpp-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Flag evaluation reason example for C++ (server-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```c
LDEvalDetail detail;
if (LDServerSDK_BoolVariationDetail(client, context, "example-flag-key", false, &detail)) {
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
