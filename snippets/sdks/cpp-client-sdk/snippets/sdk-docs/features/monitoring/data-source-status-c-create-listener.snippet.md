---
id: cpp-client-sdk/sdk-docs/features/monitoring/data-source-status-c-create-listener
sdk: cpp-client-sdk
kind: reference
lang: c
description: Data source status listener connection for C++ (client-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
struct LDDataSourceStatusListener listener;
LDDataSourceStatusListener_Init(&listener);

listener.StatusChanged = OnDataSourceStatusChanged;

/* You may optionally assign the UserData pointer, which will be passed into StatusChanged. */
/* listener.UserData = &some_struct; */

LDListenerConnection connection =
        LDClientSDK_DataSourceStatus_OnStatusChange(sdk, listener);

/* You can disconnect the listener later */
LDListenerConnection_Disconnect(connection);
```
