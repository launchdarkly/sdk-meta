---
id: cpp-client-sdk/sdk-docs/features/anonymous/anonymous-c-sdk-v2-cpp
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Anonymous user example for the C client SDK v2.x, C++ binding.
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-cpp
---

```cpp
/* anonymous user, with key specified */
struct LDUser *user = LDUserNew("example-user-key");
LDUserSetAnonymous(user, true);

/*
  In v2.x, when you mark the context as anonymous, you can
  leave the key parameter null and the client will automatically
  set it to a LaunchDarkly-specific, device-unique string
  that is consistent between app restarts and device reboots.
*/
LDUser *anonymousUser = LDUserNew(nullptr);
LDUserSetAnonymous(anonymousUser, true);
```
