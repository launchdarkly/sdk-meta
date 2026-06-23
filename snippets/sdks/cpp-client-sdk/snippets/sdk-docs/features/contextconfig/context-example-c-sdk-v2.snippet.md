---
id: cpp-client-sdk/sdk-docs/features/contextconfig/context-example-c-sdk-v2
sdk: cpp-client-sdk
kind: reference
lang: c
description: User example for C client SDK v2.x (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-c

---

```c
struct LDJSON *attributes, *groups;
groups = LDNewArray();
LDArrayPush(groups, LDNewText("Acme"));
LDArrayPush(groups, LDNewText("Global Health Services"));
attributes = LDNewObject();
LDObjectSetKey(attributes, "groups", groups);
struct LDUser *user = LDUserNew("example-user-key");
LDUserSetFirstName(user, "Sandy");
LDUserSetLastName(user, "Smith");
LDUserSetCustomAttributesJSON(user, attributes);
```
