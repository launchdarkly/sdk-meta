---
id: cpp-client-sdk/sdk-docs/features/allflags/allflags-cpp-c-v3-0
sdk: cpp-client-sdk
kind: reference
lang: c
description: All flags example for C++ (client-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
LDValue all_flags = LDClientSDK_AllFlags(client);

LDValue_ObjectIter it;

for (it = LDValue_ObjectIter_New(all_flags); !LDValue_ObjectIter_End(it); LDValue_ObjectIter_Next(it)) {

  char const* flag_key = LDValue_ObjectIter_Key(it);
  LDValue flag_val_ref = LDValue_ObjectIter_Value(it);

  if (LDValue_Type(flag_val_ref) == LDValueType_Bool) {
      printf("%s: %d\n", flag_key, LDValue_GetBool(flag_val_ref));
  }
}
```
