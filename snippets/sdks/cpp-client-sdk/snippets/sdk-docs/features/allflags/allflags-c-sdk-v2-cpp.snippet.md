---
id: cpp-client-sdk/sdk-docs/features/allflags/allflags-c-sdk-v2-cpp
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: All flags example for the C client SDK v2.x (C++ binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-cpp

---

```cpp
LDJSON *allFlagsObject = client->getAllFlags();

const LDJSON *iter;
for (iter = LDGetIter(allFlagsObject); iter; iter = LDIterNext(iter)) {
   char *serialized_value = LDJSONSerialize(iter);
   std::cout << LDIterKey(iter) << ": " << serialized_value << std::endl;
   LDFree(serialized_value);
}

LDJSONFree(allFlagsObject);
```
