---
id: cpp-client-sdk/sdk-docs/features/flagchanges/flag-changes-cpp-c-v3-0-create-connection
sdk: cpp-client-sdk
kind: reference
lang: c
description: Flag change listener connection for C++ (client-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
struct LDFlagListener listener;
LDFlagListener_Init(&listener);

listener.FlagChanged = OnFlagChange;

/* You may optionally assign the UserData pointer, which will be passed into FlagChanged. */
/* listener.UserData = &some_struct; */

LDListenerConnection connection =
       LDClientSDK_FlagNotifier_OnFlagChange(sdk, "example-flag-key", listener);

/* You can disconnect the listener later */
 LDListenerConnection_Disconnect(connection);
```
