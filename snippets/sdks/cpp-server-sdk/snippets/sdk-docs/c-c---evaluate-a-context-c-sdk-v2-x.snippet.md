---
id: cpp-server-sdk/sdk-docs/c-c---evaluate-a-context-c-sdk-v2-x
sdk: cpp-server-sdk
kind: reference
lang: c
description: "C SDK v2.x in section \"Evaluate a context\""
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c
---

```c
user = LDUserNew("example-user-key");
LDUserSetName(user, "Sandy");

LDBoolean show_feature = LDBoolVariation(client, user, "example-flag-key", false, NULL);
if (show_feature) {
    // application code to show the feature
} else {
    // the code to run if the feature is off
}
```
