---
id: cpp-client-sdk/sdk-docs/features/allflags/allflags-c-sdk-v2
sdk: cpp-client-sdk
kind: reference
lang: c
description: All flags example for the C client SDK v2.x (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-c

---

```c
struct LDJSON *allFlagsObject = LDAllFlags(client);

const struct LDJSON *iter;
for (iter = LDGetIter(allFlagsObject); iter; iter = LDIterNext(iter)) {
   char *serialized_value = LDJSONSerialize(iter);
   printf("%s: %s\n", LDIterKey(iter), serialized_value);
   LDFree(serialized_value);
}

LDJSONFree(allFlagsObject);
```
