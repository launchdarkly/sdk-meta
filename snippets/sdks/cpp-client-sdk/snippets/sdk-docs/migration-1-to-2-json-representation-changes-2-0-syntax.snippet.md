---
id: cpp-client-sdk/sdk-docs/migration-1-to-2-json-representation-changes-2-0-syntax
sdk: cpp-client-sdk
kind: reference
lang: c
description: "2.0 syntax in section \"JSON representation changes\""
---

```c
struct LDJSON *names, *tmp;

names = LDNewArray();

tmp = LDNewText("alice");
LDArrayPush(names, tmp);

tmp = LDNewText("bob");
LDArrayPush(names, tmp);

LDJSONFree(names);
```
