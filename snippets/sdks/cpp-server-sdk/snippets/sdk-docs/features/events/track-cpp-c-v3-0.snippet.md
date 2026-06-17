---
id: cpp-server-sdk/sdk-docs/features/events/track-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Custom event tracking example for C++ (server-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```c
/* track the event */
LDServerSDK_TrackEvent(client, context, "example-event-key");

/* track the event and associate data with it */
LDServerSDK_TrackData(client, context, "example-event-key", LDValue_NewNumber(42));

```
