---
id: cpp-server-sdk/sdk-docs/features/allflags/allflags-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: All flags example for C++ (server-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```c
LDAllFlagsState state = LDServerSDK_AllFlagsState(client, context, LD_ALLFLAGSSTATE_DEFAULT);

if (LDAllFlagsState_Valid(state)) {
    /* all flags were evaluated successfully! */
} else {
    /* there was an issue, but it's still possible to serialize the (empty) state. */
}

/* Retrieve a specific flag's value */
LDValue bool_value_reference = LDAllFlagsState_Value(state, "my-bool-flag");
if (LDValue_GetBool(bool_value_reference)) {
    /* value is true! */
}

/* Serialize the state, suitable for bootstrapping a client-side SDK e.g. JavaScript */
char* json = LDAllFlagsState_SerializeJSON(state);

LDMemory_FreeString(json);

LDAllFlagsState_Free(state);

```
