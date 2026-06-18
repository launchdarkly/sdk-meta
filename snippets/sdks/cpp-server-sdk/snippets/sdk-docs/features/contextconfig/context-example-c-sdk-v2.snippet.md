---
id: cpp-server-sdk/sdk-docs/features/contextconfig/context-example-c-sdk-v2
sdk: cpp-server-sdk
kind: reference
lang: c
description: User example for C server SDK v2.x (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c

---

```c
struct LDUser *user = LDUserNew("example-user-key");
LDUserSetFirstName(user, "Sandy");
LDUserSetLastName(user, "Smith");
LDUserSetEmail(user, "sandy@example.com");

struct LDJSON *tmp;
struct LDJSON *custom = LDNewObject();
struct LDJSON *groups = LDNewArray();
tmp = LDNewText("Acme");
LDArrayPush(groups, tmp);
tmp = LDNewText("Global Health Services");
LDArrayPush(groups, tmp);
LDObjectSetKey(custom, "groups", groups);

LDUserSetCustom(user, custom);
```
