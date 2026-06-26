---
id: cpp-server-sdk/sdk-docs/features/monitoring/data-source-status-c-create-listener
sdk: cpp-server-sdk
kind: reference
lang: c
description: Data source status listener connection for C++ (server-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```c
struct LDServerDataSourceStatusListener listener;
LDServerDataSourceStatusListener_Init(&listener);

listener.StatusChanged = OnDataSourceStatusChanged;

/* You may optionally assign the UserData pointer, which will be passed into
* StatusChanged. */

/* listener.UserData = &some_struct; */

LDListenerConnection connection = LDServerSDK_DataSourceStatus_OnStatusChange(client, listener);

/* You can disconnect the listener later */
LDListenerConnection_Disconnect(connection);
```
