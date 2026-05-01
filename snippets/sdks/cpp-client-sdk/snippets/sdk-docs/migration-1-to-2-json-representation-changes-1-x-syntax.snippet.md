---
id: cpp-client-sdk/sdk-docs/migration-1-to-2-json-representation-changes-1-x-syntax
sdk: cpp-client-sdk
kind: reference
lang: c
description: "1.x syntax in section \"JSON representation changes\""
---

```c
LDNode *names;

names = LDNodeCreateArray();

LDNodeAppendString(&names, "alice");
LDNodeAppendString(&names, "bob");

LDNodeFree(names);
```
