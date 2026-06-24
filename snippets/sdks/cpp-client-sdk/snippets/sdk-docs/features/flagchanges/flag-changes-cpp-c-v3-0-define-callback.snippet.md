---
id: cpp-client-sdk/sdk-docs/features/flagchanges/flag-changes-cpp-c-v3-0-define-callback
sdk: cpp-client-sdk
kind: reference
lang: c
description: Flag change callback definition for C++ (client-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-toplevel

---

```c
void OnFlagChange(char const* flag_key,
                      LDValue new_value,
                      LDValue old_value,
                      bool deleted,
                      void* user_data) {
     if (deleted) {
       printf("The flag %s was deleted\n", flag_key);
     } else {
       printf("The flag %s was updated\n", flag_key);
     }
}
```
