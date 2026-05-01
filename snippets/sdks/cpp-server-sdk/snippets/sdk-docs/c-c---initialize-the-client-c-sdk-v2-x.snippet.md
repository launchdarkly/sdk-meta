---
id: cpp-server-sdk/sdk-docs/c-c---initialize-the-client-c-sdk-v2-x
sdk: cpp-server-sdk
kind: reference
lang: c
description: "C SDK v2.x in section \"Initialize the client\""
---

```c
unsigned int maxwaitmilliseconds = 10 * 1000;
struct LDConfig *config = LDConfigNew("YOUR_SDK_KEY");
/* blocks on initialization */
struct LDClient *client = LDClientInit(config, maxwaitmilliseconds);
```
