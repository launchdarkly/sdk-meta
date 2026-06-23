---
id: cpp-client-sdk/sdk-docs/features/monitoring/data-source-status-c-define-callback
sdk: cpp-client-sdk
kind: reference
lang: c
description: Data source status callback definition for C++ (client-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-toplevel

---

```c
void OnDataSourceStatusChanged(LDDataSourceStatus status, void* user_data) {
  printf("status: %d\n", LDDataSourceStatus_GetState(status));
}
```
